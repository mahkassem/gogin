package services

type User struct {
	Id   int
	Name string
	Age  int
}

var db = []User{
	{Id: 1, Name: "MostafaAdly", Age: 22},
	{Id: 2, Name: "MahmoudKassem", Age: 35},
	{Id: 3, Name: "HossamYoussef", Age: 37},
}

func GetUserById(id int) (User, bool) {
	for _, user := range db {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}

func GetAllUsers() ([]User, bool) {
	return db, true
}
