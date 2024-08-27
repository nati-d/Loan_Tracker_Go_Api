package usercontroller

import "loan_tracker/domain"


type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(us domain.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: us,
	}
}