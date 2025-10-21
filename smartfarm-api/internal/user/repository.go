package user

import "smartfarm-api/pkg/database"

type Repository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}


// Repository untuk operasi database petani
type userRepository struct{}

func NewRepository() *userRepository {
	return &userRepository{}
}


func (r *userRepository) Create(user *User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*User, error) {
	var user User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Ambil semua user dari database
func (r *userRepository) GetAllUser() ([]User, error) {
	var users []User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}