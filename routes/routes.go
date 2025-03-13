package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathgmc/GO-GIN-API/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.GetAllStudents)
	r.Run()
}
