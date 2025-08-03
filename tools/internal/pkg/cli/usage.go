package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bastean/x/tools/internal/pkg/log"
)

func Usage(app string) {
	log.Logo(app)

	fmt.Printf("Usage: %s [flags]\n\n", strings.ToLower(app))

	flag.PrintDefaults()
}
