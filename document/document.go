package usr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	secu "../secu"

	DB "../db"
)

var err error

type Document struct {
	CreatedAt    time.Time
	DocumentName string
	DocumentType string
	DocumentURL  string
	ID           int64
	ModuleID     int64
	ModuleName   int64
	TableName    string
	Verified     int
}

type AWSBucket struct {
	Bucket   string
	Filename string
}

var db = DB.Db()

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	var documents []Document
	params := mux.Vars(r)
	var newKey = secu.Decrypt(params["id"])
	fmt.Println(newKey)
	db.Where("profile_id = ? ", newKey).Find(&documents)
	json.NewEncoder(w).Encode(&documents)
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	jsonString := secu.Decrypt(string(body))
	var document Document
	decoder := json.NewDecoder(bytes.NewBufferString(jsonString))
	err = decoder.Decode(&document)
	if err != nil {
		panic(err)
	}
	db.Create(&document)
	w.WriteHeader(http.StatusOK)
}

func Upload(c *gin.Context) {

	var link [5]string

	multipart, err := c.Request.MultipartReader()
	if err != nil {
		log.Fatalln("Failed to create MultipartReader", err)
	}

	for i := range link {
		mimePart, err := multipart.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading multipart section: %v", err)
			break
		}
		_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
		if err != nil {
			log.Printf("Invalid Content-Disposition: %v", err)
			break
		}

		uploader := s3manager.NewUploader(session.New(&aws.Config{Region: aws.String("us-east-1")}))

		result, err := uploader.Upload(&s3manager.UploadInput{
			Body:        mimePart,
			Bucket:      aws.String("verificationme"),
			Key:         aws.String(params["filename"]),
			ContentType: aws.String(mimePart.Header.Get("Content-Type")),
			ACL:         aws.String("public-read"),
		})
		if err != nil {
			log.Fatalln("Failed to upload to S3", err)
		}

		link[i] = result.Location

	}

	c.JSON(http.StatusOK, gin.H{
		"Status":   "uploaded",
		"document": link,
	})

}
