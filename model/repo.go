package model

import "_template_/domain/repository"

var _ repository.Repository = (*repo)(nil)

type repo struct {
	users repository.IUsers
}

// Users implements repository.Repository
func (r *repo) Users() repository.IUsers {
	return r.users
}

func NewRepository(users repository.IUsers) repository.Repository {
	return &repo{
		users: users,
	}
}
