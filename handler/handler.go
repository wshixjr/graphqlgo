/*
 * @Author: your name
 * @Date: 2020-09-25 16:54:05
 * @LastEditTime: 2020-09-25 16:54:46
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\handler\handler.go
 */
package graphql

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome")
	}
}

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var u []*model.User

		if err := db.Find(&u).Error; err != nil {
			// error handling here
			return err
		}

		return c.JSON(http.StatusOK, u)
	}
}
