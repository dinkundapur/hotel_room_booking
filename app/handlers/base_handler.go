package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (h *BaseHandler) SuccessResponse(c *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}

func (h *BaseHandler) ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: true,
		Message: err.Error(),
	})
}
