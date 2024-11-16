package http

import (
	"fmt"
	// "mime/multipart"
	"net/http"
	"io"
	"os"
	"strconv"

	"server/logics/db"
	"server/models"

	"github.com/labstack/echo/v4"
)

func PUTs(e *echo.Echo){
	e.PUT("/user", updateUser)
	e.PUT("/userinfo", updateUserInfo)
	e.PUT("/image", updateImage)
	e.PUT("/group", updateGroup)
	e.PUT("/message", updateMessage)
	e.PUT("/dm", updateDirectMessage)
	e.PUT("/tag", updateTag)
}

//modify error handler and return

func updateUser(c echo.Context)error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	u := models.User{
		ID: id,
		Email: email,
		Password: password,
	}
	if err := db.UpdateUser(&u); err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	return c.String(http.StatusOK, "success")
}

func updateUserInfo(c echo.Context)error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	name := c.FormValue("name")
	birthday, err := strconv.Atoi(c.FormValue("birthday"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	github := c.FormValue("github")
	description := c.FormValue("description")
	tagIDs, err := models.ParseIntSlice(c.FormValue("tagIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}
	followIDs, err := models.ParseIntSlice(c.FormValue("followIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	groupIDs, err := models.ParseIntSlice(c.FormValue("groupIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	directMessageIDs, err := models.ParseIntSlice(c.FormValue("directMessageIDs"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	ui := models.UserInfo{
		ID: id,
		Name: name,
		Birthday: birthday,
		Github: github,
		Description: description,
		TagIDs: tagIDs,
		FollowIDs: followIDs,
		GroupIDs: groupIDs,
		DirectMessageIDs: directMessageIDs,
	}

	if err := db.UpdateUserInfo(&ui); err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, ui)
}

func updateImage(c echo.Context)error{

	image, err :=c.FormFile("image")
	if err != nil{
		line := 62
		message := fmt.Sprintf("failed to get image %s %d", cpath, line)
		return c.JSON(http.StatusBadRequest, message)
	}

	userID, err := strconv.Atoi(c.FormValue("userid"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	imageID, err := strconv.Atoi(c.FormValue("imageid"))
	if err != nil{
		line := 0
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	db.DeleteImage(imageID)

	src, err := image.Open()
	if err != nil {
		line := 69
		message := fmt.Sprintf("failed to open image %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}
	defer src.Close()

	saveDir := "./images/"
	savePath := fmt.Sprintf("%s%s", saveDir, image.Filename)

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		line := 79
		message := fmt.Sprintf("failed to create directory %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	dst, err := os.Create(savePath)
	if err != nil{
		line := 86
		message := fmt.Sprintf("failed to create file %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	img := models.Image{
		UserID: userID,
		Url: savePath,
	}
	if err := db.UpdateImage(&img); err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.String(http.StatusInternalServerError, message)
	}

	return c.JSON(http.StatusOK, img)
}

func updateGroup(c echo.Context)error{
	groupID, err := strconv.Atoi(c.FormValue("groupid"))
	if err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.String(http.StatusBadRequest, message)
	}

	userIDs, err := models.ParseIntSlice(c.FormValue("userIDs"))
	if err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.String(http.StatusBadRequest, message)
	}

	messageIDs, err := models.ParseIntSlice(c.FormValue("messageIDs"))
	if err != nil{
		line := 93
		message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
		return c.String(http.StatusBadRequest, message)
	}

	g := models.Group{
		ID: groupID,
		UserIDs: userIDs,
		MessageIDs: messageIDs,
	}

	if err := db.UpdateGroup(&g); err != nil{
		if err != nil{
			line := 93
			message := fmt.Sprintf("failed to copy image %s %d", cpath, line)
			return c.String(http.StatusBadRequest, message)
	}

	return c.JSON(http.StatusOK, g)
}

// func updateMessage(c echo.Context)error{
// 	// id := c.FormValue("id")
// 	// userID := c.FormValue("userID")
// 	// message := c.FormValue("message")
// 	return nil
// 	//TODO
// }

// func updateDirectMessage(c echo.Context)error{
// 	// id := c.FormValue("id")
// 	// firstID := c.FormValue("firstID")
// 	// secondID := c.FormValue("secondID")
// 	// groupIDs := c.FormValue("groupIDs")
// 	return nil
// 	//TODO
// }

// func updateTag(c echo.Context)error{
// 	// id := c.FormValue("id")
// 	// name := c.FormValue("name")
// 	return nil
// 	//TODO
// }
