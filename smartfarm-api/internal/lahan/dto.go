package lahan

import "time"

// 🟩 Request DTO
// Digunakan untuk input dari client (misalnya dari body JSON POST/PUT).
type CreateLahanRequest struct {
	NamaLahan   string `json:"nama_lahan" validate:"required"`
	Lokasi      string `json:"lokasi" validate:"required"`
	Luas        float64 `json:"luas" validate:"required"`
	JenisTanah  string `json:"jenis_tanah" validate:"required"`
	Koordinat   string `json:"koordinat"`
}

// 🟩 Request DTO
// Digunakan untuk input dari client (misalnya dari body JSON POST/PUT).
type UpdateLahanRequest struct {
	NamaLahan   string `json:"nama_lahan" validate:"required"`
	Lokasi      string `json:"lokasi" validate:"required"`
	Luas        float64 `json:"luas" validate:"required"`
	JenisTanah  string `json:"jenis_tanah" validate:"required"`
	Koordinat   string `json:"koordinat"`
}

// 🟩 Response DTO
// Digunakan untuk output ke client (hasil GET/POST/PUT/DELETE).
type LahanResponse struct {
	ID          uint      `json:"id"`
	PetaniID    uint      `json:"petani_id"`
	NamaPetani  string    `json:"nama_petani"` // ✅ ditambahkan agar mapping tidak error
	NamaLahan   string    `json:"nama_lahan"`
	Lokasi      string    `json:"lokasi"`
	Luas        float64   `json:"luas"`
	JenisTanah  string    `json:"jenis_tanah"`
	Koordinat   string    `json:"koordinat"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
