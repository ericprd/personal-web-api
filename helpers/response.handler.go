package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHandler(c *gin.Context, data any, status int) {
	if status == http.StatusOK || status == http.StatusCreated {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data": data,
		})
		return
	}

	if status == http.StatusBadRequest || status == http.StatusBadGateway {
		errMsg := fmt.Sprint(data)
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"data": errMsg,
		})
		return
	}
}
