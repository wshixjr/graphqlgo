/*
 * @Author: your name
 * @Date: 2020-09-25 15:42:31
 * @LastEditTime: 2020-09-27 14:26:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\datastore\db.go
 */
package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	// DBMS := "mysql"
	// mySqlConfig := &mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "root",
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3306",
	// 	DBName:               "golang-with-echo-gorm-graphql-example_db",
	// 	AllowNativePasswords: true,
	// 	Params: map[string]string{
	// 		"parseTime": "true",
	// 	},
	// }
	dsn := "root:root@tcp(127.0.0.1:3306)/golang-with-echo-gorm-graphql-example_db?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
