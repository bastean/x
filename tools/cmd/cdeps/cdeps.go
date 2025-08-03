package main

import (
	"github.com/bastean/x/tools/internal/app/cdeps"
	"github.com/bastean/x/tools/internal/pkg/log"
)

func main() {
	if err := cdeps.Up(); err != nil {
		log.Fatal(err.Error())
	}
}
