package registry

import (
	"_template_/domain/repository"
	"_template_/uow"
)

var std Registry

func Repository() repository.Repository {
	return std
}

func UOW() uow.UnitOfWork {
	return std.UOW()
}

func New(
	repo repository.Repository,
	uow uow.UnitOfWork,
) Registry {
	return &registry{
		repo: repo,
		uow:  uow,
	}
}

func Init(r Registry) {
	std = r
}
