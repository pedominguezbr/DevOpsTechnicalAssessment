package resto

import (
	"fmt"
	"framework-go/pkg/config"
	"framework-go/pkg/features/auth"
	"framework-go/pkg/features/devOps"
	util "framework-go/pkg/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

// DevOpsCtrl handle control
type DevOpsCtrl struct {
	svc devOps.Service
	d   *config.Config
}

// NewDevOpsCtrl s
func NewDevOpsCtrl(svc devOps.Service, d *config.Config) *DevOpsCtrl {
	return &DevOpsCtrl{svc, d}
}

func (s *DevOpsCtrl) Login(ctx *gin.Context) {
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	var email string

	email = "pe.dominguez.br@gmail.com"

	signedToken, err := jwtWrapper.GenerateToken(email)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"msg": "error signing token",
		})
		ctx.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	ctx.JSON(200, tokenResponse)
}

// DevOps
func (s *DevOpsCtrl) DevOps(ctx *gin.Context) {
	tokenJWT := ctx.Request.Header.Get("X-JWT-KWY")

	//
	jwtWrapper := auth.JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(tokenJWT)

	if err != nil {
		log.Println("JWT error: ", err.Error())
		stringResponse := util.ResponseMessageService(err.Error())

		ctx.JSON(http.StatusBadRequest, util.StringResponseToResponseObj(stringResponse))
		return
	}

	log.Println("claims: ", claims)

	var datosDevOps devOps.RequestDevops
	//log.Println("Contenido del svc: ", s.svc)
	APIKey := ctx.Request.Header.Get("X-Parse-REST-API-Key")

	//Validar el api key
	if APIKey != "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c" {
		fmt.Printf("Found 0 results for API Key: %s\n", APIKey)
		stringResponse := util.ResponseMessageService("Authentication failed")

		ctx.JSON(http.StatusUnauthorized, util.StringResponseToResponseObj(stringResponse))
		return
	}

	if err := ctx.ShouldBindJSON(&datosDevOps); err != nil {
		log.Println("Contenido del error: ", err.Error())
		stringResponseErr := "Error en el formato" //util.ResponseService(false, 2, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, util.StringResponseToResponseObj(stringResponseErr))
		return
	}

	stringResponse := util.ResponseMessageService(fmt.Sprintf("Hellow %v your message will be send", datosDevOps.To))

	ctx.JSON(http.StatusOK, util.StringResponseToResponseObj(stringResponse))
	return
}
