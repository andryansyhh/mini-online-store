package handlers

import (
	"mini-online-store/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *olshopHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = u.userUsecase.RegisterUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "successfully registered",
	})
}
