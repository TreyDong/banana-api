package controller

import (
	"net/http"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/gin-gonic/gin"
)

// CheckUserExists checks if a user with given email exists
// This is a lightweight endpoint for frontend to decide login vs register flow
func CheckUserExists(c *gin.Context) {
	email := c.Query("email")
	if err := common.Validate.Var(email, "required,email"); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的邮箱地址",
		})
		return
	}

	exists := model.IsEmailAlreadyTaken(email)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"exists":  exists,
	})
}
