package main

import (
	"_template_/domain"

	"github.com/google/wire"
)

func InitRegistry() domain.IRegistry {
	wire.Build(
		domain.InitRegistry,
	)
	return nil
}
