package services

import (
	"github.com/kryast/crud-11.git/models"
	"github.com/kryast/crud-11.git/repositories"
)

type FeedbackService interface {
	Create(feedback *models.Feedback) error
	GetAll() ([]models.Feedback, error)
}

type feedbackService struct {
	repo repositories.FeedbackRepository
}

func NewFeedbackService(repo repositories.FeedbackRepository) FeedbackService {
	return &feedbackService{repo}
}

func (fs *feedbackService) Create(feedback *models.Feedback) error {
	return fs.repo.Create(feedback)
}

func (fs *feedbackService) GetAll() ([]models.Feedback, error) {
	return fs.repo.GetAll()
}
