package db

import(
	"gorm.io/gorm"
	"errors"
	"log"
	"server/models"
)

var DB *gorm.DB

//////////////////////////////////////////////////////////////////////////////

func CreateUser(u *models.User) error {
    r := DB.Create(&u)

    if r.Error != nil {
        log.Printf("Error in createUser (server/logics/db/db.go): %v\n", r.Error)
        return errors.New("failed to create user")
    }

    return nil
}

func GetUser(u *models.User, id int)error{
	r := DB.First(u, id)

	if r.Error != nil {
		log.Printf("Error in getUser (server/logics/db/db.go): %v\n", r.Error)
		return errors.New("failed to get user")
	}

	return nil
}

func UpdateUser(u *models.User)error{
	r := DB.Save(u)

	if r.Error != nil {
		log.Printf("Error in updateUser (server/logics/db/db.go): %v\n", r.Error)
		return errors.New("failed to update user")
	}

	return nil
}

func DeleteUser(u *models.User)error{
	r := DB.Delete(u)

	if r.Error != nil {
		log.Printf("Error in deleteUser (server/logics/db/db.go): %v\n", r.Error)
		return errors.New("failed to delete user")
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateGroup(g *models.Group)error{
	r := DB.Create(&g)

    if r.Error != nil {
      log.Printf("Error in createGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create group")
    }

    return nil
}

func GetGroup(g *models.Group)error{
	r := DB.First(&g)

    if r.Error != nil {
      log.Printf("Error in getGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get group")
    }

    return nil
}

func UpdateGroup(g *models.Group)error{
	r := DB.Save(&g)

    if r.Error != nil {
      log.Printf("Error in updateGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update group")
    }

    return nil
}

func DeleteGroup(g *models.Group)error{
	r := DB.Delete(&g)

    if r.Error != nil {
      log.Printf("Error in deleteGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete group")
    }

    return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateUserInfo(ui *models.UserInfo)error{
	r := DB.Create(&ui)

    if r.Error != nil {
      log.Printf("Error in createUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create userInfo")
    }

    return nil
}

func GetUserInfo(ui *models.UserInfo, id int)error{
	r := DB.First(&ui, id)

	if r.Error != nil {
      log.Printf("Error in getUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get userInfo")
    }

    return nil
}

func UpdateUserInfo(ui *models.UserInfo)error{
	r := DB.Save(&ui)

	if r.Error != nil {
      log.Printf("Error in updateUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update userInfo")
    }

    return nil
}

func DeleteUserInfo(ui *models.UserInfo)error{
	r := DB.Delete(&ui)

	if r.Error != nil {
      log.Printf("Error in deletUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete userInfo")
    }

    return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateImage(image *models.Image)error{
	r := DB.Create(&image)

    if r.Error != nil {
      log.Printf("Error in createImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create image")
    }

    return nil
}

func GetImage(image *models.Image)error{
	r := DB.First(&image)

	if r.Error != nil {
      log.Printf("Error in getImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get image")
    }

    return nil
}

func UpdateImage(image *models.Image)error{
	r := DB.Save(&image)

	if r.Error != nil {
      log.Printf("Error in updateImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update image")
    }

    return nil
}

func DeleteImage(image *models.Image)error{
	r := DB.Delete(&image)

	if r.Error != nil {
      log.Printf("Error in deleteImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete image")
    }

    return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateMessage(m *models.Message)error{
	r := DB.Create(&m)

    if r.Error != nil {
      log.Printf("Error in createMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create message")
    }

    return nil
}

func GetMessage(m *models.Message)error{
	r := DB.First(&m)

	if r.Error != nil {
      log.Printf("Error in getMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get message")
    }

    return nil
}

func UpdateMessage(m *models.Message)error{
	r := DB.Save(&m)

	if r.Error != nil {
      log.Printf("Error in updateMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update message")
    }

    return nil
}

func DeleteMessage(m *models.Message)error{
	r := DB.Delete(&m)

	if r.Error != nil {
      log.Printf("Error in deleteMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete message")
    }
    return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateDM(dm *models.DirectMessage)error{
	r := DB.Create(&dm)

    if r.Error != nil {
      log.Printf("Error in createDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create DM")
    }

  return nil
}

func GetDM(dm *models.DirectMessage)error{
	r := DB.First(&dm)

	if r.Error != nil {
      log.Printf("Error in getDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get DM")
    }

  return nil
}

func UpdateDM(dm *models.DirectMessage)error{
	r := DB.Save(&dm)

	if r.Error != nil {
      log.Printf("Error in updateDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update DM")
    }

  return nil
}

func DeleteDM(dm *models.DirectMessage)error{
	r := DB.Delete(&dm)

	if r.Error != nil {
      log.Printf("Error in deleteDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete DM")
    }

  return nil
}

//////////////////////////////////////////////////////////////////////////////

func CreateTag(t *models.Tag)error{
	r := DB.Create(&t)

    if r.Error != nil {
      log.Printf("Error in createTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create tag")
    }

  return nil
}

func GetTag(t *models.Tag)error{
	r := DB.First(&t)

	if r.Error != nil {
      log.Printf("Error in getTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get tag")
    }

  return nil
}

func UpdateTag(t *models.Tag)error{
	r := DB.Save(&t)

	if r.Error != nil {
      log.Printf("Error in updateTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update tag")
    }

  return nil
}

func DeleteTag(t *models.Tag)error{
	r := DB.Delete(&t)

	if r.Error != nil {
      log.Printf("Error in deleteTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete tag")
    }

  return nil
}
