package models

import (
	"gorm.io/gorm"
)

// User はアプリケーションのユーザーを表します。
type User struct {
	gorm.Model
	Email      string
	Password   string
	UserInfoID uint
	UserInfo   UserInfo `gorm:"foreignKey:UserInfoID"`
}

// UserInfo はユーザーの詳細情報を表します。
type UserInfo struct{
	gorm.Model
	Name        string
	Birthday    string
	Github      string
	Description string
	ImageID     uint     `gorm:"uniqueIndex"` // 外部キーとしてImageIDを追加
	Image       *Image   `gorm:"foreignKey:ImageID"` // ポインタを使用して循環参照を防止
	Follows     []*User  `gorm:"many2many:user_follows;"`
	Groups      []*Group `gorm:"many2many:user_groups;"`
	DirectMessages []*DirectMessage `gorm:"many2many:user_direct_messages;"`
	Tags        []*Tag   `gorm:"many2many:user_tags;"`
}

// Image はユーザーの画像を表します。
type Image struct {
	gorm.Model
	UserID uint
	Url    string
	// UserInfo 参照を削除して循環参照を防止
}

// Group はユーザーグループを表します。
type Group struct {
	gorm.Model
	Name        string
	Description string
	Users       []*User `gorm:"many2many:user_groups;"`
	Messages    []Message
	Tags        []*Tag `gorm:"many2many:group_tags;"`
}

// Message はグループまたはダイレクトメッセージ内のメッセージを表します。
type Message struct {
	gorm.Model
	UserID          uint
	Message         string
	GroupID         uint
	Group           Group `gorm:"foreignKey:GroupID"`
	DirectMessageID uint
	DirectMessage   DirectMessage  `gorm:"foreignKey:DirectMessageID"`
}

// DirectMessage はユーザー間のダイレクトメッセージを表します。
type DirectMessage struct{
	gorm.Model
	FirstID  uint
	SecondID uint
	Messages []Message `gorm:"foreignKey:DirectMessageID"`
}

// Tag はユーザーやグループに関連付けられるタグを表します。
type Tag struct{
	gorm.Model
	Name   string
	Users  []*User `gorm:"many2many:user_tags;"`
	Groups []*Group `gorm:"many2many:group_tags;"`
}
