# Using the go-git package in the browser

## Abstract

Go now supports [Web Assembly]((https://github.com/golang/go/wiki/WebAssembly)), which runs in most browsers.
Go-Git supports fully [in memory](https://github.com/src-d/go-git#in-memory-example) plumbing and "ceramic" operations.
Git can be hard to understand and DOM (document object model) visualizations of the state could be really cool.
I tried to compile a Go program that imports and uses the go-git package. The results were that currently
that package imports the [osfs](gopkg.in/src-d/go-billy.v4/osfs) convenience in various root level files. This causes compliation failure.
By refactoring when an how the imports are done in these files, Go-Git would be able to be used in the browser.

## Methodology

- Get imported modules

  `go mod vendor`

- Find usages of packages that do non-wasm-supproted os syscalls

  go-git abstracts the filesystem and OS interactions using the go-billy Filesystem. 
  The go-billy/inmem Filesystem works in browser wasm environments [see github.com/crhntr/explore-go-billy-wasm](https://github.com/crhntr/explore-go-billy-wasm).
  
  The go-billy implementation that is not supported in the browser is [go-billy/osfs](gopkg.in/src-d/go-billy.v4/osfs), which wraps the standard library filesystem package.
  
  To find uses of osfs, I ran the following
  
  ```sh
  cd vendor
  ag gopkg.in/src-d/go-billy.v4/osfs
  ```
  
  ```
  gopkg.in/src-d/go-billy.v4/osfs/os.go
  2:package osfs // import "gopkg.in/src-d/go-billy.v4/osfs"

  gopkg.in/src-d/go-git.v4/remote.go
  9:	"gopkg.in/src-d/go-billy.v4/osfs"

  gopkg.in/src-d/go-git.v4/storage/filesystem/dotgit/dotgit.go
  15:	"gopkg.in/src-d/go-billy.v4/osfs"

  gopkg.in/src-d/go-git.v4/plumbing/transport/server/loader.go
  10:	"gopkg.in/src-d/go-billy.v4/osfs"

  gopkg.in/src-d/go-git.v4/repository.go
  29:	"gopkg.in/src-d/go-billy.v4/osfs"

  modules.txt
  54:gopkg.in/src-d/go-billy.v4/osfs
  ```
  
  The osfs package is used in the following files:
  
  In repository.go
  - func PlainInit
  - func dotGitToOSFilesystems
  - func dotGitFileToOSFilesystem
  
  In remote.go
  - func PushContext
  
  In storage/filesystem/dotgit.go
  - func Alternates
  
  I did not check for network usage that would not be supported in Go wasm environments.

## Results

There are a few imports that do not allow go-git to be imported in a Go program compiled for the browser. Some refactors could make this work.

## Developer Notes

`./start.sh` is a script that starts a web server that loads the wasm into the browser.

## References

1. [Go + WebAssembly]
2. [github.com/src-d/go-git](https://github.com/src-d/go-git)
