# merge-ipynb

> Merge iPython notebook cells with ease

# Package

Package `merge-ipynb` provides a Merge function that merges iPython notebooks and writes the result to the specified `io.Writer`.

Import the package as `merge-ipynb` and refer to it as `merge`.

## Web interface

See [`merge-ipynb-web`]().

## Command-line

Install:

```
$ go get -u github.com/crime-analysis/merge-ipynb/merge-ipynb
```

Usage:

```
$ merge-ipynb <p1.ipynb> <p2.ipynb>...
```

The output is sent to `stdout`, which you can redirect to a file:

```
$ merge-ipynb p1.ipynb p2.ipynb > merged.ipynb
```

## License

MIT.