package handlers

import (
	"net/http"
	"strconv"

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

func (fh *FeedbackHandler) GetAll(c *gin.Context) {
	feedback, err := fh.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, feedback)

}

func (fh *FeedbackHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	feedback, err := fh.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, feedback)
}

func (fh *FeedbackHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	feedback.ID = uint(id)
	if err := fh.service.Update(&feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

func (fh *FeedbackHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := fh.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Deleted"})
}
