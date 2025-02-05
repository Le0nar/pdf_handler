package service

import (
	"os"

	"github.com/Le0nar/pdf_handler/internal/ticket"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateTicket(ticket ticket.Ticket) error {
	// 1) Create PDF
	CreatePDF(ticket)

	// 2) Save to Minio

	return nil
}

func (s *Service) GetTicket(id uuid.UUID) (*os.File, error) {
	return nil, nil
}
