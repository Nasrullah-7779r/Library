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
        "/all_students": {
            "get": {
                "description": "get all registered students",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "get all registered students",
                "operationId": "get-get-all-student-handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.Student"
                            }
                        }
                    }
                }
            }
        },
        "/borrow_book": {
            "get": {
                "description": "BookBorrowHandler",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "put a borrow request to acquire a book",
                "operationId": "borrow-a-book",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.BorrowRequest"
                        }
                    }
                }
            }
        },
        "/librarian/all_books": {
            "get": {
                "description": "get all books of library",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Librarian"
                ],
                "summary": "all books",
                "operationId": "get all books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/librarian.book"
                            }
                        }
                    }
                }
            }
        },
        "/librarian/book": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Librarian"
                ],
                "summary": "get single book",
                "operationId": "get single book",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/librarian.book"
                        }
                    }
                }
            }
        },
        "/librarian/borrow_requests": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Librarian"
                ],
                "summary": "all books borrow requests",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.BorrowRequest"
                            }
                        }
                    }
                }
            }
        },
        "/librarian/login": {
            "post": {
                "description": "get Librarian logged in",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "login for Librarian",
                "operationId": "login",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "obj"
                        }
                    }
                }
            }
        },
        "/librarian/register_book": {
            "post": {
                "description": "Librarian can register new Book",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Librarian"
                ],
                "summary": "registration of new Book",
                "operationId": "register",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/librarian.book"
                        }
                    }
                }
            }
        },
        "/librarian/token_refresh": {
            "post": {
                "description": "get new access token to access resources, if expired",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refreshToken",
                "operationId": "Librarian refresh-handler",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.AccessToken"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "get login to access resources",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login for students",
                "operationId": "post-login-handler",
                "responses": {
                    "201": {
                        "description": "access_token\", \"refresh_token",
                        "schema": {
                            "type": "obj"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "student can get him/herself registered with data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student"
                ],
                "summary": "register new student",
                "operationId": "post-register-student-handler",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.Student"
                        }
                    }
                }
            }
        },
        "/token_refresh": {
            "post": {
                "description": "get new access token to access resources, if expired",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "refreshToken",
                "operationId": "student refresh-handler",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.AccessToken"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AccessToken": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "common.BorrowRequest": {
            "type": "object",
            "required": [
                "book_author",
                "book_title",
                "borrower_name"
            ],
            "properties": {
                "book_author": {
                    "type": "string"
                },
                "book_title": {
                    "type": "string"
                },
                "borrower_name": {
                    "type": "string"
                },
                "request-id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/common.RequestStatus"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "common.RequestStatus": {
            "type": "string",
            "enum": [
                "pending",
                "in_process",
                "completed"
            ],
            "x-enum-varnames": [
                "PENDING",
                "IN_PROCESS",
                "COMPLETED"
            ]
        },
        "common.Student": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "description": "ID       primitive.ObjectID ` + "`" + `json:\"id\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "librarian.book": {
            "type": "object",
            "required": [
                "author",
                "description",
                "id",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_borrowed": {
                    "type": "boolean",
                    "default": false
                },
                "title": {
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
