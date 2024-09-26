package logusecase

import "loan_tracker/domain"

type LogUsecase struct {
	LogRepository domain.LogRepository
}

func NewLogUsecase(lr domain.LogRepository) *LogUsecase {
	return &LogUsecase{
		LogRepository: lr,
	}
}


// func (l *LogUsecase) GetLogs(page,limit string) ([]domain.Log, error) {
// 	return l.LogRepository.GetLogs(page,limit)
// }
