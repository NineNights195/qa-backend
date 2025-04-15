package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanarat/qa-backend/domain"
)

type Handler struct {
	usecase domain.CategoryUseCase
}

func NewHandler(e *echo.Group, u domain.CategoryUseCase) *Handler {
	h := &Handler{usecase: u}

	e.GET("", h.GetAllCategories)

	return h
}

func (h *Handler) GetAllCategories(c echo.Context) error {
	categories, err := h.usecase.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, categories)
}
