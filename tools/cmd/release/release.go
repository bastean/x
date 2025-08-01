package main

import (
	"github.com/bastean/x/tools/internal/app/release"
	"github.com/bastean/x/tools/internal/pkg/errs"
)

func main() {
	if err := release.Up(); err != nil {
		errs.Fatal(err.Error())
	}
}
