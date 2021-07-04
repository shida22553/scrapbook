package controllers

import (
	// "fmt"
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
	"strconv"
)

type LooseLeafController struct {
	LooseLeafInteractor usecase.LooseLeafInteractor
	UserInteractor      usecase.UserInteractor
	BaseController
}

type LooseLeafPutRequest struct {
	Content string `json:"content"`
}

type LooseLeafBinderIdPutRequest struct {
	BinderID uint `json:"binderId"`
}

func NewLooseLeafController(sqlHandler database.SqlHandler) *LooseLeafController {
	return &LooseLeafController{
		LooseLeafInteractor: usecase.LooseLeafInteractor{
			LooseLeafRepository: &database.LooseLeafRepository{
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

func (controller *LooseLeafController) Create(c Context) {
	user := controller.findUser(c)

	var looseLeaf domain.LooseLeaf
	c.BindJSON(&looseLeaf)
	looseLeaf.UserID = user.ID
	err := controller.LooseLeafInteractor.Add(&looseLeaf)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, looseLeaf)
}

func (controller *LooseLeafController) Update(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	looseLeaf, _ := controller.LooseLeafInteractor.LooseLeafById(user, uintId)
	requestBody := LooseLeafPutRequest{}
	c.BindJSON(&requestBody)
	looseLeaf.Content = requestBody.Content
	looseLeafErr := controller.LooseLeafInteractor.Update(&looseLeaf)
	if looseLeafErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, looseLeaf)
}

func (controller *LooseLeafController) Index(c Context) {
	user := controller.findUser(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	looseLeafs, err := controller.LooseLeafInteractor.LooseLeafs(user, page, pageSize)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, looseLeafs)
}

func (controller *LooseLeafController) Show(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	looseLeaf, err := controller.LooseLeafInteractor.LooseLeafById(user, uintId)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, looseLeaf)
}

func (controller *LooseLeafController) Delete(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	looseLeaf, _ := controller.LooseLeafInteractor.LooseLeafById(user, uintId)
	looseLeafErr := controller.LooseLeafInteractor.Remove(&looseLeaf)
	if looseLeafErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(204, nil)
}

func (controller *LooseLeafController) UpdateBinderID(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	looseLeaf, _ := controller.LooseLeafInteractor.LooseLeafById(user, uintId)
	requestBody := LooseLeafBinderIdPutRequest{}
	c.BindJSON(&requestBody)
	looseLeaf.BinderID = &requestBody.BinderID
	looseLeafErr := controller.LooseLeafInteractor.UpdateBinderID(looseLeaf)
	if looseLeafErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, looseLeaf)
}
