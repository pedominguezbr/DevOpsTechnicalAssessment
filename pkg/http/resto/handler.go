package resto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

// NewHealthCtrl healthCheck
func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

func (h healthCtrl) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
