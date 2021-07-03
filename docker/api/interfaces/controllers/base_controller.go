package controllers

import (
	"myapp/domain"
	"myapp/usecase"
)

type BaseController struct {
	UserInteractor usecase.UserInteractor
}

func (controller *BaseController) findUser(c Context) (user domain.User) {
	uid := c.MustGet("uid").(string)
	user, err := controller.UserInteractor.UserByUid(uid)
	if err != nil {
		c.JSON(500, nil)
		c.Abort()
		return
	}
	return
}
