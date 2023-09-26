package services

import "quoation-backend/models"

type QuoteService interface {
	CreateQuote(*models.Quote) error
	GetQuote(*string) (*models.Quote,error)
	GetAll()([]*models.Quote,error)
	// UpdateQuote(*string) (*models.Quote,error)
	UpdateQuote(*models.Quote) error
	DeleteQuote(*string) error

}
