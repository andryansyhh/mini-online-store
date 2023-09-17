package handlers

import (
	"mini-online-store/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OlshopHandler interface {
	SetupHandler(r *gin.Engine)
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	GetProduct(c *gin.Context)
}

type olshopHandler struct {
	userUsecase    usecase.UserUsecase
	productUsecase usecase.ProductUsecase
}

func NewOlshopHandler(userUsecase usecase.UserUsecase, productUsecase usecase.ProductUsecase) OlshopHandler {
	return &olshopHandler{
		userUsecase:    userUsecase,
		productUsecase: productUsecase,
	}
}

func (m *olshopHandler) SetupHandler(r *gin.Engine) {
	public := r.Group("/")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	userGroups := r.Group("/user")
	userGroups.POST("/register", m.RegisterUser)
	userGroups.POST("/login", m.Login)

	productGroup := r.Group("/product")
	productGroup.GET("/all-products", m.GetProduct)
	productGroup.GET("/all-products/:category", m.GetProductByCategory)

}
