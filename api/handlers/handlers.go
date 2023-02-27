package handlers

import (
	"logger-service/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoNothing(c *gin.Context) {

}

func Home(c *gin.Context) {
	helpers.ResponseSuccess(c, nil, http.StatusAccepted)
}
