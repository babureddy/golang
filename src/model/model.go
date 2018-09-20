package model

import (
	"database/sql"
	"time"
)

// /Users/cabrera/Desktop/code/gpsnow-api-server/src/github.com/Shelnutt2/db2struct/db2struct/db2struct --host gpsnow.cksiqniek8yk.ap-south-1.rds.amazonaws.com --database gpsnow -t vehicle_type --struct VehicleType -p Cabrera2018 --user gpsnow

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
	UserType  string
}

type Availability struct {
	Available  int
	CreatedAt  time.Time
	DayOfMonth sql.NullString
	DayOfWeek  sql.NullString
	FromDate   time.Time
	ID         int64
	Month      sql.NullString
	ProfileID  int64
	ToDate     time.Time
}

type Awards_Certification struct {
	Authority   string
	CreatedAt   time.Time
	Description string
	ID          int64
	Name        string
	ProfileID   int64
	URL         sql.NullString
	Validity    string
	Verified    int
}

type Category struct {
	Count          int
	CreatedAt      time.Time
	ID             int64
	IndustryID     int64
	Name           string
	Transportation int
	Verified       int
}

type Industry struct {
	Count          int
	CreatedAt      time.Time
	ID             int64
	ImageURL       string
	Name           string
	Transportation int
	Verified       int
}

type CategoryCustomization struct {
	CategoryID int64
	CreatedAt  time.Time
	ID         int64
	UserID     int64
}

type Company struct {
	AdminName      string
	CreatedAt      time.Time
	ID             int64
	Name           string
	Phone          string
	RegistrationNo sql.NullString
	Website        sql.NullString
}

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

type Education struct {
	CreatedAt        time.Time
	ID               int64
	Name             string
	ProfileID        int64
	Qualification    string
	Verified         int
	YearOfCompletion int
}

type Experience struct {
	CompanyName string
	CreatedAt   time.Time
	Description sql.NullString
	FromMonth   int
	FromYear    int
	ID          int64
	ProfileID   int64
	ToMonth     sql.NullInt64
	ToYear      sql.NullInt64
	Verified    int
}

type Expertise struct {
	CreatedAt time.Time
	ID        int64
	Name      string
	ProfileID int64
	Verified  int
}

type Feed struct {
	CreatedAt       time.Time
	Description     string
	FeedType        string
	FollowingMe     int
	FollowingOthers int
	ID              int64
	Reason          sql.NullString
	Status          int
	Title           string
	UserID          int64
}
type FeedShare struct {
	CreatedAt        time.Time
	FeedID           int64
	ID               int64
	SharedWithUserID int64
	SharerUserID     int64
}

type FeedStatistics struct {
	Action    string
	Comment   sql.NullString
	CreatedAt time.Time
	FeedID    int64
	ID        int64
	UserID    int64
}

type Group struct {
	CreatedAt time.Time
	ID        int64
	ImageURL  sql.NullString
	Name      string
	UserID    int64
}

type IndustryCustomization struct {
	CreatedAt  time.Time
	ID         int64
	IndustryID int64
	UserID     int64
}

type Item struct {
	CreatedAt            time.Time
	Currency             string
	Description          string
	ID                   int64
	Name                 string
	Price                float64
	ProfileID            int64
	Qty                  int
	ShowOnFeedsGpsEmploy int
	ShowOnFeedsGpsWork   int
	Title                string
}

type Language struct {
	CreatedAt time.Time
	ID        int64
	Name      string
	ProfileID int64
}

type LearnMore struct {
	CreatedAt time.Time
	ID        int64
	Name      string
}

type Login struct {
	AppName         string
	ID              int64
	Lat             float64
	Lng             float64
	LoginTimestamp  time.Time
	LogoutTimestamp time.Time
	UserID          int64
}
type MapSettings struct {
	CreatedAt       time.Time
	Freelance       int
	ID              int64
	OpenForRequests int
	Privacy         int
	ProfileID       int64
	Radius          int
	RadiusUnits     string
	ShowOnMap       int
}

type Network struct {
	Action    string
	Comment   sql.NullString
	CreatedAt time.Time
	ID        int64
	ProfileID int64
	UserID    int64
}
type Notification struct {
	CreatedAt  time.Time
	FromUserID int64
	ID         int64
	Message    string
	ToUserID   int64
}

type PaymentCardDetails struct {
	CreatedAt      time.Time
	CreditCardNo   sql.NullString
	CreditCardType sql.NullString
	DebitCardNo    sql.NullString
	ExpiresOnMonth sql.NullString
	ExpiresOnYear  sql.NullString
	ID             int64
	NameOnCard     string
}
type Portfolio struct {
	CreatedAt  time.Time
	DocumentID int64
	ID         int64
	Name       string
	ProfileID  int64
	Verified   int
}
type ProfileGroup struct {
	CreatedAt time.Time
	GroupID   int64
	ID        int64
	ProfileID int64
}
type Request struct {
	CreatedAt       time.Time
	CustomerID      int64
	Description     string
	Destination     sql.NullString
	Distance        sql.NullFloat64
	EndDatetime     time.Time
	Estimate        string
	ID              int64
	ItemID          sql.NullInt64
	ItemsRequired   sql.NullInt64
	Lat             sql.NullFloat64
	Lng             sql.NullFloat64
	ProfileID       sql.NullInt64
	ServiceID       sql.NullInt64
	StartDatetime   time.Time
	State           sql.NullString
	Title           string
	UpdatedAt       time.Time
	Vehicle         sql.NullInt64
	WorkersRequired int
}
type RequestStatistics struct {
	Action    string
	CreatedAt time.Time
	ID        int64
	RequestID int64
	UserID    int64
}
type RequestStatus struct {
	CreatedAt  time.Time
	ID         int64
	RequestID  int64
	RevisionID sql.NullInt64
	Status     string
}
type Revision struct {
	ActualPrice     float64
	CreatedAt       time.Time
	CustomerID      int64
	Description     string
	Destination     sql.NullString
	Distance        sql.NullFloat64
	EndDatetime     time.Time
	Estimate        string
	ID              int64
	ItemID          sql.NullInt64
	ItemsRequired   sql.NullInt64
	Lat             float64
	Lng             float64
	PriceType       string
	ProfileID       sql.NullInt64
	RequestID       int64
	ServiceID       sql.NullInt64
	StartDatetime   time.Time
	Title           string
	UpdatedAt       time.Time
	VehicleType     sql.NullInt64
	WorkersRequired int
}
type Service struct {
	CreatedAt            time.Time
	Currency             string
	Description          string
	ID                   int64
	Name                 string
	Price                float64
	PriceType            string
	ProfileID            int64
	ShowInProfileView    int
	ShowOnFeedsGpsEmploy int
	ShowOnFeedsGpsWork   int
	Title                string
}
type Transactions struct {
	Amount        float64
	CardID        sql.NullString
	CreatedAt     time.Time
	CustomerID    int64
	ID            int64
	PaymentMethod string
	ProfileID     int64
	TransactionID string
	WorkerID      int64
}
type UserProfile struct {
	About       string
	CategoryID  int64
	Count       int
	CreatedAt   time.Time
	Currency    string
	Experience  int
	ID          int64
	Price       float64
	PriceType   string
	Rating      float64
	Title       string
	UserID      int64
	VehicleType int64
}
type UserProfileStatistics struct {
	Action    string
	CreatedAt time.Time
	ID        int64
	ProfileID int64
	UserID    int64
}
type VehicleType struct {
	CreatedAt time.Time
	ID        int64
	Name      string
	Tarriff   float64
	Units     string
	Verified  int
}
