package handler

import (
	"mygram/app/middleware"
	"mygram/app/services"

	"github.com/gin-gonic/gin"

	_ "mygram/docs"

	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @version 1.0
// @description This is an API for MyGram APP. To use all of the services, please login first and get the token.
// @description After that, click the "Authorize" button at the right and a pop-up window will appear. type "Bearer <your-token>". Example: Bearer eyJhbGciOiJIUzI1...

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  togi.mare@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @host mygram-production-2f89.up.railway.app
// @BasePath /
// @swagg.NoModels

func StartServer() *gin.Engine {
	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	router.Use(cors.New(config))

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", services.MyGramUserRegister)
		userRouter.POST("/login", services.MyGramUserLogin)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	photoRouter := router.Group("/photos")
	{
		//check apakah sudah login atau belum
		photoRouter.GET("/", services.MyGramGetAllPhoto) //get all photo
		photoRouter.GET("/:id", services.MyGramGetPhoto) //get photo by id
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", services.MyGramCreatePhoto) //create photo
		photoRouter.PUT("/:id", middleware.Authorization("photo"), services.MyGramUpdatePhoto)
		photoRouter.DELETE("/:id", middleware.Authorization("photo"), services.MyGramDeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.GET("/", services.MyGramGetAllComment)
		commentRouter.GET("/:id", services.MyGramGetComment)
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/:photoID", services.MyGramCreateComment)
		commentRouter.PUT("/:id", middleware.Authorization("comment"), services.MyGramUpdateComment)
		commentRouter.DELETE("/:id", middleware.Authorization("comment"), services.MyGramDeleteComment)
	}

	socialMediaRouter := router.Group("/social-media")
	{
		socialMediaRouter.GET("/", services.MyGramGetAllSocialMedia)
		socialMediaRouter.GET("/:id", services.MyGramGetSocialMedia)
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", services.MyGramCreateSocialMedia)
		socialMediaRouter.PUT("/:id", middleware.Authorization("socialMedia"), services.MyGramUpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middleware.Authorization("socialMedia"), services.MyGramDeleteSocialMedia)
	}

	return router
}
