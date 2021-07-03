package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"infotech.umm.ac.id/auth/auth"
	"infotech.umm.ac.id/auth/config"
	"net/http"
	"strings"
	"time"
)

func (UserCredential) TableName() string {
	return "user"
}

const domain = "localhost"

//const domain = "infotech.umm.ac.id"

type UserCredential struct {
	gorm.Model
	UserName       string
	Password       string
	UserStudent    UserStudent    `gorm:"foreignkey:UserID;references:ID"`
	UserAssistant  UserAssistant  `gorm:"foreignkey:UserID;references:ID"`
	UserInstructor UserInstructor `gorm:"foreignkey:UserID;references:ID"`
	UserAdmin      UserAdmin      `gorm:"foreignkey:UserID;references:ID"`
}

type UserAdmin struct {
	gorm.Model
	UserID uint
	Active bool
}

type UserStudent struct {
	gorm.Model
	UserID uint
	Active bool
}

type UserAssistant struct {
	gorm.Model
	UserID uint
	Active bool
}

type UserInstructor struct {
	gorm.Model
	UserID uint
	Active bool
}

func (idb *InDB) GetUser(username string) (UserCredential, error) {
	var (
		data UserCredential
		err  error
	)

	fetchResult := idb.DB.
		Preload("UserInstructor").
		Preload("UserStudent").
		Preload("UserAssistant").
		Preload("UserAdmin").
		Where("user_name = ?", username).
		Last(&data)

	if fetchResult.Error != nil {
		err = fetchResult.Error
		return data, err
	}

	return data, nil
}
func (idb *InDB) UnauthorizeUser(c *gin.Context) {
	c.SetCookie("ifx-ath", "", -1, "/", domain, false, true)
	c.SetCookie("ifx-at", "", -1, "/", domain, false, false)
	c.SetCookie("ifx-st", "", -1, "/", domain, false, true)
}

func (idb *InDB) AuthorizeUser(c *gin.Context) {
	var (
		data      UserCredential
		dataInput UserCredential
		err       error
	)
	err = c.BindJSON(&dataInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Not Authorized : wrong json")
		return
	}

	data, err = idb.GetUser(dataInput.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Not Authorized : user not found")
		return
	}

	isPasswordMatch, err := CheckPassword(dataInput.Password, data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Not Authorized : %s", err.Error()))
		return
	}

	if !isPasswordMatch {
		c.JSON(http.StatusUnauthorized, "Not Authorized : password not match")
		return
	}

	key, err := auth.ImportKey(config.GetKeyStorage())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ro := "U"
	if data.UserAssistant.Active {
		ro += "A"
	}
	if data.UserInstructor.Active {
		ro += "I"
	}
	if data.UserAdmin.Active {
		ro += "D"
	}

	tok, err := auth.GenerateJWT(key, auth.Body{
		IAT: time.Now().Unix(),
		ISS: "IF-MAIN-AUTH",
		UID: data.ID,
		RO:  ro,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	tokSplit := strings.Split(tok, ".")

	c.SetCookie("ifx-ath", tokSplit[0], 0, "/", domain, false, true)
	c.SetCookie("ifx-at", tokSplit[1], 0, "/", domain, false, false)
	c.SetCookie("ifx-st", tokSplit[2], 0, "/", domain, false, true)

	c.JSON(http.StatusOK, "Authorized")
	return
}
