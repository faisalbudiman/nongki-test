package address

import (
	"nongki-test/internal/middleware"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Create, middleware.Authentication)
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
	g.PUT("/:id", h.UpdateByID, middleware.Authentication)
	g.DELETE("/soft-delete/:id", h.SoftDeleteByID, middleware.Authentication)
	g.DELETE("/:id", h.HardDeleteByID, middleware.Authentication)
}
