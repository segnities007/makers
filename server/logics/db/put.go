package db

import(
	// "gorm.io/gorm"
	"errors"
	"log"
	"server/models"
)


func UpdateUser(u *models.User)error{
	r := DB.Save(u)

	if r.Error != nil {
		log.Printf("Error in updateUser (server/logics/db/db.go): %v\n", r.Error)
		return errors.New("failed to update user")
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

func UpdateUserInfo(ui *models.UserInfo)error{
	r := DB.Save(&ui)

	if r.Error != nil {
      log.Printf("Error in updateUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update userInfo")
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

func UpdateMessage(m *models.Message)error{
	r := DB.Save(&m)

	if r.Error != nil {
      log.Printf("Error in updateMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update message")
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

func UpdateTag(t *models.Tag)error{
	r := DB.Save(&t)

	if r.Error != nil {
      log.Printf("Error in updateTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to update tag")
    }

  return nil
}
