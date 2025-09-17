package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	})
}

func CreateResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusCreated, message, data)
}

func OKResponse(c *gin.Context, message string, data interface{}) {
	SuccessResponse(c, http.StatusOK, message, data)
}

func BadRequestResponse(c *gin.Context, message string, err interface{}) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func FoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusFound, message, nil)
}

func InternalServerErrorResponse(c *gin.Context, message string, err interface{}) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}
