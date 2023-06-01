package router

import (
	"gin_test/bulletin_board/controller"
	interfaces "gin_test/bulletin_board/dao/user"
	"gin_test/bulletin_board/initializers"
	middlewares "gin_test/bulletin_board/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController, userController *controller.UsersController, postController *controller.PostController, usersInterface interfaces.UsersInterface) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static/")
	apiRouter := router.Group("/")
	apiRouter.GET("/password_reset/:token/edit", userController.ResetPasswordForm)
	apiRouter.POST("/resetPassword/:token", userController.ResetPassword)
	AuthRouter(apiRouter, authController)
	UsersRouter(apiRouter, usersInterface, userController)
	TagsRouter(apiRouter, postController, usersInterface)
	AuthForm(apiRouter, authController)
	return router
}

// Auth Form Routes
func AuthForm(router *gin.RouterGroup, authController *controller.AuthController) {
	authForm := router.Group("")
	{
		authForm.GET("/register", authController.RegisterForm)
		authForm.GET("/login", authController.LoginForm)
		authForm.GET("/forgetpassword", authController.ForgetPasswordForm)
		// authForm.GET("/password_reset/:token/edit", authController.ResetPasswordForm)
		// authForm.POST("/resetPassword/:token", authController.ResetPassword)
		authForm.POST("/forgetpassword/sendmail", authController.ForgetPassword)
	}
}

// Auth Router
func AuthRouter(router *gin.RouterGroup, authController *controller.AuthController) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/logout", authController.Logout)
	}
}

// User Router
func UsersRouter(router *gin.RouterGroup, usersInterface interfaces.UsersInterface, usersController *controller.UsersController) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", middlewares.IsAuth(usersInterface), func(ctx *gin.Context) {
			userRole := ctx.GetString("UserRole")
			usersController.GetUsers(ctx, userRole)
		})
		userRouter.GET("/profile", middlewares.IsAuth(usersInterface), usersController.ProfileForm)
		userRouter.GET("/changepassword", middlewares.IsAuth(usersInterface), usersController.ChangePasswordForm)
		router.POST("/users/changepasswords", middlewares.IsAuth(usersInterface), usersController.UpdatePassword)
		userRouter.GET("/create", middlewares.IsAuth(usersInterface), usersController.CreateUser)
		userRouter.GET("/update/:userId", middlewares.IsAuth(usersInterface), usersController.UpdateForm)
		userRouter.DELETE("/:userId", middlewares.IsAuth(usersInterface), usersController.Delete)
		userRouter.POST("/:userId", middlewares.IsAuth(usersInterface), usersController.Update)
	}
}

// PostRouter
func TagsRouter(router *gin.RouterGroup, PostsController *controller.PostController, userInterface interfaces.UsersInterface) {
	tagRouter := router.Group("/posts")
	{
		tagRouter.GET("/create", PostsController.CreateForm)
		router.POST("/posts/download", func(c *gin.Context) {
			initializers.ConnectDatabase()
			PostsController.DownloadPosts(c, initializers.DB)
		})
		tagRouter.GET("/upload", middlewares.IsAuth(userInterface), PostsController.UploadForm)
		router.POST("/posts/upload", func(c *gin.Context) {
			initializers.ConnectDatabase()
			PostsController.UploadPosts(c, initializers.DB)
		})

		// tagRouter.GET("/createConfirm", PostsController.CreateConfirmForm)
		tagRouter.GET("/update/:tagId", middlewares.IsAuth(userInterface), PostsController.UpdateForm)
		tagRouter.GET("", PostsController.FindAll)
		tagRouter.GET("/:tagId", middlewares.IsAuth(userInterface), PostsController.FindById)
		tagRouter.POST("", middlewares.IsAuth(userInterface), func(ctx *gin.Context) {
			userId := ctx.GetInt("Id")
			PostsController.Create(ctx, userId)
		})
		tagRouter.POST("/:tagId", middlewares.IsAuth(userInterface), func(ctx *gin.Context) {
			userId := ctx.GetInt("Id")
			PostsController.Update(ctx, userId)
		})
		tagRouter.DELETE("/:tagId", middlewares.IsAuth(userInterface), PostsController.Delete)
	}
	router.GET("/", PostsController.FindAll)
}
