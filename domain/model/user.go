/*
 * @Author: your name
 * @Date: 2020-09-25 16:55:59
 * @LastEditTime: 2020-09-25 17:35:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \graphgo\domain\model\user.go
 */
package model

import "time"

type User struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`
	Age        string    `json:"age"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeleteedAt time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
