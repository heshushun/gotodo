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
        "termsOfService": "http://me.shiniao.fun/",
        "contact": {
            "name": "shiniao",
            "url": "http://me.shiniao.fun/",
            "email": "zhuzhezhe5@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/sd/cpu": {
            "get": {
                "description": "Checks the cpu usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the cpu usage",
                "responses": {
                    "200": {
                        "description": "{\"status\":\"OK\", \"info\":\"Load average: 1.82, 1.05, 0.85 | Cores: 2\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/sd/disk": {
            "get": {
                "description": "Checks the disk usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the disk usage",
                "responses": {
                    "200": {
                        "description": "{\"status\":\"OK\", \"info\":\"Free space: 44429MB (43GB) / 119674MB (116GB) | Used: 39%\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/sd/health": {
            "get": {
                "description": "Shows OK as the ping-pong result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Shows OK as the ping-pong result",
                "responses": {
                    "200": {
                        "description": "{\"status\":\"OK\", \"info\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/sd/ram": {
            "get": {
                "description": "Checks the ram usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the ram usage",
                "responses": {
                    "200": {
                        "description": "{\"status\":\"OK\", \"info\":\"Free space: 5293MB (5GB) / 7665MB (7GB) | Used: 69%\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/v1/todos": {
            "get": {
                "description": "Get all todos from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"total\":233, \"todos\":[{\"ID\":91,\"title\": \"??????\", \"completed\": 1,\"CreatedAt\": \"2019-10-12T10:10:05+08:00\",\"UpdatedAt\": \"2019-10-12T10:16:24+08:00\",\"DeletedAt\": null}]}}",
                        "schema": {
                            "$ref": "#/definitions/model.TodoModel"
                        }
                    }
                }
            }
        },
        "/v1/todos/": {
            "post": {
                "description": "Add a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Add new todos to the database",
                "parameters": [
                    {
                        "description": "The todo info",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.TodoModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":\"create successful.\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/v1/todos/{id}": {
            "get": {
                "description": "Get a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Get a todo",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"ID\":91,\"title\": \"??????\", \"completed\": 1,\"CreatedAt\": \"2019-10-12T10:10:05+08:00\",\"UpdatedAt\": \"2019-10-12T10:16:24+08:00\",\"DeletedAt\": null}}",
                        "schema": {
                            "$ref": "#/definitions/model.TodoModel"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Update a todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The todo's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The todo info",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.TodoModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":\"update successful.\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todo"
                ],
                "summary": "Delete a todo",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":\"delete successful.\"}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.message": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.TodoModel": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "integer"
                },
                "title": {
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
	Host:        "todo.shiniao.fun",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "A todos application API",
	Description: "This is a todos application server.",
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
