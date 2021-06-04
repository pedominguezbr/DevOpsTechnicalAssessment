// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Pedro.Dominguez-experis",
            "email": "pe.dominguez.br@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/DevOps": {
            "post": {
                "description": "payload for the endpoint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DevOps"
                ],
                "summary": "DevOps - api test",
                "parameters": [
                    {
                        "description": "requestDevops",
                        "name": "requestDevops",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/devOps.RequestDevops"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Registros creados correctamente."
                    },
                    "400": {
                        "description": "Error en la Data enviada."
                    },
                    "500": {
                        "description": "Error Interno en el api."
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get Health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Get Health",
                "responses": {
                    "200": {
                        "description": "Respuesta de health",
                        "schema": {
                            "$ref": "#/definitions/appserver.HealthRsp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appserver.HealthRsp": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "devOps.RequestDevops": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string",
                    "example": "Rita Asturia"
                },
                "message”": {
                    "type": "string",
                    "example": "This is a test"
                },
                "timeToLifeSec": {
                    "type": "number",
                    "example": 45
                },
                "to": {
                    "type": "string",
                    "example": "Juan Perez"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8081",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Api DevOps",
	Description: "DevOps Technical Assessment.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
