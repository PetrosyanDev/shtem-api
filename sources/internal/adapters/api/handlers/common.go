package handlers

import "github.com/gin-gonic/gin"

type QuestionsHandler interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindQuestion() gin.HandlerFunc
	FindBajin() gin.HandlerFunc
}
