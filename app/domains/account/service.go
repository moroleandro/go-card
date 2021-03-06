package account

import (
	entities "github.com/guil95/go-card/app/entities/account"
	"github.com/guil95/go-card/app/vo/uuid"
	"github.com/guil95/go-card/infra/repositories"
)

type Service struct {
	repo *repositories.AccountRepo
}

func NewService(r *repositories.AccountRepo) *Service {
	return &Service{repo: r}
}

func (s Service) ListAccounts() ([]*entities.Account, error) {
	accounts, err := s.repo.List()

	if accounts == nil {
		return nil, entities.ErrorAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s Service) FindAccountByDocument(document string) (*entities.Account, error) {
	account, err := s.repo.FindAccountByDocument(document)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s Service) FindAccountByID(id uuid.ID) (*entities.Account, error) {
	account, err := s.repo.FindAccountByID(id)

	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, entities.ErrorAccountNotFound
	}

	return account, nil
}

func (s Service) CreateAccount(document string) (*entities.Account, error) {
	account := entities.NewAccount(document)

	account, err := s.repo.CreateAccount(account)

	if err != nil {
		return nil, err
	}

	return account, nil
}
