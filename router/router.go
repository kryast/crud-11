package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-11.git/handlers"
	"github.com/kryast/crud-11.git/repositories"
	"github.com/kryast/crud-11.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewFeedbackRepository(db)
	svc := services.NewFeedbackService(repo)
	h := handlers.NewFeedbackService(svc)

	r.POST("/feedback", h.Create)
	r.GET("/feedback", h.GetAll)
	r.GET("/feedback/:id", h.GetByID)

	return r
}
