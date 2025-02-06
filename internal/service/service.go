package service

import (
	"os"

	"github.com/Le0nar/pdf_handler/internal/ticket"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
)

type Service struct {
	S3Storage *minio.Client
}

func NewService() *Service {
	return &Service{
		S3Storage: initS3Storage(),
	}
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
