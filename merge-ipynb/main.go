package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/crime-data/merge-ipynb"
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

	// Open files concurrently
	filenames := args[0:]
	files := make([]io.Reader, len(filenames))
	wg := sync.WaitGroup{}
	wg.Add(len(filenames))
	ch := make(chan error, len(filenames))

	for i, f := range filenames {
		i := i
		f := f

		go func() {
			defer wg.Done()
			file, err := os.Open(f)
			ch <- err
			files[i] = file
		}()
	}

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			log.Fatal(err)
		}
	}

	merge.Merge(os.Stdout, files...)
}
