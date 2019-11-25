package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/src-d/go-billy/memfs"
)

func main() {
	fs := memfs.New()

	filename := "hello.txt"
	cf, err := fs.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(cf, "Hello, world!")
	cf.Close()

	of, err := fs.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	buf, err := ioutil.ReadAll(of)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buf))
}
