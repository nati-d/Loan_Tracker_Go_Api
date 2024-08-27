package router

import (
	usercontroller "loan_tracker/delivery/controllers/user_controller"
	"loan_tracker/delivery/middleware"
	"loan_tracker/repository/userrepository"
	userusecase "loan_tracker/usecase/user_usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func getUserController(db *mongo.Database) *usercontroller.UserController {
	UserRepository := userrepository.NewUserRepository(db)
	UserUsecase := userusecase.NewUserUsecase(UserRepository)
	UserController := usercontroller.NewUserController(UserUsecase)
	return UserController
	
}

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




	return router


	
}