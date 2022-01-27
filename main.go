package main

import (
	"github.com/leonklingele/funcresult/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.New())
}
