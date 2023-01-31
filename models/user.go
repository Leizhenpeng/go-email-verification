package models

import "time"

type LoginInputReq struct {
	Email    string `json:"email" bson:"Email"  binding:"required" default:"rivers@88.com"`
	Password string `json:"password" bson:"Password"  binding:"required" default:"123456"`
}

type UserInfo struct {
	Name         string    `json:"name" bson:"Name"  binding:"required"`
	Email        string    `json:"email" bson:"Email"  binding:"required"`
	Password     string    `json:"password" bson:"Password"  binding:"required"`
	VerifiedCode string    `json:"verifiedCode,omitempty" bson:"VerifiedCode,omitempty"`
	Verified     bool      `json:"verified" bson:"Verified" default:"false"`
	CreateAt     time.Time `json:"createAt" bson:"CreateAt"`
	UpdateAt     time.Time `json:"updateAt" bson:"UpdateAt"`
}

type DBUserResponse struct {
	Name     string `json:"name" bson:"Name"  binding:"required"`
	Email    string `json:"email" bson:"Email"  binding:"required"`
	ID       string `json:"id" bson:"_id"`
	Verified bool   `json:"verified" bson:"Verified"`
}

type UserInfoResponse struct {
	Name     string `json:"name" bson:"Name"  binding:"required"`
	Email    string `json:"email" bson:"Email"  binding:"required"`
	Verified bool   `json:"verified" bson:"Verified"`
}
