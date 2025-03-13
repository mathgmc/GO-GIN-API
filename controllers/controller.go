package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mathgmc/GO-GIN-API/models"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}
