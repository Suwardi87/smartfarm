package petani

import "time"

// Petani merepresentasikan entitas petani di sistem
// Struct untuk database (ORM)
type Petani struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Alamat   string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk input API
type PetaniRequest struct {
	Nama    string `json:"nama" example:"Suwardi"`
	Alamat string `json:"alamat" example:"Padang Panjang"`
}


