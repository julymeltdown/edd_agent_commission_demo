package handler

import (
	"edd_agent_commission/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommissionHandler struct {
	service *service.CommissionService
}

func NewCommissionHandler(service *service.CommissionService) *CommissionHandler {
	return &CommissionHandler{service: service}
}

func (h *CommissionHandler) GetCommissionsByAgentName(c echo.Context) error {
	agentName := c.Param("agentName")
	commissions, err := h.service.GetCommissionsByAgentName(c.Request().Context(), agentName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commissions)
}
