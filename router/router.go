package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gurbanli/ProductManager/controller"
	"github.com/gurbanli/ProductManager/middleware"
)

func NewRouter(env string) *gin.Engine {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			loginController := controller.LoginController{}
			authGroup.POST("/login", loginController.Login)

			registerController := controller.RegisterController{}
			authGroup.Use(middleware.SessionMiddleware())
			authGroup.POST("/register", registerController.Register)
		}
		productGroup := v1.Group("product")
		{
			productGroup.Use(middleware.SessionMiddleware())

			productController := controller.ProductController{}
			productGroup.POST("/", productController.CreateProduct)
			productGroup.PUT("/", productController.UpdateProduct)
			productGroup.GET("/:id", productController.GetProduct)
			productGroup.GET("/", productController.GetProducts)
			productGroup.DELETE("/:id", productController.DeleteProduct)

		}
	}
	return router
}
