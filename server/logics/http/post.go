package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func POSTs(e *echo.Echo){
	e.POST("/user", createUser)
	e.POST("/userinfo", createUserInfo)
	e.POST("/group", createGroup)
	e.POST("/image", createImage)
	e.POST("/message", createMessage)
	e.POST("/dm", createDirectMessage)
	e.POST("/tag", createTag)
}

func createUser(c echo.Context)error{
	email := c.QueryParam("email")
	password := c.QueryParam("password")

	//TODO
}

func createUserInfo(c echo.Context)error{
	name := c.QueryParam("name")
	birthday := c.QueryParam("birthday")
	github := c.QueryParam("github")
	description := c.QueryParam("description")
	tags := c.QueryParam("tags")
	followIDs := c.QueryParam("followIDs")
	groupIDs := c.QueryParam("groupIDs")
	directMessageIDs := c.QueryParam("directMessageIDs")

	image, err :=c.FormFile("image")
	if err != nil{
		return c.JSON(http.StatusBadRequest, err)
	}

	//TODO
}

func createGroup(c echo.Context)error{
	name := c.QueryParam("name")
	userIDs := c.QueryParam("userIDs")
	messageIDs := c.QueryParam("messageIDs")

	//TODO
}

func createImage(c echo.Context)error{
	userID := c.QueryParam("userID")
	url := c.QueryParam("url")

	//TODO
}

func createMessage(c echo.Context)error{
	userID := c.QueryParam("userID")
	message := c.QueryParam("message")

	//TODO
}

func createDirectMessage(c echo.Context)error{
	firstID := c.QueryParam("firstID")
	secondID := c.QueryParam("secondID")
	messageIDs := c.QueryParam("messageIDs")

	//TODO
}

func createTag(c echo.Context)error{
	name := c.QueryParam("name")

	//TODO
}
