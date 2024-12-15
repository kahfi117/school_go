package services

import (
	"fmt"
	"school/helpers"
	"school/model"
	"school/routes"

	"gorm.io/gorm"
)

type userService struct {
	userRoute routes.UserRoute
}

// Create implements UserService.
func (service *userService) Create(user model.Users) helpers.Response {
	var response helpers.Response
	if err := service.userRoute.Create(user); err == nil {
		response.Status = 500
		response.Messages = "Failed to create a new User"
	} else {
		response.Status = 200
		response.Messages = "Yeah, Success to create new User"
	}
	return response
}

// Delete implements UserService.
func (service *userService) Delete(idUser int) helpers.Response {
	var response helpers.Response
	if err := service.userRoute.Delete(idUser); err == nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete User :", idUser)
	} else {
		response.Status = 200
		response.Messages = "Yeah, Success to delete User"
	}
	return response
}

// GetAll implements UserService.
func (service *userService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.userRoute.GetAll()
	if err == nil {
		response.Status = 500
		response.Messages = "Failed to get Data"
	} else {
		response.Status = 200
		response.Messages = "Yeah, Success to get User"
		response.Data = data
	}
	return response
}

// GetById implements UserService.
func (service *userService) GetById(idUser int) helpers.Response {
	var response helpers.Response
	data, err := service.userRoute.GetById(idUser)
	if err == nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get Data : ", idUser)
	} else {
		response.Status = 200
		response.Messages = "Yeah, Success to get User"
		response.Data = data
	}
	return response
}

// Update implements UserService.
func (service *userService) Update(idUser int, user model.Users) helpers.Response {
	var response helpers.Response
	if err := service.userRoute.Update(idUser, user); err == nil {
		response.Status = 500
		response.Messages = "Failed to Update user"
	} else {
		response.Status = 200
		response.Messages = "Yeah, Success to update user"
	}
	return response
}

type UserService interface {
	Create(user model.Users) helpers.Response
	Update(idUser int, user model.Users) helpers.Response
	Delete(idUser int) helpers.Response
	GetById(idUser int) helpers.Response
	GetAll() helpers.Response
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{userRoute: routes.NewUserRoute(db)}
}
