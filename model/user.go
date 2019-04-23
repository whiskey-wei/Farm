package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func checkUser(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
