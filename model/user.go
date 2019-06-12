package model

import "fmt"

type User struct {
	ID              int    `json:"id"`						//id
	Username        string `json:"username"`				//用户名
	Password        string `json:"password"`				//密码
	Farmname		string `json:"farmname"`				//农场名
	Role            int    `json:"role"`					//职级，0为管理员，1为员工
	PersonalName	string `json:"personal_name"`			//姓名
	TelephoneNumber string `json:"telephone_number"`		//电话
	EMail           string `json:"e_mail"`					//邮箱
}

func CheckUser(username string, password string) int {
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
		flag = !db.Where(maps).Find(&users).Offset(pageNum).Limit(pageSize).RecordNotFound()
	} else {
		flag = !db.Where(maps).Find(&users).RecordNotFound()
	}
	return
}
