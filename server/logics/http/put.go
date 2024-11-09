package http

import (
	// "net/http"

	"github.com/labstack/echo/v4"
)

func PUTs(e *echo.Echo){
	e.PUT("/user", updateUser)
	e.PUT("/userinfo", updateUserInfo)
	e.PUT("/group", updateGroup)
	e.PUT("/message", updateMessage)
	e.PUT("/dm", updateDirectMessage)
	e.PUT("/tag", updateTag)
}

func updateUser(c echo.Context)error{
	// id := c.QueryParam("id")
	// email := c.QueryParam("email")
	// password := c.QueryParam("password")
	return nil
	//TODO
}

func updateUserInfo(c echo.Context)error{
	// id := c.QueryParam("id")
	// name := c.QueryParam("name")
	// birthday := c.QueryParam("birthday")
	// github := c.QueryParam("github")
	// description := c.QueryParam("description")
	// tags := c.QueryParam("tags")
	// followIDs := c.QueryParam("followIDs")
	// groupIDs := c.QueryParam("groupIDs")
	// directMessageIDs := c.QueryParam("directMessageIDs")

	// image, err :=c.FormFile("image")
	// if err != nil{
	// 	return c.JSON(http.StatusBadRequest, err)
	// }
	return nil
	//TODO
}

func updateGroup(c echo.Context)error{
	// id := c.QueryParam("id")
	// name := c.QueryParam("name")
	// userIDs := c.QueryParam("userIDs")
	// messageIDs := c.QueryParam("messageIDs")
	return nil
	//TODO
}

// func updateImage(c echo.Context)error{
// 	id := c.QueryParam("id")
// 	userID := c.QueryParam("userID")
// 	url := c.QueryParam("url")
//  多分これはいらない
// }

func updateMessage(c echo.Context)error{
	// id := c.QueryParam("id")
	// userID := c.QueryParam("userID")
	// message := c.QueryParam("message")
	return nil
	//TODO
}

func updateDirectMessage(c echo.Context)error{
	// id := c.QueryParam("id")
	// firstID := c.QueryParam("firstID")
	// secondID := c.QueryParam("secondID")
	// groupIDs := c.QueryParam("groupIDs")
	return nil
	//TODO
}

func updateTag(c echo.Context)error{
	// id := c.QueryParam("id")
	// name := c.QueryParam("name")
	return nil
	//TODO
}
