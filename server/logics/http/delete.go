package http

import "github.com/labstack/echo/v4"

func DELETEs(e *echo.Echo){
	e.DELETE("/user", deleteUser)
	e.DELETE("/userinfo", deleteUserInfo)
	e.DELETE("/group", deleteGroup)
	e.DELETE("/image", deleteImage)
	e.DELETE("/message", deleteMessage)
	e.DELETE("/dm", deleteDirectMessage)
	e.DELETE("/tag", deleteTag)
}

func deleteUser(c echo.Context)error{
	return nil
}

func deleteUserInfo(c echo.Context)error{
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
