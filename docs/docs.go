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
        "/tags": {
            "get": {
                "description": "Return list of tags.",
                "tags": [
                    "tags"
                ],
                "summary": "Get All tags.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "obejct"
                        }
                    }
                }
            },
            "post": {
                "description": "Save tags data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Create tags",
                "parameters": [
                    {
                        "description": "Create tags",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTagsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tags/{tagID}": {
            "delete": {
                "description": "Remove tags data by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Delete tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/tags/{tagId}": {
            "get": {
                "description": "Return the tahs whoes tagId valu mathes id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Get Single tags by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update tags by id",
                        "name": "tagId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update tags data.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Update tags",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update tags by id",
                        "name": "tagId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update tags",
                        "name": "tags",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateTagsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateTagsRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
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
	BasePath:         "/api",
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
