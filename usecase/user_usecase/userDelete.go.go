package userusecase



func(u *UserUsecase) DeleteUser(username string) error {
	err := u.UserRepository.DeleteUser(username)
	if err != nil {
		return err
	}
	return nil
}