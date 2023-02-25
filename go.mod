module github.com/crhntr/explore-go-git-wasm

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/net v0.7.0 // indirect
)

replace github.com/go-git/go-git/v5 v5.2.0 => github.com/crhntr/go-git/v5 v5.2.1-0.20201117062125-09378a6e8d4b
