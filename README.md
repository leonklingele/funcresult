# funcresult — a Go linter to analyze function result parameters

## Installation

```sh
go get -u github.com/leonklingele/funcresult/...
funcresult -help
```

## Run analyzer

```sh
funcresult -require-unnamed ./...

# Example output:
GOPATH/src/github.com/leonklingele/funcresult/pkg/analyzer/analyzer.go:18:13: should use unnamed result parameter
```

### Available flags

```
  -require-named
    	require the use of named function result parameters only
  -require-unnamed
    	require the use of unnamed function result parameters only
```
