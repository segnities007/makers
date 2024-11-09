package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type IntArray []int

func (a *IntArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, a)
}

func (a IntArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type Image struct {
	ID     int   `gorm:"primaryKey" json:"id"`
	UserID int   `json:"userid"`
	Url    string `json:"url"`
}

type User struct {
	ID         int   `gorm:"primaryKey" json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserInfoID int   `json:"userinfoid"`
}

type UserInfo struct {
	ID               int     `gorm:"primaryKey" json:"id"`
	Name             string   `json:"name"`
	Birthday         int     `json:"birthday"`
	Github           string   `json:"github"`
	Description      string   `json:"description"`
	ImageID          int     `json:"imageid" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FollowIDs        IntArray `gorm:"serializer:json" json:"followids"`
	GroupIDs         IntArray `gorm:"serializer:json" json:"groupids"`
	DirectMessageIDs IntArray `gorm:"serializer:json" json:"directmessageids"`
	TagIDs           IntArray `gorm:"serializer:json" json:"tagids"`
}

type Group struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserIDs     IntArray `gorm:"serializer:json" json:"userids"`
	MessageIDs  IntArray `gorm:"serializer:json" json:"messageids"`
	TagIDs      IntArray `gorm:"serializer:json" json:"tagids"`
}

type Message struct {
	ID      int   `gorm:"primaryKey" json:"id"`
	UserID  int   `json:"userid"`
	Message string `json:"message"`
}

type DirectMessage struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	FirstID    int     `json:"firstid"`
	SecondID   int     `json:"secondid"`
	MessageIDs IntArray `gorm:"serializer:json" json:"messageids"`
}

type Tag struct {
	ID   int   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
