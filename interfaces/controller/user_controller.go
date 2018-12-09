package controller

import (
	"strconv"

	"github.com/yoshikawataiki/simple-api/domain"
	"github.com/yoshikawataiki/simple-api/interfaces/database"
	"github.com/yoshikawataiki/simple-api/usecase"
)

// UserController model
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController return the UserController
func NewUserController(SQLHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the user
func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, user)
}

// Index is to index the users
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, users)
}

// Show is to show the user
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	user, err := controller.Interactor.UserByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
