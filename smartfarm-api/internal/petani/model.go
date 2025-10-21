package petani

import "time"

// Petani merepresentasikan entitas petani di sistem
// Struct untuk database (ORM)
type Petani struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk input API
type PetaniRequest struct {
	Name    string `json:"name" example:"Suwardi"`
	Address string `json:"address" example:"Padang Panjang"`
}


