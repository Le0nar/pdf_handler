package main

import (
	"github.com/Le0nar/pdf_handler/internal/handler"
	"github.com/Le0nar/pdf_handler/internal/service"
)

func main() {
	service := service.NewService()
	handler := handler.NewHandler(service)

	router := handler.InitRouter()

	router.Run()
}
