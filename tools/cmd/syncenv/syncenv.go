package main

import (
	"github.com/bastean/x/tools/internal/app/syncenv"
	"github.com/bastean/x/tools/internal/pkg/log"
)

func main() {
	if err := syncenv.Up(); err != nil {
		log.Fatal(err.Error())
	}
}
