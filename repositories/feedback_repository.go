package repositories

import "gorm.io/gorm"

type FeedbackRepository interface {
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db}
}
