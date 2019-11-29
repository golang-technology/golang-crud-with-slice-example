package main

import (
	controller "./controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	/*	fs := http.FileServer(http.Dir("static"))
		http.Handle("/static/css", http.StripPrefix("/static/", fs))*/
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/create",controller.SaveUser).Methods("GET")
	r.HandleFunc("/create",controller.DoSaveUser).Methods("POST")
	r.HandleFunc("/update/{id}",controller.UpdateUser).Methods("GET")
	r.HandleFunc("/update/{id}",controller.DoUpdateUser).Methods("POST")
	r.HandleFunc("/details/{id}",controller.DetailUser).Methods("GET")
	r.HandleFunc("/delete/{id}",controller.DeteleUser)

	//http.Handle(r)
	http.ListenAndServe(":9090", r)
	//fmt.Println("Hello")
}
