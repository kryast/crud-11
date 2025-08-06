package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-11.git/models"
	"github.com/kryast/crud-11.git/services"
)

type FeedbackHandler struct {
	service services.FeedbackService
}

func NewFeedbackService(service services.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service}
}

func (fh *FeedbackHandler) Create(c *gin.Context) {
	var feedback models.Feedback

	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := fh.service.Create(&feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, feedback)
}
