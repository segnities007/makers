package http

import (
	"fmt"
	"net/http"
	"strconv"

	"server/logics/db"
	// "server/models"

	"github.com/labstack/echo/v4"
)

var dpath string = "server/logics/http/delete.go"

func DELETEs(e *echo.Echo){
	e.DELETE("/user", deleteUser)
	// e.DELETE("/userinfo", deleteUserInfo)
	e.DELETE("/group", deleteGroup)
	e.DELETE("/image", deleteImage)
	e.DELETE("/message", deleteMessage)
	e.DELETE("/dm", deleteDirectMessage)
	e.DELETE("/tag", deleteTag)
}

func deleteUser(c echo.Context)error{

	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil{
		line := 0
		msg := fmt.Sprintf("failed deleteUser %s %d", dpath, line)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err = db.DeleteUser(id)
	if err != nil{
		line := 0
		msg := fmt.Sprintf("failed deleteUser %s %d", dpath, line)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	id, err = strconv.Atoi(c.QueryParam("uiid"))
	if err != nil{
		line := 0
		msg := fmt.Sprintf("failed deleteUser %s %d", dpath, line)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err = deleteUserInfo(c, id)
	if err != nil{
		line := 0
		msg := fmt.Sprintf("failed deleteUser %s %d", dpath, line)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, "success")
}

func deleteUserInfo(c echo.Context, id int)error{

	err := db.DeleteUserInfo(id)
	if err != nil{
		line := 0
		msg := fmt.Sprintf("failed deleteUserInfo %s %d", dpath, line)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return nil
}

func deleteGroup(c echo.Context)error{
	return nil
}

func deleteImage(c echo.Context)error{
	return nil
}

func deleteMessage(c echo.Context)error{
	return nil
}

func deleteDirectMessage(c echo.Context)error{
	return nil
}

func deleteTag(c echo.Context)error{
	return nil
}
