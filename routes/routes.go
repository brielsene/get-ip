package routes

import (
	"get-ip/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/nubank-valid-pix", controllers.GetIp)
	r.Run(":8000")
}
