package controllers

import (
	"github.com/antibomberman/junior_test/internal/models"
	"github.com/antibomberman/junior_test/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (uc *UserController) Index(c *gin.Context) {
	users, err := uc.userService.All()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, users)
}
func (uc *UserController) Show(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.userService.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "not found"})
		return
	}

	c.JSON(200, user)
}
func (uc *UserController) Create(c *gin.Context) {
	validate := validator.New()

	data := models.UserCreate{
		Name:        c.Param("name"),
		Surname:     c.Param("surname"),
		Patronymic:  c.Param("patronymic"),
		Gender:      c.Param("gender"),
		Age:         c.Param("age"),
		Nationality: c.Param("nationality"),
	}
	err := validate.Struct(data)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, err := uc.userService.Create(data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, user)
}
func (uc *UserController) Update(c *gin.Context) {
	validate := validator.New()

	id := c.Param("id")

	data := models.UserUpdate{
		Name:        c.Param("name"),
		Surname:     c.Param("surname"),
		Patronymic:  c.Param("patronymic"),
		Gender:      c.Param("gender"),
		Age:         c.Param("age"),
		Nationality: c.Param("nationality"),
	}
	err := validate.Struct(data)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, err := uc.userService.Update(id, data)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, user)
}
func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := uc.userService.Delete(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "deleted"})
}
