package petani

import "time"

//
// 🟩 Request DTO
// Digunakan untuk input dari client (misalnya dari body JSON POST/PUT).
//
type CreatePetaniRequest struct {
	Name    string `json:"name" validate:"required,min=3"`
	Address string `json:"address" validate:"required,min=5"`
}

type UpdatePetaniRequest struct {
	Name    string `json:"name" validate:"required,min=3"`
	Address string `json:"address" validate:"required,min=5"`
}

//
// 🟦 Response DTO
// Digunakan untuk output ke client (hasil GET/POST/PUT/DELETE).
//
type PetaniResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
