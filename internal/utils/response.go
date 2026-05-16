package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Response
	Meta PaginatedMeta `json:"meta"`
}

type PaginatedMeta struct {
	Page       string `json:"status"`
	Limit      string `json:"message,omitempty"`
	Total      int64  `json:"total_items"`
	TotalPages int    `json:"total_pages"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func CreateResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := Response{
		Status:  false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)

}

func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

func InternalServerErrorResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

func PaginatedSuccessResponse(c *gin.Context, message string, data interface{}, meta PaginatedMeta) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Response: Response{
			Status:  true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
