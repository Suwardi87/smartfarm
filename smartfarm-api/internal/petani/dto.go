package petani

import "time"

//
// 🟩 Request DTO
// Digunakan untuk input dari client (misalnya dari body JSON POST/PUT).
//
type CreatePetaniRequest struct {
	Nama    string `json:"nama" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=5"`
}

type UpdatePetaniRequest struct {
	Nama    string `json:"nama" validate:"required,min=3"`
	Alamat string `json:"alamat" validate:"required,min=5"`
}

//
// 🟦 Response DTO
// Digunakan untuk output ke client (hasil GET/POST/PUT/DELETE).
//
type PetaniResponse struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Alamat   string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
