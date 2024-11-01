package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string
	Password string
	UserInfoID uint
}

type UserInfo struct{
	gorm.Model
	Name string
	Birthday string
	Github string
	Description string
	Image Image
	FollowIDs []uint
	GroupIDs []uint
	DirectMessageIDs []uint
	Tag []uint
}

type Group struct {
	gorm.Model
	Name string
	Description string
	UserIDs []uint
	MessageIDs []uint
	Tag []uint
}

type Image struct {
	gorm.Model
	UserID uint
	Url string
}

type Message struct {
	gorm.Model
	UserID uint
	Message string
}

type DirectMessage struct{
	gorm.Model
	FirstID uint
	SecondID uint
	MessageIDs []uint
}

type Tags struct{
	gorm.Model
	TagID uint
	UserIDs []uint
	GroupIDs []uint
}

type Tag struct{
	gorm.Model
  Name string
}
