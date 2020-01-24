package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/moricho/errcmp"
)

func main() { singlechecker.Main(errcmp.Analyzer) }
