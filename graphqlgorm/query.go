/*
 * @Author: your name
 * @Date: 2020-09-25 16:56:48
 * @LastEditTime: 2020-09-27 14:28:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\graphql\query.go
 */
package graphqlgorm

import (
	"graphgo/graphqlgorm/field"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func newQuery(db *gorm.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": field.NewUsers(db),
		},
	})
}
