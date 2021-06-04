package appserver

import (
	"framework-go/pkg/config"
	"log"

	"github.com/gin-gonic/gin"

	// The follow packages were added for Hygen
	"framework-go/pkg/features/devOps"

	resto "framework-go/pkg/http/resto"
)

// HealthRsp representa un modelo de respuesta para Health
type HealthRsp struct {
	Status string `json:"status" example:"success"`
}

func (as *appserver) MapRoutes(c *config.Config) {

	// Group : v1
	apiV1 := as.router.Group("/")

	as.healthRoutes(apiV1)

	as.LoginRoutes(apiV1, c)
	//DevOps
	as.DevOpsRoutes(apiV1, c)

}

// Health godoc
// @Summary Get Health
// @Description Get Health
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthRsp "Respuesta de health"
// @Router /health [get]
func (as *appserver) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := resto.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}

// devOps godoc
// @Summary DevOps - api test
// @Description payload for the endpoint
// @Tags DevOps
// @Accept  json
// @Produce  json
// @Success 201 "Enviados correctamente."
// @Success 400 "Error en la Data enviada."
// @Success 500 "Error Interno en el api."
// @Param requestDevops body devOps.RequestDevops true "requestDevops"
// @Router /DevOps [post]
func (as *appserver) DevOpsRoutes(api *gin.RouterGroup, d *config.Config) {
	routes := api.Group("/")
	{
		var devOpsSvco devOps.Service
		as.cont.Invoke(func(s devOps.Service) {
			devOpsSvco = s
		})

		log.Println("Contenido del devOpsSvco: ", devOpsSvco)

		svro := resto.NewDevOpsCtrl(devOpsSvco, d)
		routes.POST("/DevOps", svro.DevOps)
	}
}

// devOps godoc
// @Summary login - api test
// @Description payload for the endpoint
// @Tags login
// @Accept  json
// @Produce  json
// @Success 201 "login correctamente."
// @Success 400 "Error en la Data enviada."
// @Success 500 "Error Interno en el api."
// @Param requestDevops body devOps.RequestDevops true "requestDevops"
// @Router /login [post]
func (as *appserver) LoginRoutes(api *gin.RouterGroup, d *config.Config) {
	routes := api.Group("/")
	{
		var devOpsSvco devOps.Service
		as.cont.Invoke(func(s devOps.Service) {
			devOpsSvco = s
		})

		log.Println("Contenido del devOpsSvco: ", devOpsSvco)

		svro := resto.NewDevOpsCtrl(devOpsSvco, d)
		routes.POST("/login", svro.Login)
	}
}
