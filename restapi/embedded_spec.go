// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Authentication Service API",
    "title": "Auth Service",
    "contact": {
      "name": "NEWS API Support"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "$ref": "./paths/default.yml"
    },
    "/v1/auth/email-confirmation/{token}": {
      "$ref": "./paths/auth/registerConfirmation.yml"
    },
    "/v1/auth/refresh-token": {
      "$ref": "./paths/auth/refreshToken.yml"
    },
    "/v1/auth/signin": {
      "$ref": "./paths/auth/login.yml"
    },
    "/v1/auth/signup": {
      "$ref": "./paths/auth/register.yml"
    },
    "/v1/auth/user-details": {
      "$ref": "./paths/auth/registerDetails.yml"
    },
    "/v1/oauth2/oauth-signin-cb": {
      "$ref": "./paths/auth/callbackLogin.yml"
    },
    "/v1/oauth2/oauth-signup-cb": {
      "$ref": "./paths/auth/callbackRegister.yml"
    }
  },
  "definitions": {
    "generalResponse": {
      "$ref": "./definitions/generalResponse.yml"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Authentication Service API",
    "title": "Auth Service",
    "contact": {
      "name": "NEWS API Support"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/auth/email-confirmation/{token}": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "RegisterConfirmation",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/auth/refresh-token": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "RefreshToken",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/auth/signin": {
      "post": {
        "tags": [
          "auth"
        ],
        "operationId": "LoginUser",
        "parameters": [
          {
            "name": "loginRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/auth/signup": {
      "post": {
        "tags": [
          "auth"
        ],
        "operationId": "RegisterUser",
        "parameters": [
          {
            "name": "registerRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/auth/user-details": {
      "post": {
        "tags": [
          "auth"
        ],
        "operationId": "RegisterUserDetails",
        "parameters": [
          {
            "name": "registerRequest",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerDetailsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/oauth2/oauth-signin-cb": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "LoginCallback",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    },
    "/v1/oauth2/oauth-signup-cb": {
      "get": {
        "tags": [
          "auth"
        ],
        "operationId": "RegisterCalback",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "400": {
            "description": "BAD REQUEST",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "401": {
            "description": "UNAUTHORIZED",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "403": {
            "description": "FORBIDDEN",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "404": {
            "description": "NOT FOUND",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          },
          "500": {
            "description": "INTERNAL SERVER ERROR",
            "schema": {
              "$ref": "#/definitions/generalResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "generalResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "object",
          "properties": {
            "code": {
              "type": "integer"
            },
            "message": {
              "type": "string"
            }
          }
        },
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "object",
      "properties": {
        "data": {
          "type": "object",
          "properties": {
            "accessToken": {
              "type": "string"
            },
            "customizations": {
              "type": "object",
              "properties": {
                "backgroundColor": {
                  "type": "string"
                },
                "backgroundImageURL": {
                  "type": "string"
                },
                "header1Font": {
                  "type": "string"
                },
                "header1FontColor": {
                  "type": "string"
                },
                "header1FontSize": {
                  "type": "string"
                },
                "header2Font": {
                  "type": "string"
                },
                "header2FontColor": {
                  "type": "string"
                },
                "header2FontSize": {
                  "type": "string"
                },
                "logoURL": {
                  "type": "string"
                },
                "navbarColor": {
                  "type": "string"
                },
                "navbarFont": {
                  "type": "string"
                },
                "navbarFontColor": {
                  "type": "string"
                },
                "navbarFontSize": {
                  "type": "string"
                },
                "navbarSelectColor": {
                  "type": "string"
                },
                "textFont": {
                  "type": "string"
                },
                "textFontColor": {
                  "type": "string"
                },
                "textFontSize": {
                  "type": "string"
                }
              }
            },
            "domain": {
              "type": "string"
            },
            "expiresIn": {
              "type": "string"
            },
            "role": {
              "type": "string"
            },
            "userRef": {
              "type": "string"
            },
            "userType": {
              "type": "string"
            }
          }
        },
        "error": {
          "type": "object",
          "properties": {
            "code": {
              "type": "integer"
            },
            "message": {
              "type": "string"
            }
          }
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "registerDetailsRequest": {
      "type": "object",
      "properties": {
        "details": {
          "type": "string"
        },
        "domain": {
          "type": "string"
        }
      }
    },
    "registerRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    }
  }
}`))
}