package main

import (
	"fmt"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
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
