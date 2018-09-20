package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	language "./language"
	userprofile "./userprofile"
	usr "./usr"

	secu "./secu"
)

var err error

func main() {
	// testEncryption()
	router := mux.NewRouter()

	router.HandleFunc("/d/{id}", d).Methods("GET")
	router.HandleFunc("/e/{id}", e).Methods("GET")
	router.HandleFunc("/user", usr.Users).Methods("GET")
	router.HandleFunc("/user/{id}", usr.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", usr.UpdateUser).Methods("PUT")
	router.HandleFunc("/user", usr.CreateUser).Methods("POST")
	router.HandleFunc("/login", usr.Login).Methods("POST")

	router.HandleFunc("/userprofile/{id}", userprofile.GetUserProfile).Methods("GET")
	router.HandleFunc("/userprofile", userprofile.CreateUserProfile).Methods("POST")
	router.HandleFunc("/userprofile/{id}", userprofile.UpdateUserProfile).Methods("PUT")
	router.HandleFunc("/userprofiles/{id}", userprofile.GetUserProfiles).Methods("GET")

	router.HandleFunc("/language/{id}", language.GetLanguages).Methods("GET")
	router.HandleFunc("/language", language.CreateLanguage).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", router))
}

func e(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	e, err := secu.Encrypt(params["id"])
	if err != nil {
		panic(err)
	} else {
		fmt.Println(e)
	}
}
func d(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	e := secu.Decrypt(params["id"])
	fmt.Println(e)
}
