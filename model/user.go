package model

import "time"

type Users struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	SchoolID  int       `json:"school_id" gorm:"column:school_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

func (Users) TableName() string {
	return "users"
}
