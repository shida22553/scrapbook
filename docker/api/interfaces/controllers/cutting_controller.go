package controllers

import (
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
	"strconv"
)

type CuttingController struct {
	CuttingInteractor usecase.CuttingInteractor
	UserInteractor    usecase.UserInteractor
	BaseController
}

type CuttingPutRequest struct {
	Note string `json:"note"`
}

func NewCuttingController(sqlHandler database.SqlHandler) *CuttingController {
	return &CuttingController{
		CuttingInteractor: usecase.CuttingInteractor{
			CuttingRepository: &database.CuttingRepository{
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

func (controller *CuttingController) Create(c Context) {
	user := controller.findUser(c)

	var cutting domain.Cutting
	c.BindJSON(&cutting)
	cutting.UserID = user.ID
	err := controller.CuttingInteractor.Add(&cutting)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, cutting)
}

func (controller *CuttingController) Update(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	cutting, _ := controller.CuttingInteractor.CuttingById(user, uintId)
	requestBody := CuttingPutRequest{}
	c.BindJSON(&requestBody)
	cutting.Note = requestBody.Note
	cuttingErr := controller.CuttingInteractor.Update(&cutting)
	if cuttingErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, cutting)
}

func (controller *CuttingController) Index(c Context) {
	user := controller.findUser(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	cuttings, err := controller.CuttingInteractor.Cuttings(user, page, pageSize)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, cuttings)
}

func (controller *CuttingController) Show(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	cutting, err := controller.CuttingInteractor.CuttingById(user, uintId)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, cutting)
}

func (controller *CuttingController) Delete(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	cutting, _ := controller.CuttingInteractor.CuttingById(user, uintId)
	cuttingErr := controller.CuttingInteractor.Remove(&cutting)
	if cuttingErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(204, nil)
}
