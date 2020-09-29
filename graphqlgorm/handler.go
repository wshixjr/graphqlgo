/*
 * @Author: your name
 * @Date: 2020-09-25 16:56:43
 * @LastEditTime: 2020-09-27 14:27:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\graphql\handler.go
 */
package graphqlgorm

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQuery(db),
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
