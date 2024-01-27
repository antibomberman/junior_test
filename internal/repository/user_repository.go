package repository

import (
	"database/sql"
	"fmt"
	"github.com/antibomberman/junior_test/internal/models"
)

type UserRepository interface {
	GetUserByID(id string) (models.User, error)
	CreateUser(data models.UserCreate) (int64, error)
	UpdateUser(id string, data models.UserUpdate) (int64, error)
	GetAllUsers(minAge, maxAge uint8, nameFilter string) ([]models.User, error)
	HasUser(id string) (bool, error)
	DeleteUser(id string) error
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByID(id string) (models.User, error) {
	user := models.User{}
	err := r.db.QueryRow("SELECT id, name, surname,patronymic,gender,age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Gender, &user.Age)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(data models.UserCreate) (int64, error) {
	result, err := r.db.Exec("INSERT INTO users (name) VALUES ($1) RETURNING id", data.Name)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
func (r *userRepository) UpdateUser(id string, data models.UserUpdate) (int64, error) {
	fmt.Println(id, data)
	result, err := r.db.Exec("INSERT INTO users (name) VALUES ($1) RETURNING id", data.Name)
	if err != nil {
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *userRepository) GetAllUsers(minAge, maxAge uint8, nameFilter string) ([]models.User, error) {
	query := "SELECT * FROM users WHERE true"

	if minAge > 0 {
		query += fmt.Sprintf(" AND age >= %d", minAge)
	}
	if maxAge > 0 {
		query += fmt.Sprintf(" AND age <= %d", maxAge)
	}
	if nameFilter != "" {
		query += fmt.Sprintf(" AND name ILIKE '%%%s%%'", nameFilter)
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Gender, &user.Nationality, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (r *userRepository) HasUser(id string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (r *userRepository) DeleteUser(userID string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
