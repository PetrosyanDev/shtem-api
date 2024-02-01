package ports

import "github.com/gin-gonic/gin"

type APIHandler interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindQuestion() gin.HandlerFunc
	FindBajin() gin.HandlerFunc
	GetShtems() gin.HandlerFunc
}
