package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        string    `gorm:"column:id;primaryKey;type:uuid"`
	PartnerId string    `gorm:"column:partner_id"`
	Total     int       `gorm:"column:total"`
	UserName  string    `gorm:"column:user_name"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserModel) TableName() string {
	return "user_tbl"
}

func (u *UserModel) BeforeCreate(ctx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()
	return
}

func (u *UserModel) BeforeUpdate(ctx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *UserModel) GetTotal() int {
	return u.Total
}
