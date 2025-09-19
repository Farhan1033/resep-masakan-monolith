package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(30);not null;unique" json:"email"`
	UserName  string    `gorm:"type:varchar(100);not null;unique" json:"user_name"`
	FullName  string    `gorm:"type:varchar(255);not null" json:"full_name"`
	Password  string    `gorm:"type:varchar(100);not null" json:"-"`
	AvatarUrl string    `gorm:"type:text" json:"avatar_url"`
	Bio       string    `gorm:"type:text" json:"bio"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (User) TableName() string {
	return "users"
}
