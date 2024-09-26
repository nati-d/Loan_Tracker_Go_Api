package userusecase

import "loan_tracker/domain"


func(u *UserUsecase) GetUserByUsernameOrEmail(usernameOrEmail string) (*domain.User, error) {
	user, err := u.UserRepository.GetUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}