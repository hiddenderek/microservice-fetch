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
        "/receipts/process": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Process a receipt",
                "parameters": [
                    {
                        "description": "Receipt Data",
                        "name": "receipt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Receipt"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/receipts/{id}/points": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Calculate points for a receipt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Item": {
            "type": "object",
            "required": [
                "price",
                "shortDescription"
            ],
            "properties": {
                "price": {
                    "type": "string"
                },
                "shortDescription": {
                    "type": "string"
                }
            }
        },
        "models.Receipt": {
            "type": "object",
            "required": [
                "items",
                "purchaseDate",
                "purchaseTime",
                "retailer",
                "total"
            ],
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchaseTime": {
                    "type": "string"
                },
                "retailer": {
                    "type": "string"
                },
                "total": {
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