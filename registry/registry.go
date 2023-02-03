package registry

import (
	"_template_/domain/repository"
	"_template_/uow"
)

var _ Registry = (*registry)(nil)

type Registry interface {
	repository.Repository
	UOW() uow.UnitOfWork
}

type registry struct {
	repo repository.Repository
	uow  uow.UnitOfWork
}

// UOW implements Registry
func (r *registry) UOW() uow.UnitOfWork {
	return r.uow
}

// Users implements Registry
func (r *registry) Users() repository.IUsers {
	return r.repo.Users()
}
