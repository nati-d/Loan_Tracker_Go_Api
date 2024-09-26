package userusecase

import "loan_tracker/domain"


func (u *UserUsecase) GetAllUsers() (*domain.User, error) {
	users, err := u.UserRepository.GetAllUsers()

	

	if err != nil {
		return nil, err
	}
	return users, nil
}