package model

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Role            int    `json:"role"`
	TelephoneNumber string `json:"telephone_number"`
	EMail           string `json:"e_mail"`
}

func CheckUser(username, password string) bool {
	var user User
	db.Select("id").Where(&User{Username: username, Password: password}).First(&user)
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
	return db.Model(&User{}).Update(&user).RecordNotFound()
}

func DeleteUser(id int) bool {
	return db.Where("id = ?", id).Delete(&User{}).RecordNotFound()
}
