package data

import "time"

// sql config details
const (
	UserName                   = "root"
	PassWord                   = "Soumya@1234"
	Hostname                   = "127.0.0.1:3306"
	Dbname                     = "mysql"
	UserTableName              = "userdetails"
	TigerInfoTableName         = "tigerdetails"
	TigerSightingInfoTableName = "tigersighting"
)

type UserDetails struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
}

type TigerDetails struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
	// Pass the timestamp as ISO 8601 format in req body
	// for ex: "timestamp": "2015-07-05T22:16:18Z"
	LastSeen  time.Time `json:"lastSeen"`
	Latitude  string    `json:"lastSeenlatitude"`
	Longitude string    `json:"lastSeenlongitude"`
}

type TigerSightingDetails struct {
	Name string `json:"name"`
	// Pass the timestamp as ISO 8601 format in req body
	// for ex: "timestamp": "2015-07-05T22:16:18Z"
	TimeStamp   time.Time `json:"timestamp"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	UploadImage bool      `json:"uploadImage"`
}
