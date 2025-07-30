package cli

import (
	"flag"
	"fmt"
)

func Usage(name string) {
	fmt.Printf("Usage: %s [flags]\n\n", name)
	flag.PrintDefaults()
}
