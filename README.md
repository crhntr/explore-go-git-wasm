# Using the go-git package in the browser

## Abstract

Go now supports [Web Assembly]((https://github.com/golang/go/wiki/WebAssembly)), which runs in most browsers.
Go-Git supports fully [in memory](https://github.com/go-git/go-git#in-memory-example) plumbing and "ceramic" operations.
Git can be hard to understand and DOM (document object model) visualizations of the state could be really cool.
I tried to compile a Go program that imports and uses the go-git package.
By refactoring when and how the osfs package gets imported in go-git we can use git in the browser.

I have a tagged release here with the required go-git changes: https://github.com/crhntr/go-git/releases/tag/wip-js-wasm

To get my proof-of-concept release run "go get github.com/crhntr/go-git@wip-js-wasm"
and add `replace github.com/go-git/go-git/v5 v5.2.0 =>  github.com/crhntr/go-git/v5 <the version you see in the 'go get' output>` to your module file. 

## Methodology

- Get imported modules

  `go mod vendor`

- Find usages of packages that do non-wasm-supproted os syscalls

  go-git abstracts the filesystem and OS interactions using the go-billy Filesystem. 
  The go-billy/inmem Filesystem works in browser wasm environments [see github.com/crhntr/explore-go-billy-wasm](https://github.com/crhntr/explore-go-billy-wasm).
  
  The go-billy implementation that is not supported in the browser is [go-billy/osfs](https://github.com/go-git/go-billy/tree/master/osfs), which wraps the standard library filesystem package.
  
  To find uses of osfs, I ran the following
  
  ```sh
  cd vendor
  ag gopkg.in/src-d/go-billy.v4/osfs
  ```
  
  ```
  github.com/go-git/go-git/v5/remote.go
  9:      "github.com/go-git/go-billy/v5/osfs"
  
  github.com/go-git/go-git/v5/storage/filesystem/dotgit/dotgit.go
  17:     "github.com/go-git/go-billy/v5/osfs"
  
  github.com/go-git/go-git/v5/plumbing/transport/server/loader.go
  10:     "github.com/go-git/go-billy/v5/osfs"
  
  github.com/go-git/go-git/v5/repository.go
  33:     "github.com/go-git/go-billy/v5/osfs"
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
  
  I then factored out the usages of osfs (see diff here: https://github.com/crhntr/go-git/commit/09378a6e8d4bb3078f6f4de990c22c40077bafd5).

  I did not check for network usage that would not be supported in Go wasm environments.

## Results

There are a few imports that do not allow go-git to be imported in a Go program compiled for the browser. Some refactors could make this work.

## Developer Notes

`./start.sh` is a script that starts a web server that loads the wasm into the browser.
