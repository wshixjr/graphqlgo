/*
 * @Author: your name
 * @Date: 2020-09-23 15:16:17
 * @LastEditTime: 2020-09-27 14:27:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\main.go
 */
package main

import (
	"graphgo/datastore"
	"graphgo/graphqlgorm"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

type Params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}

func helloHandler() *handler.Handler {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"pallat": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "anchaleechamaikorn", nil
			},
		},
		"age": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 40, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
}

func main() {
	e := echo.New()
	//e.SetBinder(&binding.EchoBinder{})

	db, err := datastore.NewDB()
	e.Logger.Printf("datastore.NewDB", err)

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	}))

	e.POST("/user", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return context.JSON(http.StatusOK, u)
		}
		return context.JSON(http.StatusOK, u)
	})

	e.POST("/query", func(context echo.Context) error {
		params := new(Params)
		err := context.Bind(&params)
		if err != nil {
			return context.JSON(http.StatusOK, "error")
		}
		e.Logger.Printf("params query : ", params.Query)
		e.Logger.Printf("params OperationName : ", params.OperationName)
		e.Logger.Printf("params Variables : ", params.Variables)
		return context.JSON(http.StatusOK, params)
	})

	e.POST("/graphql", echo.WrapHandler(helloHandler()))

	// graphql
	ggh, err := graphqlgorm.NewHandler(db)

	e.POST("/graphqlgorm", echo.WrapHandler(ggh))

	// e.POST("/graphql", func(context echo.Context) error {
	// 	params := Params{}
	// 	err := context.Bind(&params)
	// 	if err != nil {
	// 		return context.JSON(http.StatusOK, "")
	// 	}
	// 	data := schema.Exec(context.Request().Context(), params.Query, params.OperationName, params.Variables)
	// 	return context.JSON(http.StatusOK, data)
	// })
	e.Logger.Debug("register /hello")
	e.Start(":8080") // HTTP 服务监听在 8080 端口
}
