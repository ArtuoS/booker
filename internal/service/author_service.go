package service

import "github.com/ArtuoS/booker-api/internal/infra"

type AuthorService struct {
	Repository infra.AuthorRepository
}

func NewAuthorService(repository infra.AuthorRepository) AuthorService {
	return AuthorService{
		Repository: repository,
	}
}

func (a *AuthorService) Execute(cmd string) error {
	err := a.Repository.Execute(cmd)
	if err != nil {
		return err
	}

	return nil
}
