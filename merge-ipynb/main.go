package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nishanths/merge-ipynb"
)

const (
	helpText = `merge-ipynb: merge ipython notebooks
usage: merge-ipynb <p1.ipynb> <p2.ipynb>...`
)

func main() {
	flag.Parse()
	args := os.Args[1:]

	// Print help and exit
	if len(args) == 0 || (len(args) > 0 && (args[0] == "-h" || args[0] == "--help")) {
		fmt.Println(helpText)
		os.Exit(0)
	}

	f1, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	merge.Merge(bytes.NewBuffer(nil), f1, f2)
}
