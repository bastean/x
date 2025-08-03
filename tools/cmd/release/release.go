package main

import (
	"github.com/bastean/x/tools/internal/app/release"
	"github.com/bastean/x/tools/internal/pkg/log"
)

func main() {
	if err := release.Up(); err != nil {
		log.Fatal(err.Error())
	}
}
