package lahan

import "smartfarm-api/pkg/database"

type lahanRepository struct{}

func NewLahanRepository() *lahanRepository {
	return &lahanRepository{}
}

type LahanRepository interface {
	// Create(petani *Petani) error
	FindAll() ([]Lahan, error)
	// FindByID(id uint) (*Petani, error)
	// Update(petani *Petani) error
	// Delete(id uint) error
	// 🟢 Tambahkan ini
	FindPaginated(page, limit int) ([]Lahan, int64, error)
}

func (r *lahanRepository) FindAll() ([]Lahan, error) {
	var lahans []Lahan
	err := database.DB.Find(&lahans).Error
	return lahans, err
}

func (r *lahanRepository) FindPaginated(page, limit int) ([]Lahan, int64, error) {
	var lahans []Lahan
	var total int64

	offset := (page - 1) * limit

	if err := database.DB.Model(&Lahan{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.
		Limit(limit).
		Offset(offset).
		Order("id desc").
		Find(&lahans).Error; err != nil {
		return nil, 0, err
	}

	return lahans, total, nil
}
