package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	fs := memfs.New()

	repo, err := git.Init(memory.NewStorage(), fs)
	if err != nil {
		fmt.Println(err)
		return
	}

	wt, err := repo.Worktree()
	if err != nil {
		fmt.Println(err)
		return
	}

	exampleTxt, err := wt.Filesystem.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = exampleTxt.Write([]byte(`Hello, world!`))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = exampleTxt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = wt.Add("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	author := object.Signature{Name: "crhntr", Email: "crhntr.com@gmail.com", When: time.Now()}
	_, err = wt.Commit("initial commit", &git.CommitOptions{
		Author:    &author,
		Committer: &author,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(repo.Head())
}
