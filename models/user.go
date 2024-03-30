package models

type User struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

func CreateUser(name string, location Location) User {

	return User{
		Name:     name,
		Location: location,
	}
}
