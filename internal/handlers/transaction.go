package handlers

import (
	"mini-online-store/internal/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *olshopHandler) BuyProduct(c *gin.Context) {
	userUuid := c.GetString("user_uuid")
	var req models.Transaction
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserUuid = userUuid
	res, err := o.transactionUsecase.CreateTransaction(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
