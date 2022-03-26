package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" `
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" `
	Email    string
	Role     int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" `
}

//check user existed or not
func CheckUser(name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

//update checkuser
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

//add user
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//get single user
func GetUser(id int) (User, int) {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

//get user table
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var user User
	var users []User
	var total int64

	db.Model(&user).Count(&total)
	if username == "" {
		err = db.Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	}
	err = db.Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//edit user
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//delete user
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//encrypt password
func ScryptPw(password string) string {
	const Keylen = 10
	salt := make([]byte, 8)

	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, Keylen)

	if err != nil {
		log.Fatal(err)
	}

	fpw := base64.StdEncoding.EncodeToString(HashPw)

	return fpw
}

func CheckLogin(username string, password string) (User, int) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCSE
}
