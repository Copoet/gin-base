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
        "/articles/add": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "添加文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/delete": {
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/info": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "文章详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/articles/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "文章列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/articles/update": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "登陆",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/logout": {
            "get": {
                "tags": [
                    "Auth"
                ],
                "summary": "退出登陆",
                "responses": {}
            }
        },
        "/category/add": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "添加分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分类名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/category/delete": {
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/category/info": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "分类详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/category/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "分类列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/category/tree": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "获取栏目树形分类",
                "responses": {}
            }
        },
        "/category/update": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "更新分类",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分类ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分类名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/manager/add": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "添加管理员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/manager/delete": {
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "删除管理员",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Manager ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/manager/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "管理员列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/manager/update": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Manage"
                ],
                "summary": "更新管理员",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Manager ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/menu/add": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "添加菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "pid",
                        "name": "pid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "icon",
                        "name": "icon",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "path",
                        "name": "path",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "status",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/menu/delete": {
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "删除菜单",
                "responses": {}
            }
        },
        "/menu/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/menu/tree": {
            "get": {
                "tags": [
                    "Menu"
                ],
                "summary": "获取菜单树",
                "responses": {}
            }
        },
        "/menu/update": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "更新菜单",
                "responses": {}
            }
        },
        "/system/add": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "添加系统配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sys_name",
                        "name": "sys_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sys_value",
                        "name": "sys_value",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sys_type",
                        "name": "sys_type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/system/delete": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "删除系统配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/system/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "获取系统配置列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sys_name",
                        "name": "sys_name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/system/update": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "更新系统配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sys_name",
                        "name": "sys_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sys_value",
                        "name": "sys_value",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sys_type",
                        "name": "sys_type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v-0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "gin-base api doc",
	Description:      "This is a sample Server pets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
