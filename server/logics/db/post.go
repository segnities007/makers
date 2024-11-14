package db

import(
	// "gorm.io/gorm"
	"errors"
	"log"
	"server/models"
)


func CreateUser(u *models.User) error {
    r := DB.Create(&u)

    if r.Error != nil {
        log.Printf("Error in createUser (server/logics/db/db.go): %v\n", r.Error)
        return errors.New("failed to create user")
    }

    return nil
}

func CreateGroup(g *models.Group)error{
	r := DB.Create(&g)

    if r.Error != nil {
      log.Printf("Error in createGroup (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create group")
    }

    return nil
}

func CreateUserInfo(ui *models.UserInfo)error{
	r := DB.Create(&ui)

    if r.Error != nil {
      log.Printf("Error in createUserInfo (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create userInfo")
    }

    return nil
}

func CreateImage(image *models.Image)error{
	r := DB.Create(&image)

    if r.Error != nil {
      log.Printf("Error in createImage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create image")
    }

    return nil
}

func CreateMessage(m *models.Message)error{
	r := DB.Create(&m)

    if r.Error != nil {
      log.Printf("Error in createMessage (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create message")
    }

    return nil
}

func CreateDM(dm *models.DirectMessage)error{
	r := DB.Create(&dm)

    if r.Error != nil {
      log.Printf("Error in createDM (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create DM")
    }

  return nil
}


func CreateTag(t *models.Tag)error{
	r := DB.Create(&t)

    if r.Error != nil {
      log.Printf("Error in createTag (server/logics/db/db.go): %v\n", r.Error)
      return errors.New("failed to create tag")
    }

  return nil
}
