package services

import "github.com/kryast/crud-11.git/repositories"

type FeedbackService interface {
}

type feedbackService struct {
	repo repositories.FeedbackRepository
}

func NewFeedbackService(repo repositories.FeedbackRepository) FeedbackService {
	return &feedbackService{repo}
}
