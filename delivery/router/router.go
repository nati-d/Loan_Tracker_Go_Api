package router

import (
	loancontroller "loan_tracker/delivery/controllers/loan_controller"
	// "loan_tracker/delivery/controllers/log_controller"
	usercontroller "loan_tracker/delivery/controllers/user_controller"
	"loan_tracker/delivery/middleware"
	"loan_tracker/repository"
	loanusecase "loan_tracker/usecase/loan_usecase"
	// logusecase "loan_tracker/usecase/log_usecase"
	userusecase "loan_tracker/usecase/user_usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func getUserController(db *mongo.Database) *usercontroller.UserController {
	UserRepository := repository.NewUserRepository(db)
	UserUsecase := userusecase.NewUserUsecase(UserRepository)
	UserController := usercontroller.NewUserController(UserUsecase)
	return UserController	
}

func getLoanController(db *mongo.Database) *loancontroller.LoanController {
	LoanRepository := repository.NewLoanRepository(db)
	UserRepository := repository.NewUserRepository(db)
	LoanUsecase := loanusecase.NewLoanUsecase(LoanRepository)
	UserUsecase := userusecase.NewUserUsecase(UserRepository)
	LoanController := loancontroller.NewLoanController(LoanUsecase, UserUsecase)
	return LoanController
}

// func getLogController(db *mongo.Database) *log_controller.LogController {
// 	LogRepository := repository.NewLogRepository(db)
// 	LogUsecase := logusecase.NewLogUsecase(LogRepository)
// 	UserUsecase := userusecase.NewUserUsecase(repository.NewUserRepository(db))
// 	LogController := log_controller.NewLogController(LogUsecase, UserUsecase)
// 	return LogController
// }

func publicRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.POST("/register", userController.RegisterUser)
	router.GET("/users/verify-email", userController.VerifyUser)
	router.POST("/login", userController.Login)
	router.POST("/password-reset-request", userController.PasswordResetRequest)
}

func privateRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.GET("/users/profile", middleware.AuthMiddleware("access"), userController.GetUserProfile)
	router.GET("/users", middleware.AuthMiddleware("access"), userController.GetAllUsers)
}

func protectedRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.POST(
		"/tokens/refresh",
		middleware.AuthMiddleware("refresh"),
		userController.RefreshToken,
	)
}

func privateLoanRouter(router *gin.Engine, loanController *loancontroller.LoanController) {
	router.POST("/loans", middleware.AuthMiddleware("access"), loanController.ApplyLoan)
	router.GET("/admin/loans", middleware.AuthMiddleware("access"), loanController.GetAllLoans)
	router.GET("/myloans", middleware.AuthMiddleware("access"), loanController.GetMyLoans)
	router.PATCH("/admin/approve/:id", middleware.AuthMiddleware("access"), loanController.ApproveLoan)
}





func SetupRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()

	db := client.Database("loan_tracker")
	UserController := getUserController(db)

	//create public routes
	publicRouter(router, UserController)
	//create protected routes
	protectedRouter(router, UserController)

	//create private routes
	privateRouter(router, UserController)

	//create private routes for loans
	privateLoanRouter(router, getLoanController(db))




	return router


	
}