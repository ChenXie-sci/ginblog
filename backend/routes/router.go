package routes

import (
	"ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		//User module

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//category module
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//article module
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		//upload file
		auth.POST("upload", v1.UpLoad)
		// update user profile
		auth.PUT("profile/:id", v1.UpdateProfile)
	}

	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)
		router.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)
}
