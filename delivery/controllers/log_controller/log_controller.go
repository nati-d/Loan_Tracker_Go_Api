package log_controller

import (
	"loan_tracker/domain"
	// "net/http"

	// "github.com/gin-gonic/gin"
)

type LogController struct {
	logUsecase domain.LogUsecase
	userUsecase domain.UserUsecase
}

func NewLogController(logUsecase domain.LogUsecase,userUsecase domain.UserUsecase) *LogController {
	return &LogController{
		logUsecase: logUsecase,
		userUsecase: userUsecase,
	}
}

// GetLogs handles the request to retrieve system logs
// func (lc *LogController) GetLogs(c *gin.Context) {
// 	claims, ok := c.MustGet("claims").(*domain.LoginClaims)
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}
	
// 	// Check if the user is an admin
// 	_, err := lc.userUsecase.GetUserByUsernameOrEmail(claims.Username)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}	

// 	// Extract pagination parameters
// 	page := c.Query("page")
// 	limit := c.Query("limit")

// 	// Get logs from usecase
// 	logs, count, err := lc.logUsecase.GetLogs(page, limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Respond with logs and total count
// 	c.JSON(http.StatusOK, gin.H{"logs": logs, "totalCount": count})
// }
