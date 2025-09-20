package ingrediententity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ingredient struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	IsActive  bool      `gorm:"type:boolean;default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (i *Ingredient) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}

	return nil
}

func (Ingredient) TableName() string {
	return "ingredient"
}
