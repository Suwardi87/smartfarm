package rencana

import "time"

type RencanaTanam struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PetaniID        uint      `json:"petani_id" gorm:"not null"`
	LahanID         uint      `json:"lahan_id" gorm:"not null"`
	TanamanID       uint      `json:"tanaman_id" gorm:"not null"`
	TanggalTanam    time.Time `json:"tanggal_tanam" gorm:"type:date"`
	PerkiraanPanen  time.Time `json:"perkiraan_panen" gorm:"type:date"`
	Status          string    `json:"status" gorm:"type:enum('dijadwalkan','berlangsung','selesai');default:'dijadwalkan'"`
	Keterangan      string    `json:"keterangan" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// Relasi
	// Petani  Petani  `json:"petani" gorm:"foreignKey:PetaniID"`
	// Lahan   Lahan   `json:"lahan" gorm:"foreignKey:LahanID"`
	// Tanaman Tanaman `json:"tanaman" gorm:"foreignKey:TanamanID"`
}
