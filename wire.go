//go:build wireinject
// +build wireinject

package main

import (
	"_template_/model"
	"_template_/registry"
	"_template_/uow"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func NewRegistry(conn *gorm.DB) registry.Registry {
	wire.Build(
		registry.New,
		model.NewRepository,
		model.NewUserRepository,
		uow.New,
	)
	return nil
}
