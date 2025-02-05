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
	return nil
}

func (s *Service) GetTicket(id uuid.UUID) (*os.File, error) {
	return nil, nil
}
