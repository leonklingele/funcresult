# funcresult — a Go linter to analyze function result parameters

![build](https://github.com/leonklingele/funcresult/actions/workflows/build.yml/badge.svg)

## Installation

```sh
go install github.com/leonklingele/funcresult@latest
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
