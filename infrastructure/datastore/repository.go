package datastore

import "_template_/domain"

var _ domain.IRepository = (*repository)(nil)

type repository struct{}
