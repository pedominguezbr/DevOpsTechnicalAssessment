package appserver

import (
	"fmt"
	_ "framework-go/docs"
	"framework-go/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"

	// the following packages were injected for Hygen
	"database/sql"
)

type appserver struct {
	router *gin.Engine
	cont   *dig.Container
}

func NewAppServer(e *gin.Engine, c *dig.Container) *appserver {
	return &appserver{
		router: e,
		cont:   c,
	}
}

// the following setup db were injected for Hygen

func (as *appserver) SetupDBOracle() error {
	var db []*sql.DB

	if err := as.cont.Invoke(func(d []*sql.DB) { db = d }); err != nil {
		return err
	}
	fmt.Printf("db ORACLE = %v \n", db)

	return nil
}

func (as *appserver) Start() (*gin.Engine, error) {

	var cfg *config.Config
	if err := as.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.Port)
	url := ginSwagger.URL(uri)
	as.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return as.router, nil
}
