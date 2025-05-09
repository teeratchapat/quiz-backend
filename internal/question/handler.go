package question

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	repo := &Repository{DB: db}
	service := &Service{Repo: repo}
	handler := &Handler{Service: service}

	e.GET("/questions", handler.GetQuestions)
	e.GET("/questions/:id", handler.GetQuestionByID)
	e.POST("/questions", handler.CreateQuestion)
	e.DELETE("/questions/:id", handler.DeleteQuestion)
}

type Handler struct {
	Service *Service
}

func (h *Handler) GetQuestions(c echo.Context) error {
	questions, err := h.Service.GetQuestions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, questions)
}

func (h *Handler) GetQuestionByID(c echo.Context) error {
	id := c.Param("id")
	question, err := h.Service.GetQuestionByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Question not found"})
	}
	return c.JSON(http.StatusOK, question)
}

func (h *Handler) CreateQuestion(c echo.Context) error {
	var q Question
	if err := c.Bind(&q); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	id, err := h.Service.CreateQuestionReturningID(q.QuestionText)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert question"})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message":     "Question created successfully",
		"question_id": id,
	})
}

func (h *Handler) DeleteQuestion(c echo.Context) error {
	id := c.Param("id")
	err := h.Service.DeleteQuestion(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete question"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Question deleted"})
}
