package controllers

import (
	"fmt"
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
	BaseController
}

type UserPutRequest struct {
	Name string `json:"name"`
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
		BaseController: BaseController{
			UserInteractor: usecase.UserInteractor{
				UserRepository: &database.UserRepository{
					SqlHandler: sqlHandler,
				},
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	fmt.Printf("create: \n")
	requestBody := UserPutRequest{}
	c.BindJSON(&requestBody)
	name := requestBody.Name
	// fmt.Printf("must uid: %v\n", uid)
	// fmt.Printf("uid: %v\n", c.Param("uid"))
	fmt.Printf("name: %v\n", name)
	u := domain.User{
		Name: name,
		Uid:  c.MustGet("uid").(string),
	}
	id, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, id)
}

func (controller *UserController) Update(c Context) {
	user := controller.findUser(c)
	requestBody := UserPutRequest{}
	c.BindJSON(&requestBody)
	user.Name = requestBody.Name
	userErr := controller.Interactor.Update(&user)
	if userErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, user)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) ShowCurrentUser(c Context) {
	user := controller.findUser(c)
	c.JSON(200, user)
}
