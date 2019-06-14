package model

import "fmt"

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Role            int    `json:"role"`
	TelephoneNumber string `json:"telephone_number"`
	EMail           string `json:"e_mail"`
}

func CheckUser(username, password string) int {
	var user User
	db.Select("role").Where(&User{Username: username, Password: password}).First(&user)
	return user.Role
}

func ExistUserByUsername(username string) bool {
	var user User
	db.Select("id").Where(&User{Username: username}).First(&user)
	fmt.Println(user.ID)
	if user.ID > 0 {
		return true
	}
	return false
}

func AddUser(user User) bool {
	db.Create(&user)
	return true
}

func UpdateUser(user User) bool {
	return !db.Model(&User{}).Update(&user).RecordNotFound()
}

func DeleteUser(id int) bool {
	return !db.Where("id = ?", id).Delete(&User{}).RecordNotFound()
}

func GetUserByUsername(username string) (user User, flag bool) {
	flag = !db.Where("username = ?", username).First(&user).RecordNotFound()
	return
}

func GetUserTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)
	return
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (users []User, flag bool) {
	if pageNum > 0 && pageSize > 0 {
		flag = !db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).RecordNotFound()
	} else {
		flag = !db.Where(maps).Find(&users).RecordNotFound()
	}
	return
}
