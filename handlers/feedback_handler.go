package handlers

import (
	"github.com/kryast/crud-11.git/services"
)

type FeedbackHandler struct {
	service services.FeedbackService
}

func NewFeedbackService(service services.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service}
}
