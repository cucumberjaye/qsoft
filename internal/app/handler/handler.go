package handler

import (
	"fmt"
	"github.com/cucumberjaye/qtsoft/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Service interface {
	CalculateDays(year int) int
}

type Handler struct {
	s Service
}

func New(s Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	when := r.Group("/when", h.PingPong)
	{
		when.GET("/:year", h.DaysToYear)
	}

	return r
}

func (h *Handler) DaysToYear(c *gin.Context) {
	header := c.Param("year")

	year, err := strconv.Atoi(header)
	if err != nil {
		c.String(http.StatusBadRequest, "year must be integer")
		logger.ErrorLogger.Printf("header: %s, error: %s", header, "year must be integer")
		return
	}

	days := h.s.CalculateDays(year)

	s := fmt.Sprintf("%d", days)

	c.String(http.StatusOK, s)
}
