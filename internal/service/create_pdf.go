package service

import (
	"fmt"

	"github.com/Le0nar/pdf_handler/internal/ticket"
	"github.com/jung-kurt/gofpdf"
)

// Функция для создания PDF на основе данных Ticket
func CreatePDF(ticket ticket.Ticket) error {
	// Создаем новый объект PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Устанавливаем шрифт для заголовков
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, fmt.Sprintf("Ticket for %s", ticket.PassengerName))
	pdf.Ln(10)

	// Устанавливаем шрифт для содержания
	pdf.SetFont("Arial", "", 12)

	// Добавляем информацию из структуры
	pdf.Cell(0, 10, fmt.Sprintf("Ticket ID: %s", ticket.ID.String()))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Flight Number: %s", ticket.FlightNumber))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("From: %s", ticket.From))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("To: %s", ticket.To))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Departure: %s", ticket.Departure.Format("2006-01-02 15:04:05")))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Arrival: %s", ticket.Arrival.Format("2006-01-02 15:04:05")))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Seat Number: %s", ticket.SeatNumber))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Class: %s", ticket.Class))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Price: $%.2f", ticket.Price))

	// Сохраняем PDF в файл
	err := pdf.OutputFileAndClose("ticket.pdf")
	if err != nil {
		return fmt.Errorf("error saving PDF: %w", err)
	}

	return nil
}
