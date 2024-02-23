// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "https://github.com/marcusadriano",
            "email": "marcusadriano.pereira@gmail.com"
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
        "/clientes/saldos": {
            "get": {
                "description": "Saldo e somatoria das transacoes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clientes"
                ],
                "summary": "Obtem todos os saldos e a soma de todas as transacoes.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.GetBalanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/clientes/{id}/extrato": {
            "get": {
                "description": "Extrato",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clientes"
                ],
                "summary": "Obtem o extrato com as 10 ultimas transacoes e o saldo atual.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Statements"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/clientes/{id}/transacoes": {
            "post": {
                "description": "Eh necessario informacao do valor, tipo (debito ou credito) e descricao.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clientes"
                ],
                "summary": "Criar uma nova transacao para o usuario.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transacao",
                        "name": "transacao",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.createTransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.TransactionCreated"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.GetBalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "limit": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "sum_transactions": {
                    "type": "integer"
                },
                "transactions": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "http.createTransactionRequest": {
            "type": "object",
            "required": [
                "descricao",
                "tipo",
                "valor"
            ],
            "properties": {
                "descricao": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 1
                },
                "tipo": {
                    "type": "string",
                    "enum": [
                        "c",
                        "d"
                    ]
                },
                "valor": {
                    "type": "integer"
                }
            }
        },
        "service.Balance": {
            "type": "object",
            "properties": {
                "data_extrato": {
                    "type": "string"
                },
                "limite": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "service.Statements": {
            "type": "object",
            "properties": {
                "saldo": {
                    "$ref": "#/definitions/service.Balance"
                },
                "ultimas_transacoes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.Transaction"
                    }
                }
            }
        },
        "service.Transaction": {
            "type": "object",
            "properties": {
                "descricao": {
                    "type": "string"
                },
                "realizado_em": {
                    "type": "string"
                },
                "tipo": {
                    "$ref": "#/definitions/service.TransactionType"
                },
                "valor": {
                    "type": "integer"
                }
            }
        },
        "service.TransactionCreated": {
            "type": "object",
            "properties": {
                "limite": {
                    "type": "integer"
                },
                "saldo": {
                    "type": "integer"
                }
            }
        },
        "service.TransactionType": {
            "type": "string",
            "enum": [
                "d",
                "c"
            ],
            "x-enum-varnames": [
                "Debit",
                "Credit"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.2",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Rinha Backend API - Concorrencia",
	Description:      "Servidor Web \"Rinha de Backend 2 - Concorrencia\".",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
