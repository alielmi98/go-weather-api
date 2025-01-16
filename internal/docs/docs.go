package docs

import "github.com/swaggo/swag"

var docTemplate = `{
    "openapi": "3.0.0",
    "info": {
        "title": "{{.Title}}",
        "version": "{{.Version}}",
        "description": "{{.Description}}"
    },
    "servers": [
        {
            "url": "{{.BasePath}}"
        }
    ],
    "paths": {
        // Define your API paths and operations here
    }
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.example.com",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "My API",
	Description:      "This is a sample API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
