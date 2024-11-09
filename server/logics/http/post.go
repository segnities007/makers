package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"server/logics/db"
	"server/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func POSTs(e *echo.Echo){
	e.POST("/user", createUser)
	e.POST("/userinfo", createUserInfo)
	e.POST("/group", createGroup)
	// e.POST("/image", createImage) これは多分使用しない。updateとdeleteとgetだけしか使用しない気がする
	e.POST("/message", createMessage)
	e.POST("/dm", createDirectMessage)
	e.POST("/tag", createTag)
}

var cpath string = "server/logics/http/post.go"

func createUser(c echo.Context)error{
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	user := models.User{Email: email, Password: password}

	db.CreateUser(&user)

		return c.JSON(http.StatusOK, user)
}

///////

func createUserInfo(c echo.Context)error{ //TODO modify line value
	id, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		line := 45
		message := fmt.Sprintf("failed to convert %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	name := c.QueryParam("name")
	birthday, err := strconv.Atoi(c.QueryParam("birthday"))

	if err != nil {
		line := 45
		message := fmt.Sprintf("failed to convert %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	github := c.QueryParam("github")
	description := c.QueryParam("description")

	ui := models.UserInfo{//imageid followids tagids groupids
		Name: name,
		Birthday: birthday,
		Github: github,
		Description: description,
	}

	db.CreateUserInfo(&ui)

	err, imageID := createImage(c, ui.ID, err)
	if err != nil { return err }

	ui.ImageID = imageID
	db.UpdateUserInfo(&ui)
	var u models.User
	db.GetUser(&u, id)
	u.UserInfoID = ui.ID
	db.UpdateUser(&u)

	return c.JSON(http.StatusOK, "create user infomation")
}

///////

func createImage(c echo.Context, userid int, err error) (error, int){	//TODO lineの修正

	if err != nil {
		line := 55
		message := fmt.Sprintf("failed to convert %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message), 0
	}

	image, err :=c.FormFile("image")
	if err != nil{
		line := 62
		message := fmt.Sprintf("failed to get image %s %d", cpath, line)
		return c.JSON(http.StatusBadRequest, message), 0
	}

	src, err := image.Open()
	if err != nil {
		line := 69
		message := fmt.Sprintf("failed to open image %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message), 0
	}
	defer src.Close()

	saveDir := "../images"
	savePath := fmt.Sprintf("%s%s", saveDir, image.Filename)

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		line := 79
		message := fmt.Sprintf("failed to create directory %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message), 0
	}

	dst, err := os.Create(savePath)
	if err != nil{
		line := 86
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message), 0
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message), 0
	}

	img := models.Image{
		UserID: userid,
		Url: savePath,
	}

	db.CreateImage(&img)

	return nil, img.ID
}

///////

func createGroup(c echo.Context)error{
	// name := c.QueryParam("name")
	// userIDs := c.QueryParam("userIDs")
	// messageIDs := c.QueryParam("messageIDs")
	return nil
	//TODO
}

///////

func createMessage(c echo.Context)error{
	// userID := c.QueryParam("userID")
	// message := c.QueryParam("message")
	return nil
	//TODO
}

///////

func createDirectMessage(c echo.Context)error{
	// firstID := c.QueryParam("firstID")
	// secondID := c.QueryParam("secondID")
	// messageIDs := c.QueryParam("messageIDs")
	return nil
	//TODO
}

///////

func createTag(c echo.Context)error{
	// name := c.QueryParam("name")
	return nil
	//TODO
}

///////
