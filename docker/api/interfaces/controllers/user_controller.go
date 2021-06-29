package controllers

import (
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	id, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, id)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	uid := c.Param("uid")
	user, err := controller.Interactor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, user)
}
