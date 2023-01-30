package models

import "time"

type SignUpInput struct {
	Name     string    `json:"name" bson:"Name"  binding:"required"`
	Email    string    `json:"email" bson:"Email"  binding:"required"`
	Password string    `json:"password" bson:"Password"  binding:"required"`
	Verified bool      `json:"verified" bson:"Verified"`
	CreateAt time.Time `json:"createAt" bson:"CreateAt"`
	UpdateAt time.Time `json:"updateAt" bson:"UpdateAt"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"Email"  binding:"required"`
	Password string `json:"password" bson:"Password"  binding:"required"`
}
