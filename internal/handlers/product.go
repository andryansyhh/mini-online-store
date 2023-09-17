package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *olshopHandler) GetProduct(c *gin.Context) {
	res, err := o.productUsecase.GetAllProduct()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (o *olshopHandler) GetProductByCategory(c *gin.Context) {
	category := c.Param("category")
	res, err := o.productUsecase.GetProductByCategory(category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
