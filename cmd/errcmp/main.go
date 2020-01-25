package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/hashicorp/go-version"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/moricho/errcmp"
)

func main() {
	runtimev := runtime.Version()
	runtimev = runtimev[2:]
	v1, err := version.NewVersion(runtimev)
	if !errors.Is(err, nil) {
		log.Fatalf("failed to convert: %v", err)
		return
	}
	v2, err := version.NewVersion("1.13")
	if !errors.Is(err, nil) {
		log.Fatalf("failed to convert: %v", err)
		return
	}
	if v1.LessThan(v2) {
		return
	} else {
		fmt.Println("errcmp checking...")
	}

	singlechecker.Main(errcmp.Analyzer)
}
