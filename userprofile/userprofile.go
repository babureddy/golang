package usrProfile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	secu "../secu"

	DB "../db"
)

var db = DB.Db()

type UserProfile struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	About       string    `json:"about"`
	Title       string    `json:"title"`
	Experience  int       `json:"experience"`
	UserID      int64     `json:"user_id"`
	Rating      float64   `json:"rating"`
	VehicleType int64     `json:"vehicle_type"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	PriceType   string    `json:"price_type"`
	Count       int       `json:"count"`
}

/*
{"id" : 1,	"category_id" : 1,	"about" : "I am an experienced concrete mixer.","title" : "Concrete Specialist","experience" : 20,	"user_id" : 12,	"rating" : 4.00,"vehicle_type" : 0,	"price" : 20.000,"currency" : "USD","price_type" : "hourly","count" : 0	}
{"id" : 1,	"experience" : 20}
*/
func GetUserProfiles(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var newKey = secu.Decrypt(params["id"])
	fmt.Println("decrypted key=" + newKey)

	var userProfiles []UserProfile
	res := db.Where("user_id = ?", newKey)

	res.Find(&userProfiles)
	json.NewEncoder(w).Encode(&userProfiles)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var newKey = secu.Decrypt(params["id"])

	var userProfile UserProfile
	res := db.Where("id=?", newKey)

	res.Find(&userProfile)
	json.NewEncoder(w).Encode(&userProfile)
}

// {"category_id" : 2,	"about" : "I am an experienced school teacher.","title" : "Math Teacher","experience" : 15,	"user_id" : 12,	"rating" : 4.00,"vehicle_type" : 0,	"price" : 25.000,"currency" : "USD","price_type" : "hourly","count" : 0	}
func CreateUserProfile(w http.ResponseWriter, r *http.Request) {
	var userProfile UserProfile
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(body))
		d := secu.Decrypt(string(body))
		fmt.Println(d)

		err = json.Unmarshal([]byte(d), &userProfile)
		if err != nil {
			panic(err)
		}
		db.Create(&userProfile)
		json.NewEncoder(w).Encode(&userProfile)
	}
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	jsonStream1 := secu.Decrypt(string(body))
	//	const jsonStream1 = `{"id" : 1,	"category_id" : 1,	"about" : "I am an experienced concrete mixer.","title" : "Concrete Specialist","experience" : 20,	"user_id" : 12,	"rating" : 4.00,"vehicle_type" : 0,	"price" : 20.000,"currency" : "USD","price_type" : "hourly","count" : 0	}`
	dec := json.NewDecoder(strings.NewReader(jsonStream1))
	var m UserProfile
	if err := dec.Decode(&m); err == io.EOF {
		panic(err)
	} else if err != nil {
		log.Fatal(err)
	}
	m.CreatedAt = time.Now()
	db.Save(&m)
	json.NewEncoder(w).Encode(&m)
}
