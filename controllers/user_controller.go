package controllers

import (
	"net/http"
	"school/model"
	"school/services"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserContoller struct {
	userService services.UserService
}

func (controller UserContoller) Create(c echo.Context) error {
	type payload struct {
		SchoolID  int       `json:"school_id" validate:"required"`
		Name      string    `json:"name" validate:"required"`
		Email     string    `json:"email" validate:"required"`
		Phone     string    `json:"phone" validate:"required"`
		CreatedAt time.Time `json:"created_at"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	result := controller.userService.Create(model.Users{
		Name:      payloadValidator.Name,
		Email:     payloadValidator.Email,
		Phone:     payloadValidator.Phone,
		SchoolID:  payloadValidator.SchoolID,
		CreatedAt: payloadValidator.CreatedAt,
	})

	return c.JSON(http.StatusOK, result)
}

func (controller UserContoller) Update(c echo.Context) error {
	type payload struct {
		SchoolID  int       `json:"school_id" validate:"required"`
		Name      string    `json:"name" validate:"required"`
		Email     string    `json:"email" validate:"required"`
		Phone     string    `json:"phone" validate:"required"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}
	idUser, _ := strconv.Atoi(c.Param("id"))
	result := controller.userService.Update(idUser, model.Users{
		Name:      payloadValidator.Name,
		Email:     payloadValidator.Email,
		Phone:     payloadValidator.Phone,
		SchoolID:  payloadValidator.SchoolID,
		UpdatedAt: payloadValidator.UpdatedAt,
	})

	return c.JSON(http.StatusOK, result)
}

func (controller UserContoller) Delete(c echo.Context) error {

	idUser, _ := strconv.Atoi(c.Param("id"))
	result := controller.userService.Delete(idUser)

	return c.JSON(http.StatusOK, result)
}

func (controller UserContoller) GetAll(c echo.Context) error {
	result := controller.userService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller UserContoller) GetById(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.userService.GetById(idUser)
	return c.JSON(http.StatusOK, result)

}

func NewUserController(db *gorm.DB) UserContoller {
	service := services.NewUserService(db)
	controller := UserContoller{
		userService: service,
	}
	return controller
}
