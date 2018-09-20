package usr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	secu "../secu"

	DB "../db"
)

type User struct {
	CompanyID int64
	CreatedAt time.Time
	Dob       time.Time
	Email     string
	FirstName string
	ID        int64
	LastName  string
	Lat       float64
	Lng       float64
	Password  string
	Phone     string
	Salt      string
}

type UserLogin struct {
	AppName         string
	ID              int64
	Lat             float64
	Lng             float64
	LoginTimestamp  time.Time
	LogoutTimestamp time.Time
	UserID          int64
}

var db = DB.Db()

func Login(w http.ResponseWriter, r *http.Request) {
	var validUser User
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	} else {
		d := secu.Decrypt(string(body))
		json.NewDecoder(strings.NewReader(d)).Decode(&validUser)
		db.Where("email = ? and password=?", validUser.Email, validUser.Password).Find(&user)
		if user.ID > 0 {
			fmt.Println(user)
			db.First(&user, user.ID)
			var login UserLogin
			login.UserID = user.ID
			login.AppName = "Worker"
			login.Lat = user.Lat
			login.Lng = user.Lng
			db.Create(&login)
			json.NewEncoder(w).Encode(&user)
		}
	}

}

func Users(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var newKey = secu.Decrypt(params["id"])
	fmt.Println("decrypted key=" + newKey)
	db.First(&user, newKey)
	json.NewEncoder(w).Encode(&user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(body))
		d := secu.Decrypt(string(body))
		fmt.Println(d)

		json.NewDecoder(strings.NewReader(d)).Decode(&user)
		db.Create(&user)
		json.NewEncoder(w).Encode(&user)
	}

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(body))
		d := secu.Decrypt(string(body))
		fmt.Println(d)

		json.NewDecoder(strings.NewReader(d)).Decode(&user)
		db.Update(&user)
		json.NewEncoder(w).Encode(&user)
	}

}
