package database

import (
	"time"

	"github.com/yoshikawataiki/simple-api/domain"
)

// UserRepository model
type UserRepository struct {
	SQLHandler
}

// Store insert values into user table
func (repo *UserRepository) Store(u domain.User) (id int64, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int64(id64)
	return
}

// FindByID find the user by id
func (repo *UserRepository) FindByID(identifier int64) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var name string
	var email string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &name, &email, &createdAt, &updatedAt); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return
}

// FindAll find all users
func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, name, email, created_at, updated_at FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var name string
		var email string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &name, &email, &createdAt, &updatedAt); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			Name:      name,
			Email:     email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		users = append(users, user)
	}
	return
}
