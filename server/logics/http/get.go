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
	e.GET("/user", getUserWithEP);
	e.GET("/user/:id", getUserWithI)
	e.GET("/userinfo", getUserInfo)
	e.GET("/group", getGroup)
	e.GET("/image", getImage)
	e.GET("/message", getMessage)
	e.GET("/dm", getDirectMessage)
	e.GET("/tag", getTag)
}

func getUserWithI(c echo.Context)error{//modify error handler
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

func getUserWithEP(c echo.Context)error{//modify error handler
	email := c.QueryParam("email")
	password := c.QueryParam("password")

	var u models.User
	err := db.GetUserWithEP(&u, email, password)
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to getUserWithEP %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, u)
}

func getUserInfo(c echo.Context)error{//modify error handler
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

func getImage(c echo.Context)error{//modify error handler
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var g models.Group
	if err := db.GetGroup(&g, id); err != nil {
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, g)
}

func getMessage(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var m models.Message
	if err := db.GetMessage(&m, id); err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, m)
}

func getDirectMessage(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var dm models.DirectMessage
	if err := db.GetDM(&dm, id); err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, dm)
}

func getTag(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	var t models.Tag
	if err := db.GetTag(&t, id); err != nil{
		line := 0
		message := fmt.Sprintf("failed to convert %s %d", gpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, t)
}
