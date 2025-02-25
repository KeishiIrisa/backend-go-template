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
        "/login": {
            "post": {
                "description": "Log in by providing email and password to receive a JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Log in a user",
                "parameters": [
                    {
                        "description": "User Login Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserLoginSchemaIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "Create a new user by providing email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserSchemaIn"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-09-06T14:00:00Z"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2023-09-06T14:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "john.doe@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "Doe"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-09-06T14:00:00Z"
                }
            }
        },
        "schemas.UserLoginSchemaIn": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "string@string.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8,
                    "example": "stringstring"
                }
            }
        },
        "schemas.UserSchemaIn": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1,
                    "example": "John"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1,
                    "example": "Doe"
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Simple Gin Backend API",
	Description:      "This is a simple backend using Gin and GORM.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
