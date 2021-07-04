package controllers

import (
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
	"strconv"
)

type BinderController struct {
	BinderInteractor usecase.BinderInteractor
	UserInteractor   usecase.UserInteractor
	BaseController
}

type BinderPutRequest struct {
	Name string `json:"name"`
}

func NewBinderController(sqlHandler database.SqlHandler) *BinderController {
	return &BinderController{
		BinderInteractor: usecase.BinderInteractor{
			BinderRepository: &database.BinderRepository{
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

func (controller *BinderController) Create(c Context) {
	user := controller.findUser(c)

	var binder domain.Binder
	c.BindJSON(&binder)
	binder.UserID = user.ID
	err := controller.BinderInteractor.Add(&binder)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, binder)
}

func (controller *BinderController) Update(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	binder, _ := controller.BinderInteractor.BinderById(user, uintId)
	requestBody := BinderPutRequest{}
	c.BindJSON(&requestBody)
	binder.Name = requestBody.Name
	binderErr := controller.BinderInteractor.Update(&binder)
	if binderErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, binder)
}

func (controller *BinderController) Index(c Context) {
	user := controller.findUser(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	binders, err := controller.BinderInteractor.Binders(user, page, pageSize)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, binders)
}

func (controller *BinderController) Show(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	binder, err := controller.BinderInteractor.BinderById(user, uintId)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, binder)
}

func (controller *BinderController) Delete(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	binder, _ := controller.BinderInteractor.BinderById(user, uintId)
	binderErr := controller.BinderInteractor.Remove(&binder)
	if binderErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(204, nil)
}
