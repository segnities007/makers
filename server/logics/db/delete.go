package db

import(
	// "gorm.io/gorm"
	"errors"
	"log"
	"server/models"
)

func DeleteUser(id int)error{
	var u models.User
	r := DB.Delete(&u, id)

	if r.Error != nil {
		log.Printf("Error in deleteUser (server/logics/db/db.go): %v\n", r.Error)
		return errors.New("failed to delete user")
	}

	return nil
}

func DeleteGroup(id int)error{
	var g models.Group
	r := DB.Delete(&g)

    if r.Error != nil {
      log.Printf("Error in deleteGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete group")
    }

    return nil
}

func DeleteUserInfo(id int)error{
	var ui models.UserInfo
	if id == 0 {
		return nil
	}
	r := DB.Delete(&ui)

	if r.Error != nil {
      log.Printf("Error in deletUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete userInfo")
    }

    return nil
}

func DeleteImage(id int)error{
	var image models.Image
	r := DB.Delete(&image)

	if r.Error != nil {
      log.Printf("Error in deleteImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete image")
    }

    return nil
}

func DeleteMessage(id int)error{
	var m models.Message
	r := DB.Delete(&m)

	if r.Error != nil {
      log.Printf("Error in deleteMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete message")
    }
    return nil
}

func DeleteDM(id int)error{
	var dm models.DirectMessage
	r := DB.Delete(&dm)

	if r.Error != nil {
      log.Printf("Error in deleteDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete DM")
    }

  return nil
}

func DeleteTag(id int)error{
	var t models.Tag
	r := DB.Delete(&t)

	if r.Error != nil {
      log.Printf("Error in deleteTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to delete tag")
    }

  return nil
}
