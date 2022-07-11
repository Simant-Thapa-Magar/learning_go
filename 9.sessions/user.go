package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Firstname string
	Lastname  string
	Username  string
	Password  []byte
}

var sessionUser = map[string]string{}

var userData = map[string]User{}

const cookieKey = "sessionId"

func isLoggedIn(req *http.Request) bool {
	id, err := req.Cookie(cookieKey)

	if err != nil {
		return false
	}

	if uId, ok := sessionUser[id.Value]; ok {
		log.Println("User is logged in ", uId)
		return true
	}

	return false
}

func createUser(req *http.Request) bool {
	fname := req.FormValue("first_name")
	lname := req.FormValue("last_name")
	username := req.FormValue("username")
	password := req.FormValue("password")

	if fname == "" || lname == "" || username == "" || password == "" {
		return false
	}

	for _, u := range userData {
		if u.Username == username {
			return false
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return false
	}

	newUser := User{
		fname,
		lname,
		username,
		hashedPassword,
	}

	fmt.Println("user ", newUser)

	userId := uuid.NewV4()

	fmt.Println("user id ", userId)

	userData[userId.String()] = newUser

	return true

}

func attemptLogin(w http.ResponseWriter, req *http.Request) bool {
	var usr User
	var userId string
	userFound := false
	username := req.FormValue("username")
	password := req.FormValue("password")

	if username == "" || password == "" {
		return false
	}

	for key, u := range userData {
		if u.Username == username {
			userId = key
			usr = u
			userFound = true
			break
		}
	}

	if !userFound {
		return false
	}

	err := bcrypt.CompareHashAndPassword(usr.Password, []byte(password))

	if err != nil {
		return false
	}

	userSessionKey := uuid.NewV4()

	sessionUser[userSessionKey.String()] = userId

	userCookie := &http.Cookie{
		Name:  cookieKey,
		Value: userSessionKey.String(),
	}

	http.SetCookie(w, userCookie)
	return true
}

func logUserOut(w http.ResponseWriter, req *http.Request) {
	id, err := req.Cookie(cookieKey)

	if err != nil {
		return
	}

	if uId, ok := sessionUser[id.Value]; ok {
		log.Println("Current logged in User ", uId)
		cookie := &http.Cookie{
			Name:   cookieKey,
			Value:  "",
			MaxAge: -1,
		}
		delete(sessionUser, id.Value)
		http.SetCookie(w, cookie)
	}

}

func getUser(w http.ResponseWriter, req *http.Request) User {
	var usr User
	id, err := req.Cookie(cookieKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if uId, ok := sessionUser[id.Value]; ok {
		if u, ook := userData[uId]; ook {
			usr = u
		}
	}

	return usr
}
