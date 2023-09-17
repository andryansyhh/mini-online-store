package handlers

import (
	"mini-online-store/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *olshopHandler) AddToCart(c *gin.Context) {
	userUuid := c.GetString("user_uuid")
	var req dto.AddToCart
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = o.shoppingCartUsecase.AddToCart(&dto.AddToCart{
		UserUuid:    userUuid,
		ProductUuid: req.ProductUuid,
		Qty:         req.Qty,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success product to cart"})
}

func (o *olshopHandler) ListItemInCart(c *gin.Context) {
	userUuid := c.GetString("user_uuid")
	res, err := o.shoppingCartUsecase.ListItemInCart(userUuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (o *olshopHandler) DeleteCart(c *gin.Context) {
	userUuid := c.GetString("user_uuid")
	productUuid := c.Param("cart_uuid")
	err := o.shoppingCartUsecase.DeleteCartItem(productUuid, userUuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted cart"})
}