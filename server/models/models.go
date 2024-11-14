package models

import (
	"strconv"
	"strings"
)

func ParseIntSlice(s string) ([]int, error) {
    if s == "" {
        return nil, nil
    }

    var slice []int
    parts := strings.Split(s, ",")

    for _, part := range parts {
        num, err := strconv.Atoi(strings.TrimSpace(part))
        if err != nil {
            return nil, err
        }
        slice = append(slice, num)
    }
    return slice, nil
}

type Image struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	UserID int    `json:"userid"`
	Url    string `json:"url"`
}

type User struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	UserInfoID int    `json:"userinfoid"`
}

type UserInfo struct {
	ID               int      `gorm:"primaryKey" json:"id"`
	Name             string   `json:"name"`
	Birthday         int      `json:"birthday"`
	Github           string   `json:"github"`
	Description      string   `json:"description"`
	ImageID          int      `json:"imageid"`
	FollowIDs        []int `gorm:"type:json" json:"followids"`
	GroupIDs         []int `gorm:"type:json" json:"groupids"`
	DirectMessageIDs []int `gorm:"type:json" json:"directmessageids"`
	TagIDs           []int `gorm:"type:json" json:"tagids"`
}

type Group struct {
	ID          int      `gorm:"primaryKey" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UserIDs     []int `gorm:"type:json" json:"userids"`
	MessageIDs  []int `gorm:"type:json" json:"messageids"`
	TagIDs      []int `gorm:"type:json" json:"tagids"`
}

type Message struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	UserID  int    `json:"userid"`
	Message string `json:"message"`
}

type DirectMessage struct {
	ID         int      `gorm:"primaryKey" json:"id"`
	FirstID    int      `json:"firstid"`
	SecondID   int      `json:"secondid"`
	MessageIDs []int `gorm:"type:json" json:"messageids"`
}

type Tag struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
