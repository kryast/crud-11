package repositories

import (
	"github.com/kryast/crud-11.git/models"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *models.Feedback) error
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
