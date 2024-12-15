package routes

import (
	"school/model"

	"gorm.io/gorm"
)

type dbUser struct {
	Conn *gorm.DB
}

// Create implements UserRoute.
func (db *dbUser) Create(user model.Users) error {
	return db.Conn.Create(user).Error
}

// Delete implements UserRoute.
func (db *dbUser) Delete(idUser int) error {
	return db.Conn.Delete(idUser).Error
}

// GetAll implements UserRoute.
func (db *dbUser) GetAll() ([]model.Users, error) {
	var data []model.Users
	result := db.Conn.Find(&data)

	return data, result.Error
}

// GetById implements UserRoute.
func (db *dbUser) GetById(idUser int) (model.Users, error) {
	var data model.Users
	result := db.Conn.Where("id", idUser).First(&data)
	return data, result.Error
}

// Update implements UserRoute.
func (db *dbUser) Update(idUser int, user model.Users) error {
	return db.Conn.Where("id", idUser).Updates(user).Error
}

type UserRoute interface {
	Create(user model.Users) error
	Update(idUser int, user model.Users) error
	Delete(idUser int) error
	GetById(idUser int) (model.Users, error)
	GetAll() ([]model.Users, error)
}

func NewUserRoute(Conn *gorm.DB) UserRoute {
	return &dbUser{Conn: Conn}
}
