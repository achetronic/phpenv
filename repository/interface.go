package repository

type RepositoryManager interface {
	AddRepository() (bool, string)
	UpdatePackages() (bool, string)
}
