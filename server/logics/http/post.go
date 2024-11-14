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
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := models.User{Email: email, Password: password}

	db.CreateUser(&user)

		return c.JSON(http.StatusOK, user)
}

///////

func createUserInfo(c echo.Context)error{ //TODO modify line value
	id, err := strconv.Atoi(c.FormValue("id"))

	if err != nil {
		line := 45
		message := fmt.Sprintf("failed to convert %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	name := c.FormValue("name")
	birthday, err := strconv.Atoi(c.FormValue("birthday"))

	if err != nil {
		line := 45
		message := fmt.Sprintf("failed to convert %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	github := c.FormValue("github")
	description := c.FormValue("description")

	ui := models.UserInfo{//imageid followids tagids groupids
		Name: name,
		Birthday: birthday,
		Github: github,
		Description: description,
	}

	if err := db.CreateUserInfo(&ui); err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	err, imageID := createImage(c, ui.ID, err)
	if err != nil { return err }

	ui.ImageID = imageID
	if err := db.UpdateUserInfo(&ui); err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	var u models.User
	if err := db.GetUser(&u, id); err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	u.UserInfoID = ui.ID
	if err := db.UpdateUser(&u); err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

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

	saveDir := "./images/"
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

func createGroup(c echo.Context)error{//TODO modify error handler
	name := c.FormValue("name")

	userIDs, err := models.ParseIntSlice(c.FormValue("userids"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusBadRequest, message)
	}

	messageIDs, err := models.ParseIntSlice(c.FormValue("messageIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	g := models.Group{
		Name: name,
		UserIDs: userIDs,
		MessageIDs: messageIDs,
	}

	if err := db.CreateGroup(&g); err != nil {
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, g)
}

///////

func createMessage(c echo.Context)error{//modify err handler
	userID, err := strconv.Atoi(c.FormValue("userID"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	message := c.FormValue("message")

	m := models.Message{
		UserID: userID,
		Message: message,
	}

	if err := db.CreateMessage(&m); err != nil {
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}
	return c.JSON(http.StatusOK, m)
}

///////

func createDirectMessage(c echo.Context)error{
	firstID, err := strconv.Atoi(c.FormValue("firstID"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	secondID, err := strconv.Atoi(c.FormValue("secondID"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	messageIDs, err := models.ParseIntSlice(c.FormValue("messageIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	dm := models.DirectMessage{
		FirstID: firstID,
		SecondID: secondID,
		MessageIDs: messageIDs,
	}

	if err := db.CreateDM(&dm); err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.JSON(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, dm)
}

///////

func createTag(c echo.Context)error{
	// name := c.FormValue("name")
	return nil
	//TODO
}

///////
