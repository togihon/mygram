// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "togi.mare@gmail.com"
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
        "/comments": {
            "get": {
                "description": "User can retrieve all comments and no need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get all comments",
                "responses": {
                    "200": {
                        "description": "Will send all comments",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If there is no comment, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/comments/{id}": {
            "get": {
                "description": "User can retrieve a comment and no need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get one comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If a comment's id matches with the parameter",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If the comments's id doesn't match with the parameter, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can edit their own comment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Edit a comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "your comment",
                        "name": "message",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If all the parameters are valid",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If there is something wrong, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can delete their own comment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete a comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If comment is exist and it's your own comment",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "If the comment's id is not your own and if the comment doesn't exist, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/comments/{photoID}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can create a comment.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Create a comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photo id",
                        "name": "photoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "your comment",
                        "name": "message",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "If all of the parameters filled and you're login",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "401": {
                        "description": "If you are not login or some parameters not filled, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    },
                    "404": {
                        "description": "If photo id's not found",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/photos": {
            "get": {
                "description": "User can retrieve all photos and no need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get all photos",
                "responses": {
                    "200": {
                        "description": "Will send all photo datas",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If there is no photo, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can upload a photo.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Upload a photo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "photo title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "photo caption",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "photo url",
                        "name": "photo_url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "If all of the parameters filled and you're logged in",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If you are not login or some parameters not filled, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/photos/{id}": {
            "get": {
                "description": "User can retrieve a photo and no need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get one photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photo id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If a photo's id matches with the parameter",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If the photo's id doesn't match with the parameter, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can edit their own photo.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Edit a photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photo id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "photo title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "photo caption",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "photo url",
                        "name": "photo_url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If the parameters are valid",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "401": {
                        "description": "If there is something wrong, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can delete their own photo.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Delete a photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photo id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If photo is exist and it's your own photo",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "400": {
                        "description": "If the photo is not your own or if the photo doesn't exist, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/social-media": {
            "get": {
                "description": "User can retrieve all social media and no need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "social-medias"
                ],
                "summary": "Get all social media",
                "responses": {
                    "200": {
                        "description": "Will send all social media datas",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If there is no social media, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can create a social media.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "social-medias"
                ],
                "summary": "Create a social media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "social media name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "social media url",
                        "name": "social_media_url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "If all of the parameters filled and you are logged in",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "401": {
                        "description": "If you are not login or some parameters not filled, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/social-media/{id}": {
            "get": {
                "description": "User can retrieve a social media and doesn't need to login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "social-medias"
                ],
                "summary": "Get one social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "social media id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If a social media's id matches with the parameter",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "404": {
                        "description": "If the social media's id doesn't match with the parameter, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can edit their own social media.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "social-medias"
                ],
                "summary": "Edit a social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "social media id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "social media name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "social media url",
                        "name": "social_media_url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If all the parameters are valid",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    },
                    "400": {
                        "description": "If there is something wrong, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "User can delete their own social media.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "social-medias"
                ],
                "summary": "Delete a social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "social media id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "If social media is exist and it's your own social media",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "If social media's id is not your own or if the comment doesn't exist, error will appear",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login to system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Jika email dan password benar, maka akan mendapatkan token",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseLogin"
                        }
                    },
                    "401": {
                        "description": "Jika email dan password salah, maka akan muncul error",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register an account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User's password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User's age",
                        "name": "age",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Jika semua field benar, maka akun akan dibuat ",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseRegister"
                        }
                    },
                    "400": {
                        "description": "Jika terdapat kesalahan akan muncul error",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseFailed"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "created"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "entity.ResponseFailed": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "entity.ResponseLogin": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "token": {
                            "type": "string",
                            "example": "eyJhbGciOiJI...."
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "entity.ResponseRegister": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "age": {
                            "type": "integer",
                            "example": 18
                        },
                        "email": {
                            "type": "string",
                            "example": "user@mail.com"
                        },
                        "id": {
                            "type": "integer",
                            "example": 1
                        },
                        "username": {
                            "type": "string",
                            "example": "user"
                        }
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "entity.ResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://mygram-production-2f89.up.railway.app/",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyGram API",
	Description:      "This is an API for MyGram APP. To use all of the services, please login first and get the token.\nAfter that, click the \"Authorize\" button at the right and a pop-up window will appear. type \"Bearer <your-token>\". Example: Bearer eyJhbGciOiJIUzI1...",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
