package petani

import "smartfarm-api/pkg/database"

type PetaniRepository interface {
	Create(petani *Petani) error
	FindAll() ([]Petani, error)
	FindByID(id uint) (*Petani, error)
	Update(petani *Petani) error
	Delete(id uint) error
	// 🟢 Tambahkan ini
	FindPaginated(page, limit int) ([]Petani, int64, error)
}

type petaniRepository struct{}

func NewPetaniRepository() PetaniRepository {
	return &petaniRepository{}
}

func (r *petaniRepository) Create(petani *Petani) error {
	return database.DB.Create(&petani).Error
}

func (r *petaniRepository) FindAll() ([]Petani, error) {
	var petanis []Petani
	err := database.DB.Find(&petanis).Error
	return petanis, err
}

func (r *petaniRepository) FindByID(id uint) (*Petani, error) {
	var petani Petani
	err := database.DB.First(&petani, id).Error
	return &petani, err
}

func (r *petaniRepository) Update(petani *Petani) error {
	return database.DB.Save(petani).Error
}

func (r *petaniRepository) Delete(id uint) error {
	return database.DB.Delete(&Petani{}, id).Error
}

func (r *petaniRepository) FindPaginated(page, limit int) ([]Petani, int64, error) {
	var petanis []Petani
	var total int64

	offset := (page - 1) * limit

	if err := database.DB.Model(&Petani{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Find(&petanis).Error; err != nil {
		return nil, 0, err
	}

	return petanis, total, nil
}



