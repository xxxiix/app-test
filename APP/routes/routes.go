package routes

import (
	"main/controllers"
	"main/logger"
	"main/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	// 用户登录注册相关
	u := r.Group("/user")
	u.POST("/signup", controllers.SignUpHandler)
	u.POST("/sms", controllers.SmsHandler)
	u.POST("/login", controllers.LoginHandler)

	u_login := u.Group("/:user_id")
	u_login.Use(middlewares.JWTAuthMiddleware())

	// 账单处理相关
	b := u_login.Group("/bill")
	b.POST("/add", controllers.AddBillHandler)
	b.POST("/delete", controllers.DeleteBillHandler)
	b.POST("/change", controllers.ChangeBillHandler)
	b.POST("/search-by-day", controllers.SearchBillHandlerByDay)
	b.POST("/search-by-week", controllers.SearchBillHandlerByWeek)
	b.POST("/search-by-month", controllers.SearchBillHandlerByMonth)
	b.POST("/search-by-year", controllers.SearchBillHandlerByYear)
	// 尚未实现
	// b.POST("/add-by-photo", controllers.AddByPhotoBillHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
