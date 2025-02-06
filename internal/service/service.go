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
	err := CreatePDF(ticket)
	if err != nil {
		return err
	}

	// TODO: try to not save to disk
	// 2) Save to Minio
	fileName := getTicketFileName(ticket.ID.String())
	objectName := getObjectName(ticket.ID.String())

	// TODO: mb use "/" + filename
	_, err = s.S3Storage.FPutObject(bucketName, objectName, fileName, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	// 3) Remove file from disk
	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetTicket(id uuid.UUID) (*minio.Object, string, error) {
	objectName := getObjectName(id.String())

	object, err := s.S3Storage.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}

	return object, objectName, nil
}
