// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/todo": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Create todo",
                "parameters": [
                    {
                        "description": "Todo to create",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todo.ReqBodyTodo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The created todo",
                        "schema": {
                            "$ref": "#/definitions/todo.Todo"
                        }
                    }
                }
            }
        },
        "/api/todo/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Delete todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of todo to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The deleted todo",
                        "schema": {
                            "$ref": "#/definitions/todo.Todo"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Finish todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of todo to finish",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The updated todo",
                        "schema": {
                            "$ref": "#/definitions/todo.Todo"
                        }
                    }
                }
            }
        },
        "/api/todos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "List todo",
                "responses": {
                    "200": {
                        "description": "List of all todos",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todo.Todo"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "todo.ReqBodyTodo": {
            "type": "object",
            "properties": {
                "finished": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "todo.Todo": {
            "type": "object",
            "properties": {
                "finished": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
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
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Todo API",
	Description: "This is a exmaple of Todo API for K8S traning",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}