package domain

var registryObj IRegistry

var _ IRegistry = (*registry)(nil)

type IRegistry interface {
	Repository() IRepository
}

type registry struct {
	repository IRepository
}

// Repository implements IRegistry
func (r *registry) Repository() IRepository {
	return r.repository
}

func Registry() IRegistry {
	return registryObj
}

func SetRegistry(r IRegistry) {
	registryObj = r
}

func InitRegistry() IRegistry {
	return nil
}
