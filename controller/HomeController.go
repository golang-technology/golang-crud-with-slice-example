package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Id       uint64
	Username string
	Password string
	Fullname string
}

var db = []Data{
	{1, "haipv", "hai@123", "Phan Van Hai"},
	{2, "maipt", "mai@123", "Phan Thi Mai"},
	{3, "tuannv", "tuan@123", "Nguyen Van Tuan"},
	{4, "dungnv", "dung@123", "Nguyen Van Dung"},
	{5, "anhpt", "anh@123", "Phan Thi Anh"},
}

var (
	key   = []byte("super-scret-key")
	store = sessions.NewCookieStore(key)
)

// Home
func Home(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("view/home.html"))
	tmp.Execute(w, struct {
		DB [] Data
	}{db})
}

// Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.ParseUint(vars["id"], 10, 32)
	tmp := template.Must(template.ParseFiles("view/create-user.html"))
	for _, data := range db {
		if data.Id == uid {
			tmp.Execute(w, data)
			return
		}
	}
	tmp.Execute(w, nil)
}

// Create
func SaveUser(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("view/create-user.html"))
	tmp.Execute(w, nil)
}

func DoSaveUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fullname := r.FormValue("fullname")
	id := uint64(len(db) + 1)
	ls := Data{id, username, password, fullname}
	db = append(db, ls)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	fmt.Println(db)
}

// Update User
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.ParseUint(vars["id"], 10, 32)
	tmp := template.Must(template.ParseFiles("view/update-user.html"))
	for _, data := range db {
		if data.Id == uid {
			tmp.Execute(w, data)
			return
		}
	}
	tmp.Execute(w, nil)
}

// Do Update User
func DoUpdateUser(w http.ResponseWriter, r *http.Request) {
	//id := r.FormValue("id")
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	username := r.FormValue("username")
	password := r.FormValue("password")
	fullname := r.FormValue("fullname")
	for i, data := range db {
		if data.Id == id {
			db[i].Id = id
			db[i].Username = username
			db[i].Password = password
			db[i].Fullname = fullname
			break
		}
	}
	fmt.Println(id, "&&", db)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Details
func DetailUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.ParseUint(vars["id"], 10, 32)
	tmp := template.Must(template.ParseFiles("view/details-user.html"))
	for _, data := range db {
		if data.Id == uid {
			tmp.Execute(w, data)
			return
		}
	}
	tmp.Execute(w, nil)
}

func DeteleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.ParseUint(vars["id"], 10, 32)
	for i, data := range db {
		if data.Id == uid {
			db = append(db[:i], db[i+1:]...)
			break
		}
	}
	fmt.Println(db)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
