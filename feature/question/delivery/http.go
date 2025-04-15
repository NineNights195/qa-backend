package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
)

type Handler struct {
	usecase domain.QuestionUseCase
}

func NewHandler(e *echo.Group, u domain.QuestionUseCase) *Handler {
	h := &Handler{usecase: u}

	e.GET("", h.GetAllQuestions)
	e.GET("/:id", h.GetQuestionByID)
	e.POST("", h.CreateQuestion)
	e.PUT("/:id", h.UpdateQuestion)
	e.DELETE("/:id", h.DeleteQuestion)

	return h
}

func (h *Handler) GetAllQuestions(c echo.Context) error {
	questions, err := h.usecase.GetAllQuestions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, questions)
}

func (h *Handler) GetQuestionByID(c echo.Context) error {
	id := c.Param("id")
	question, err := h.usecase.GetQuestionByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, question)
}

func (h *Handler) CreateQuestion(c echo.Context) error {
	var question entity.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	fmt.Println(question)

	if question.CategoryID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "String is empty",
		})
	}

	if err := h.usecase.CreateQuestion(&question); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, question)
}

func (h *Handler) UpdateQuestion(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id format",
		})
	}

	var question entity.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	question.ID = uint(idUint)
	if err := h.usecase.UpdateQuestion(&question); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, question)
}

func (h *Handler) DeleteQuestion(c echo.Context) error {
	id := c.Param("id")
	if err := h.usecase.DeleteQuestion(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
