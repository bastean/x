package main

import (
	"github.com/bastean/x/tools/internal/app/cdeps"
	"github.com/bastean/x/tools/internal/pkg/errs"
)

func main() {
	if err := cdeps.Up(); err != nil {
		errs.Fatal(err.Error())
	}
}
