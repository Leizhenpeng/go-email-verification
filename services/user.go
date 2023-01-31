package services

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"leizhenpeng/go-email-verification/models"
	"leizhenpeng/go-email-verification/utils"
	"time"
)

type UserServices interface {
	SignUp(user *models.UserInfo) (*models.DBUserResponse, error)
	Login(user *models.LoginInputReq) (*models.DBUserResponse, error)
	FindUserByEmail(email string) (*models.DBUserResponse, error)
	FindUserById(id string) (*models.UserInfoResponse, error)
	SendEmailVerification(email string) error
	VerifyEmail(info string) error
}

type UserServicesImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (u UserServicesImpl) SendEmailVerification(email string) error {
	info := utils.GenEmailVerificationInfo(email)
	// update database
	SetMailInfo(u.ctx, info, 15*time.Minute)
	// send email
	EmailData := GenEmailData(email, info)
	err := SendEmail(email, EmailData)
	if err != nil {
		return err
	}
	return nil
	// send email
}

func (u UserServicesImpl) VerifyEmail(info string) error {
	//if in redis
	ifExist, err2 := GetMailInfo(u.ctx, info)
	if err2 != nil {
		return err2
	}
	if !ifExist {
		return err2
	}
	DeleteMailInfo(u.ctx, info)
	email, _, err := utils.ParseEmailVerificationInfo(info)
	if err != nil {
		return err
	}
	// check if email is verified
	res, err := GetUserByEmail(u.ctx, email)
	if err != nil {
		return err
	}
	var dbUser models.UserInfo
	res.Decode(&dbUser)
	if dbUser.Verified {
		return errors.New("email already verified")
	}
	// update database
	_, err3 := UpdateUserFieldByEmail(u.ctx, email, "Verified", true)
	if err3 != nil {
		return err
	}
	return nil
	// check if code is correct

}

func (u UserServicesImpl) FindUserById(id string) (*models.UserInfoResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	res, err := GetUserByID(u.ctx, oid)
	if err != nil {
		return nil, errors.New("user not found")
	}
	var dbUserResponse models.UserInfoResponse
	res.Decode(&dbUserResponse)
	return &dbUserResponse, nil
}

func (u UserServicesImpl) Login(user *models.LoginInputReq) (*models.DBUserResponse, error) {
	res, err := GetUserByEmail(u.ctx, user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	var dbUser models.UserInfo
	res.Decode(&dbUser)
	err = utils.ComparePassword(dbUser.Password, user.Password)
	if err != nil {
		return nil, errors.New("password is not correct")
	}
	var dbUserResponse models.DBUserResponse
	res.Decode(&dbUserResponse)
	return &dbUserResponse, nil
}

func NewUserServicesImpl(collection *mongo.Collection, ctx context.Context) *UserServicesImpl {
	return &UserServicesImpl{
		ctx:        ctx,
		collection: collection,
	}
}

func (u UserServicesImpl) SignUp(user *models.UserInfo) (*models.DBUserResponse, error) {
	now := time.Now()
	user.CreateAt = now
	user.UpdateAt = now
	user.Email = utils.FormatEmail(user.Email)
	user.Verified = false

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	result, err := AddUser(u.ctx, user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}
	ObjectID := result.InsertedID
	re, err := GetUserByID(u.ctx, ObjectID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", re)
	var dbUser models.DBUserResponse
	re.Decode(&dbUser)
	return &dbUser, nil
}

func (u UserServicesImpl) FindUserByEmail(email string) (*models.DBUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ UserServices = &UserServicesImpl{}
