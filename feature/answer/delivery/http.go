package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/thanarat/qa-backend/domain"
	"github.com/thanarat/qa-backend/entity"
)

type Handler struct {
	usecase domain.AnswerUseCase
}

func NewHandler(e *echo.Group, u domain.AnswerUseCase) *Handler {
	h := &Handler{usecase: u}

	e.GET("/:id", h.GetAnswersByQuestionID)
	e.POST("", h.CreateAnswer)
	e.PUT("/:id", h.UpdateAnswer)
	e.DELETE("/:id", h.DeleteAnswer)

	return h
}

func (h *Handler) GetAnswersByQuestionID(c echo.Context) error {
	questionID := c.Param("id")
	answers, err := h.usecase.GetAnswersByQuestionID(questionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, answers)
}

func (h *Handler) CreateAnswer(c echo.Context) error {
	var answer entity.Answer
	if err := c.Bind(&answer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := h.usecase.CreateAnswer(&answer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, answer)
}

func (h *Handler) UpdateAnswer(c echo.Context) error {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id format",
		})
	}

	var answer entity.Answer
	if err := c.Bind(&answer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	answer.ID = uint(idUint)
	if err := h.usecase.UpdateAnswer(&answer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, answer)
}

func (h *Handler) DeleteAnswer(c echo.Context) error {
	id := c.Param("id")
	if err := h.usecase.DeleteAnswer(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
