package usecase

import "github.com/yoshikawataiki/simple-api/domain"

// UserRepository model
type UserRepository interface {
	Store(domain.User) (int64, error)
	FindByID(int64) (domain.User, error)
	FindAll() (domain.Users, error)
}
