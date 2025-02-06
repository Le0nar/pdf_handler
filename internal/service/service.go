package service

import (
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
	buf, err := CreatePDF(ticket)
	if err != nil {
		return err
	}

	// 2) Save to Minio
	objectName := getObjectName(ticket.ID.String())

	// Загружаем PDF в MinIO
	_, err = s.S3Storage.PutObject(
		bucketName,
		objectName,
		buf,
		int64(buf.Len()),
		minio.PutObjectOptions{ContentType: "application/pdf"},
	)
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
