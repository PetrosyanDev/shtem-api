package ports

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	GenerateToken() gin.HandlerFunc
	ValidateToken() gin.HandlerFunc
}
