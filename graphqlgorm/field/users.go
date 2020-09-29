/*
 * @Author: your name
 * @Date: 2020-09-25 16:58:00
 * @LastEditTime: 2020-09-27 14:28:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\graphql\field\users.go
 */
package field

import (
	"graphgo/domain/model"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var user = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"name":      &graphql.Field{Type: graphql.String},
			"age":       &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
			"updatedAt": &graphql.Field{Type: graphql.String},
			"deletedAt": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

func NewUsers(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(user),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"age": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"createdAt": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			// 偏移量
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			}, // 返回的数据个数
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var u []*model.User

			if err := db.Last(&u, p.Args).Error; err != nil {
				// do something
			}

			return u, nil
		},
		Description: "user",
	}
}
