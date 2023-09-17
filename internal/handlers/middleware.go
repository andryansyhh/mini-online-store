package handlers

import (
	"mini-online-store/internal/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *olshopHandler) TokenChecker(c *gin.Context) {
	token := c.GetHeader("token")
	reply, err := u.userUsecase.Auth(&dto.AuthRequest{
		Token: token,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.Set("user_uuid", reply.UserUuid)
	c.Set("email", reply.Email)
	c.Set("phone", reply.Phone)
	c.Next()
}
