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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.userUsecase.RegisterUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully registered"})
}

func (u *olshopHandler) Login(c *gin.Context) {
	var req dto.LoginUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := u.userUsecase.Login(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
