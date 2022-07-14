package main

type User struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Age       int    `json:"age" bson:"age"`
}

func (u User) GetUsers() {

}

func (u User) CreateUser(usr User) {

}

func (u User) DeleteUser(usr User) {

}
