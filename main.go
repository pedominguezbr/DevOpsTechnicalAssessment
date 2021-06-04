package main

import (
	"bytes"
	"fmt"
	"framework-go/cmd/appserver"
	"framework-go/pkg/config"
	"framework-go/pkg/di"
	"io"
	"io/ioutil"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Api DevOps
// @version 1.0
// @description DevOps Technical Assessment.

// @contact.name Pedro.Dominguez-experis
// @contact.email pe.dominguez.br@gmail.com

// @host localhost:8081
// @BasePath /
func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	g := gin.Default()
	g.Use(RequestLogger())
	d := di.BuildContainer()
	c, _ := config.NewConfig()

	g.Use(cors.Default())
	svr := appserver.NewAppServer(g, d)
	svr.MapRoutes(c)
	// the followings setups were injecteds by Hygen
	if err := svr.SetupDBOracle(); err != nil {
		return err
	}

	engine, _ := svr.Start()
	return engine.Run(fmt.Sprintf(":%s", c.Port))
}

// RequestLogger Read request logger
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)

		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		fmt.Printf("\n\n\nHost: %s\n", c.Request.Host)
		fmt.Printf("METHOD: %s\n", c.Request.Method)
		fmt.Printf("Header: %s\n", c.Request.Header)
		fmt.Printf("Body: %s\n", readBody(rdr1))

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()

	return s
}
