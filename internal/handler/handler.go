package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Le0nar/pdf_handler/internal/ticket"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
)

type service interface {
	CreateTicket(ticket ticket.Ticket) error
	GetTicket(id uuid.UUID) (*minio.Object, string, error)
}

type Handler struct {
	service service
	// TODO: create interface for validator
	validator *validator.Validate
}

func NewHandler(s service) *Handler {
	return &Handler{
		service:   s,
		validator: validator.New(),
	}
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var dto ticket.Ticket

	// Прочитаем JSON в структуру
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Валидируем структуру
	if err := h.validator.Struct(dto); err != nil {
		// Ошибки валидации
		var validationErrors []string

		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, e.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	err := h.service.CreateTicket(dto)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetTicket(c *gin.Context) {
	// TODO: придумать как корерктно отдать файл

	stringedId := c.Param("id")

	id, err := uuid.Parse(stringedId)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	object, objectName, err := h.service.GetTicket(id)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer object.Close()

	// Устанавливаем правильные HTTP-заголовки для PDF
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", objectName))

	// Отправляем PDF файл как поток, используя io.Copy
	_, err = io.Copy(c.Writer, object)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при отправке файла пользователю",
		})
		log.Fatalln(err)
		return
	}

}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		ticketGroup := api.Group("/ticket")
		{
			ticketGroup.POST("/file", h.CreateTicket)
			ticketGroup.GET("/:id/file", h.GetTicket)
		}
	}

	return r
}
