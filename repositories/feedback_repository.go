package repositories

import (
	"github.com/kryast/crud-11.git/models"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *models.Feedback) error
	GetAll() ([]models.Feedback, error)
	GetByID(id uint) (*models.Feedback, error)
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db}
}

func (fr *feedbackRepository) Create(feedback *models.Feedback) error {
	return fr.db.Create(feedback).Error
}

func (fr *feedbackRepository) GetAll() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	err := fr.db.Find(&feedbacks).Error

	return feedbacks, err
}

func (fr *feedbackRepository) GetByID(id uint) (*models.Feedback, error) {
	var feedback models.Feedback
	err := fr.db.First(&feedback, id).Error

	return &feedback, err
}
