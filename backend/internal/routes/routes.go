package routes

import (
	"github.com/KeishiIrisa/backend-go-template/internal/controllers"
	"github.com/KeishiIrisa/backend-go-template/internal/middleware"

	"github.com/gin-gonic/gin"

	docs "github.com/KeishiIrisa/backend-go-template/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes sets up the application routes
func RegisterRoutes(router *gin.Engine) {
	// Swagger routes
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Hello world routes
	router.GET("/", controllers.GetHelloWorld)

	router.POST("/sign-up", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected routes (Require authentication)
	api := router.Group("/users")
	api.Use(middleware.JWTAuthMiddleware())
	{
		// CRUD Routes for users
		api.GET("/me", controllers.GetLoggedInUser)
		api.GET("/:id", controllers.GetUserById)
		api.PUT("/:id", controllers.UpdateUser)
	}
}
