package httphelper

import (
	"net/http"

	"github.com/Shachindra/SuperTags/gateway/api/types"
	"github.com/Shachindra/SuperTags/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// TODO: add method for internet error with common msg
func ErrResponse(c *gin.Context, statusCode int, errMessage string) {
	response := types.ApiResponse{
		Status: statusCode,
		Error:  errMessage,
	}
	c.JSON(response.Status, response)
}

func CErrResponse(c *gin.Context, status int, customStatusCode int, errMessage string) {
	response := types.ApiResponse{
		Status: customStatusCode,
		Error:  errMessage,
	}
	c.JSON(status, response)
}

func SuccessResponse(c *gin.Context, message string, payload interface{}) {
	response := types.ApiResponse{
		Status:  http.StatusOK,
		Payload: payload,
		Message: message,
	}
	c.JSON(response.Status, response)
}

func InternalServerError(c *gin.Context) {
	response := types.ApiResponse{
		Status: http.StatusInternalServerError,
		Error:  "unexpected error occurred",
	}
	c.JSON(response.Status, response)
}

func NewInternalServerError(c *gin.Context, message string, format string, args ...interface{}) {
	logwrapper.Errorf(format, args...)
	response := types.ApiResponse{
		Status:  http.StatusInternalServerError,
		Error:   "unexpected error occurred",
		Message: message,
	}
	c.JSON(response.Status, response)
}
