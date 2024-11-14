package db

import(
	"errors"
	"log"
	"server/models"
)

func GetUserWithEP(u *models.User, e string, p string)error{
	r := DB.Where("email = ? AND password = ?",e,p).First(&u)

	if r.Error != nil{
		return r.Error
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

func GetGroup(g *models.Group)error{
	r := DB.First(&g)

    if r.Error != nil {
      log.Printf("Error in getGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get group")
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

func GetImage(image *models.Image, id int)error{
	r := DB.First(&image, id)

	if r.Error != nil {
      log.Printf("Error in getImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get image")
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

func GetDM(dm *models.DirectMessage)error{
	r := DB.First(&dm)

	if r.Error != nil {
      log.Printf("Error in getDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to get DM")
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
