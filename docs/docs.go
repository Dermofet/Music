// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Dermofet3_3"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/signin": {
            "post": {
                "description": "Авторизация пользователя с использованием имени пользователя и пароля.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Вход пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя для входа",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Токен авторизации",
                        "schema": {
                            "$ref": "#/definitions/view.TokenView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Ошибка авторизации"
                    },
                    "422": {
                        "description": "Ошибка при обработке данных"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Регистрация нового пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя для регистрации",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Токен авторизации",
                        "schema": {
                            "$ref": "#/definitions/view.TokenView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "422": {
                        "description": "Ошибка при обработке данных"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/music/catalog": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение всех треков",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Получение всех треков",
                "responses": {
                    "200": {
                        "description": "Данные трека",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MusicView"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/music/new": {
            "post": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Создание нового трека",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Создание трека",
                "parameters": [
                    {
                        "description": "Данные трека",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.MusicCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Трек создан"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/music/popular": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение треков отсортированных по популярности",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Получение треков отсортированных по популярности",
                "responses": {
                    "200": {
                        "description": "Список треков",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MusicView"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/music/release": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение треков отсортированных по популярности",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Получение треков отсортированных по популярности",
                "responses": {
                    "200": {
                        "description": "Список треков",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MusicView"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/music/{id}": {
            "put": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Обновление трека",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Обновление трека",
                "parameters": [
                    {
                        "description": "Данные трека",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.MusicCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Трек обновлен"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Удаление трека",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Music"
                ],
                "summary": "Удаление трека",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id трека",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Трек удален"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/add-track/{id}": {
            "post": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Добавление трека в список понравившихся",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Добавить трека в список понравившихся",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор трека (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Трек успешно добавлен в список понравившихся"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Трек не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/get-liked-tracks": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение списка понравившихся треков",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Показать понравившиеся треки",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.ListMusicView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Трек не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/id/{id}": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение информации о пользователе по его уникальному идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получение пользователя по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор пользователя (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/view.UserView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение пользователя по его уникальному идентификатору из JWT токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получение пользователя по JWT токену",
                "responses": {
                    "200": {
                        "description": "Данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/view.UserView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Обновление информации о пользователе по его уникальному идентификатору из JWT токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Обновление пользователя по JWT токену",
                "parameters": [
                    {
                        "description": "Данные пользователя для обновления",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленные данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/view.UserView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "422": {
                        "description": "Ошибка при обработке данных"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Удаление пользователя по его уникальному идентификатору из JWT токена.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Удаление пользователя по JWT токену",
                "responses": {
                    "204": {
                        "description": "Пользователь успешно удален"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/remove-track/{id}": {
            "delete": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Удаление трека из списка понравившихся",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Удалить трек из списка понравившихся",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор трека (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Трек успешно удален из списка понравившихся"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Трек не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/username/{username}": {
            "get": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Получение информации о пользователе по его уникальному идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Получение пользователя по Username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username пользователя",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/view.UserView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Обновление информации о пользователе по его уникальному идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Обновление пользователя по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор пользователя (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные пользователя для обновления",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Обновленные данные пользователя",
                        "schema": {
                            "$ref": "#/definitions/view.UserView"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "422": {
                        "description": "Ошибка при обработке данных"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JwtAuth": []
                    }
                ],
                "description": "Удаление пользователя по его уникальному идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Удаление пользователя по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор пользователя (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Пользователь успешно удален"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "401": {
                        "description": "Неавторизованный запрос"
                    },
                    "404": {
                        "description": "Пользователь не найден"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.CustomDate": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "entity.MusicCreate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "release": {
                    "$ref": "#/definitions/entity.CustomDate"
                }
            }
        },
        "entity.UserCreate": {
            "type": "object",
            "properties": {
                "password": {
                    "description": "Пароль",
                    "type": "string"
                },
                "username": {
                    "description": "Имя пользователя",
                    "type": "string"
                }
            }
        },
        "view.ListMusicView": {
            "type": "object",
            "properties": {
                "musics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.MusicView"
                    }
                }
            }
        },
        "view.MusicView": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "view.TokenView": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "view.UserView": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JwtAuth": {
            "description": "JWT Bearer токен для аутентификации",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Golang Test API",
	Description:      "API for Golang Test Project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
