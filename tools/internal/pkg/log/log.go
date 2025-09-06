package log

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"

	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

const (
	Font = "speed"
)

var (
	Log     = log.New()
	Debug   = Log.Debug
	Info    = Log.Info
	Success = Log.Success
)

func Logo(name string) {
	figure.NewFigure(name, Font, true).Print()
	println()
}

func Error(what string) {
	Log.Error("Error: %s", what)
}

func Fatal(what string) {
	Log.Fatal("Error: %s", what)
}

func Starting() {
	Info("Starting...")
}

func Created(values ...string) {
	for _, value := range values {
		Success(fmt.Sprintf("Created: %q", value))
	}
}

func Completed() {
	Success("Completed!")
}
