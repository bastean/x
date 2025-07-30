package errs

import (
	"fmt"
	"os"
)

func Fatal(what string) {
	fmt.Printf("Error: %q\n", what)
	os.Exit(1)
}
