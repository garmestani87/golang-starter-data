// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/v1/redis/{key}": {
            "get": {
                "description": "get value by key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show value for key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/redis/{key}/{value}": {
            "put": {
                "description": "put key and value",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "put key and value",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "value",
                        "name": "value",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/": {
            "post": {
                "description": "save user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "save user",
                "parameters": [
                    {
                        "description": "User Dto",
                        "name": "UserDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseModel-dto_UserDto-dto_UserDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ResponseModel-dto_UserDto-dto_UserDto"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "description": "find user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "find user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ExceptionModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "stackTrace": {
                    "type": "string"
                }
            }
        },
        "common.ResponseModel-dto_UserDto-dto_UserDto": {
            "type": "object",
            "properties": {
                "exceptionModel": {
                    "$ref": "#/definitions/common.ExceptionModel"
                },
                "input": {
                    "$ref": "#/definitions/dto.UserDto"
                },
                "output": {
                    "$ref": "#/definitions/dto.UserDto"
                },
                "responseCode": {
                    "type": "integer"
                }
            }
        },
        "dto.UserDto": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
