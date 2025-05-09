package choice

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	repo := &Repository{DB: db}
	service := &Service{Repo: repo}
	handler := &Handler{Service: service}

	e.GET("/questions/:id/choices", handler.GetChoices)
	e.POST("/questions/:id/choices", handler.CreateChoice)
}

type Handler struct {
	Service *Service
}

func (h *Handler) GetChoices(c echo.Context) error {
	id := c.Param("id")
	choices, err := h.Service.GetChoices(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, choices)
}

func (h *Handler) CreateChoice(c echo.Context) error {
	questionID := c.Param("id")

	var choice Choice
	if err := c.Bind(&choice); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	choice.QuestionID = questionID

	if err := h.Service.CreateChoice(choice); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to insert choice"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Choice created successfully"})
}
