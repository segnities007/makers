package models

import (
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Email      string	`json:"email"`
	Password   string	`json:"password"`
	UserInfoID uint	`json:"userinfoid"`
}

type UserInfo struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name        string	`json:"name"`
	Birthday    string	`json:"birthday"`
	Github      string	`json:"github"`
	Description string	`json:"description"`
	ImageID     uint	`json:"imageid"`
	FollowIDs     []uint	`json:"followids"`
	GroupIDs      []uint	`json:"groupids"`
	DirectMessageIDs []uint 	`json:"directmessageids"`
	TagIDs        []uint	`json:"tagids"`
}

type Image struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID uint	`json:"userid"`
	Url    string	`json:"url"`
}

type Group struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name        string	`json:"name"`
	Description string	`json:"description"`
	UserIDs       []uint	`json:"userids"`
	MessageIDs    []uint	`json:"messageids"`
	TagIDs        []uint	`json:"tagids"`
}

type Message struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID          uint	`json:""`
	Message         string	`json:""`
}

type DirectMessage struct{
	ID uint `gorm:"primaryKey" json:"id"`
	FirstID  uint	`json:"firstid"`
	SecondID uint	`json:"secondid"`
	MessageIDs []uint	`json:"messageids"`
}

type Tag struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name   string	`json:"name"`
}
