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
		UserInteractor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CuttingController) Create(c Context) {
	uid := c.MustGet("uid").(string)
	user, err := controller.UserInteractor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	var cutting domain.Cutting
	c.BindJSON(&cutting)
	cutting.UserID = user.ID
	id, err := controller.CuttingInteractor.Add(cutting)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, id)
}

func (controller *CuttingController) Update(c Context) {
	uid := c.MustGet("uid").(string)
	user, err := controller.UserInteractor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	cutting, _ := controller.CuttingInteractor.CuttingById(user, id)
	requestBody := CuttingPutRequest{}
	c.BindJSON(&requestBody)
	cutting.Note = requestBody.Note
	cuttingId, cuttingErr := controller.CuttingInteractor.Update(cutting)
	if cuttingErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, cuttingId)
}

func (controller *CuttingController) Index(c Context) {
	uid := c.MustGet("uid").(string)
	user, err := controller.UserInteractor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		return
	}
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
	uid := c.MustGet("uid").(string)
	user, err := controller.UserInteractor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	cutting, err := controller.CuttingInteractor.CuttingById(user, id)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, cutting)
}
