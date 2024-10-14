package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:20;null;unique;default:null"`
	Lastname string `gorm:"type:string;size:20;null;unique;default:null"`
	Age      int    `gorm:"type:int;;null;default:0"`
}

func NewUserEntity() *User {
	return &User{}
}
