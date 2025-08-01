package main

import (
	"github.com/bastean/x/tools/internal/app/syncenv"
	"github.com/bastean/x/tools/internal/pkg/errs"
)

func main() {
	if err := syncenv.Up(); err != nil {
		errs.Fatal(err.Error())
	}
}
