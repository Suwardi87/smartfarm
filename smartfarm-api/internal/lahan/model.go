package lahan

import (
	"smartfarm-api/internal/petani"
	"time"
)

type Lahan struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	PetaniID    uint           `json:"petani_id" gorm:"not null"`
	NamaLahan   string         `json:"nama_lahan" gorm:"type:varchar(100);not null"`
	Lokasi      string         `json:"lokasi" gorm:"type:varchar(255);not null"`
	Luas        float64        `json:"luas" gorm:"type:decimal(10,2)"`
	JenisTanah  string         `json:"jenis_tanah" gorm:"type:varchar(100)"`
	Koordinat   string         `json:"koordinat,omitempty" gorm:"type:varchar(100)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Petani petani.Petani `json:"petani" gorm:"foreignKey:PetaniID"`
}
