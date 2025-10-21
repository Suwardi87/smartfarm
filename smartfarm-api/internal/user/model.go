package user

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey;type:char(36)" json:"id"`
	NamaLengkap string `gorm:"type:varchar(150);not null" json:"nama_lengkap"`
	Username    string `gorm:"type:varchar(80);uniqueIndex;not null" json:"username"`
	Password    string `gorm:"type:varchar(255);not null" json:"-"` // disembunyikan dari JSON
	Level       string `gorm:"type:varchar(20);not null" json:"level"` // contoh: petani, nagari, dinas
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	u.Username = strings.TrimSpace(strings.ToLower(u.Username))
	return
}

type LoginRequest struct {
	Username string `json:"username" example:"suwardi"`
	Password string `json:"password" example:"123456"`
}


type RegisterRequest struct {
	NamaLengkap string `json:"nama_lengkap" example:"Suwardi"`
	Username    string `json:"username" example:"suwardi"`
	Password    string `json:"password" example:"123456"`
	Level       string `json:"level" example:"admin"`
}
