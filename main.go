package main

import (
	"fmt"

	"github.com/src-d/go-billy/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func main() {
	fs := memfs.New()

	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: "https://github.com/crhntr/explore-go-git-wasm",
	})
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(repo.Head())
}
