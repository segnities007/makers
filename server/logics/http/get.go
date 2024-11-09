package http

import "github.com/labstack/echo/v4"

func GETs(e *echo.Echo){
	e.GET("/user", getUser)
	e.GET("/userinfo", getUserInfo)
	e.GET("/group", getGroup)
	e.GET("/image", getImage)
	e.GET("/message", getMessage)
	e.GET("/dm", getDirectMessage)
	e.GET("/tag", getTag)
}

func getUser(c echo.Context)error{
	// id := c.QueryParam("")
	return nil
	//TODO
}

func getUserInfo(c echo.Context)error{

	return nil
	//TODO
}

func getGroup(c echo.Context)error{

	return nil
	//TODO
}

func getImage(c echo.Context)error{

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
