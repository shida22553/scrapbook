package controllers

import (
	"myapp/domain"
	"myapp/interfaces/database"
	"myapp/usecase"
	"strconv"
)

type BookController struct {
	BookInteractor usecase.BookInteractor
	UserInteractor usecase.UserInteractor
	BaseController
}

type BookPutRequest struct {
	Name string `json:"name"`
}

func NewBookController(sqlHandler database.SqlHandler) *BookController {
	return &BookController{
		BookInteractor: usecase.BookInteractor{
			BookRepository: &database.BookRepository{
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

func (controller *BookController) Create(c Context) {
	user := controller.findUser(c)

	var book domain.Book
	c.BindJSON(&book)
	book.UserID = user.ID
	err := controller.BookInteractor.Add(&book)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, book)
}

func (controller *BookController) Update(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	book, _ := controller.BookInteractor.BookById(user, uintId)
	requestBody := BookPutRequest{}
	c.BindJSON(&requestBody)
	book.Name = requestBody.Name
	bookErr := controller.BookInteractor.Update(&book)
	if bookErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(201, book)
}

func (controller *BookController) Index(c Context) {
	user := controller.findUser(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	books, err := controller.BookInteractor.Books(user, page, pageSize)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, books)
}

func (controller *BookController) Show(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	book, err := controller.BookInteractor.BookById(user, uintId)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, book)
}

func (controller *BookController) Delete(c Context) {
	user := controller.findUser(c)

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	uintId := uint(id)
	book, _ := controller.BookInteractor.BookById(user, uintId)
	bookErr := controller.BookInteractor.Remove(&book)
	if bookErr != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(204, nil)
}
