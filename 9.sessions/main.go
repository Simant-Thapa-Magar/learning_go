package main

import (
	"html/template"
	"net/http"
)

var tpl template.Template

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./templates/index.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	loginStatus := isLoggedIn(req)

	tpl.ExecuteTemplate(w, "index.gohtml", loginStatus)
}

func signup(w http.ResponseWriter, req *http.Request) {
	loginStatus := isLoggedIn(req)

	if loginStatus {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		newUser := createUser(req)

		if newUser {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
		}
	}

	tpl, err := template.ParseFiles("./templates/signup.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", "")
}

func login(w http.ResponseWriter, req *http.Request) {
	loginStatus := isLoggedIn(req)

	if loginStatus {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		isValidUser := attemptLogin(w, req)

		if isValidUser {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
	}

	tpl, err := template.ParseFiles("./templates/login.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", "")
}

func profile(w http.ResponseWriter, req *http.Request) {
	loginStatus := isLoggedIn(req)

	if !loginStatus {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	tpl, err := template.ParseFiles("./templates/profile.gohtml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(w, req)

	tpl.ExecuteTemplate(w, "profile.gohtml", user)
}

func logout(w http.ResponseWriter, req *http.Request) {
	loginStatus := isLoggedIn(req)

	if !loginStatus {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	logUserOut(w, req)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
