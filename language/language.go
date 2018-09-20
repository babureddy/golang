package usr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	secu "../secu"

	DB "../db"
)

type Language struct {
	CreatedAt time.Time
	ID        int64
	Name      string
	ProfileID int64
}

var db = DB.Db()

func GetLanguages(w http.ResponseWriter, r *http.Request) {
	var languages []Language
	params := mux.Vars(r)
	var newKey = secu.Decrypt(params["id"])
	fmt.Println(newKey)
	db.Where("profile_id = ? ", newKey).Find(&languages)
	json.NewEncoder(w).Encode(&languages)
}

func CreateLanguage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	jsonString := secu.Decrypt(string(body))
	// jsonString := `[{"ProfileID":9,"Name":"Spanish"}]`
	result := make([]Language, 0)
	decoder := json.NewDecoder(bytes.NewBufferString(jsonString))
	err = decoder.Decode(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result[0].ProfileID, result[0].Name)
	// var oldLanguages Language
	// db.Where("profile_id = ?", result[0].ProfileID).Find(&oldLanguages)
	// fmt.Println(oldLanguages)
	tx := db.Begin()
	// db.Where("profile_id = ?", result[0].ProfileID).Delete(Language{})
	deleteQuery := "DELETE FROM language WHERE profile_id=?;"
	_ = tx.Exec(deleteQuery, result[0].ProfileID).Error
	query := "INSERT INTO language(profile_id, name) VALUES (?,?);"
	for x, value := range result {
		fmt.Println(x, ":", value)
		_ = tx.Exec(query, value.ProfileID, value.Name).Error
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)

}
