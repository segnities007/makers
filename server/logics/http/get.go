package http

import (
	"fmt"
	"net/http"
	"strconv"

	"server/logics/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

var gpath string = "server/logics/http/get.go"

func GETs(e *echo.Echo){
	e.GET("/user/:id", getUser)
	e.GET("/userinfo", getUserInfo)
	e.GET("/group", getGroup)
	e.GET("/image", getImage)
	e.GET("/message", getMessage)
	e.GET("/dm", getDirectMessage)
	e.GET("/tag", getTag)
}

func getUser(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var u models.User
	db.GetUser(&u, id)

	return c.JSON(http.StatusOK, u)
}

func getUserInfo(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var ui models.UserInfo
	db.GetUserInfo(&ui, id)

	return c.JSON(http.StatusOK, ui)
}

func getImage(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var image models.Image
	if err := db.GetImage(&image, id); err != nil {
		line := 0
		message := fmt.Sprintf("failed to get image %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.File(image.Url)
}

func getGroup(c echo.Context)error{

	return nil
	//TODO
}

func getMessage(c echo.Context)error{

	return nil
	//TODO
}

func getDirectMessage(c echo.Context)error{

	return nil
	//TODO
}

func getTag(c echo.Context)error{

	return nil
	//TODO
}
