package release

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomModuleRelease() *Module {
	module, err := NewModuleRelease(
		services.Create.LoremIpsumWord(),
		services.Create.RandomString([]string{"patch", "minor", "major"}),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func RandomModuleFirstRelease() *Module {
	module, err := NewModuleFirstRelease(
		services.Create.LoremIpsumWord(),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func ModuleWithInvalidIncrement() (*Module, string) {
	value := "x"

	module := RandomModuleRelease()

	module.Increment = value

	return module, value
}
