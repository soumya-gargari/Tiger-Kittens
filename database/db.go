package database

import (
	"Tiger-Kittens/data"
	"database/sql"
	"errors"
	"fmt"
)

type Database struct {
	Db *sql.DB
}

// dnsName is nothing but giving credentials to sql to connect to the database
// for ex: root:password@tcp(127.0.0.1:3306)/{dbName}
// database name is optional
func dnsName() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", data.UserName, data.PassWord, data.Hostname, data.Dbname)
}

func (d *Database) InitDatabse() *sql.DB {
	// creating connection with database
	dnsHost := dnsName()
	fmt.Println(dnsHost)
	database, err := sql.Open("mysql", dnsHost)
	if err != nil {
		fmt.Println("failed to establish sql connection", err)
		return nil
	}
	err = database.Ping()
	if err == nil {
		fmt.Println("sql connection created successfully")
	}
	database.SetMaxIdleConns(20)
	database.SetMaxOpenConns(20)
	resp, err := database.Exec("CREATE DATABASE IF NOT EXISTS " + data.Dbname)
	if err != nil {
		fmt.Println("failed to create database", err)
		return nil
	}
	rows, _ := resp.RowsAffected()
	fmt.Println("num of rows affected:", rows)
	return database
}

func (d *Database) CreateUserTable(tableName string) error {
	d.Db = d.InitDatabse()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(username VARCHAR(255), password VARCHAR(255), email VARCHAR(255))", tableName)
	_, err := d.Db.Exec(query)
	if err != nil {
		fmt.Println("failed to create table with rows", err)
		return err
	}
	fmt.Println("Sucessfully created table:", tableName)
	return nil
}

func (d *Database) InsertUserData(tableName string, userData data.UserDetails) error {
	exists, err := recordExists(d.Db, tableName, "username", userData.UserName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("failed to insert data, duplicate")
	} else {
		query := fmt.Sprintf("INSERT INTO %s VALUES('%s','%s','%s')", tableName, userData.UserName, userData.PassWord, userData.Email)
		fmt.Println("query statement for insertion is:", query)
		_, err := d.Db.Query(query)
		if err != nil {
			fmt.Println("failed to insert data", err)
			return err
		}
	}
	return nil
}

func (d *Database) GetUserData(tableName, userName, passWord string) (data.UserDetails, error) {
	var userData data.UserDetails
	query := fmt.Sprintf("SELECT * FROM %s WHERE username='%s' AND password='%s'", tableName, userName, passWord)
	fmt.Println("query statement for insertion is:", query)
	rows, err := d.Db.Query(query)
	if err != nil {
		fmt.Println("failed to get data from mysql table", err)
		return userData, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userData.UserName, &userData.PassWord, &userData.Email)
		if err != nil {
			fmt.Println("failed to get data from mysql table", err)
			return userData, err
		}
	}
	return userData, nil
}

func (d *Database) CreateTigerInfoTable(tableName string) error {
	d.Db = d.InitDatabse()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(name VARCHAR(255), dateOfBirth VARCHAR(255), lastSeen TIMESTAMP DEFAULT CURRENT_TIMESTAMP, lastSeenlatitude VARCHAR(255), lastSeenlongitude VARCHAR(255))", tableName)
	fmt.Println("query statement for creation of table is:", query)
	_, err := d.Db.Exec(query)
	if err != nil {
		fmt.Println("failed to create table with rows", err)
		return err
	}
	fmt.Println("Sucessfully created table:", tableName)
	return nil
}

func (d *Database) CreateTigerSightingInfoTable(tableName string) error {
	d.Db = d.InitDatabse()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(username VARCHAR(255), tigername VARCHAR(255), latitude VARCHAR(255), longitude VARCHAR(255), timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, uploadImage BOOL)", tableName)
	fmt.Println("query statement for creation of table is:", query)
	_, err := d.Db.Exec(query)
	if err != nil {
		fmt.Println("failed to create table with rows", err)
		return err
	}
	fmt.Println("Sucessfully created table:", tableName)
	return nil
}

func (d *Database) InsertTigerData(tableName string, tigerData data.TigerDetails) error {
	exists, err := recordExists(d.Db, tableName, "name", tigerData.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("failed to insert data, duplicate")
	} else {
		_, err := d.Db.Exec("INSERT INTO tigerdetails (name, dateOfBirth, lastSeen, lastSeenlatitude, lastSeenlongitude) VALUES (?, ?, ?, ?, ?)", tigerData.Name, tigerData.DateOfBirth, tigerData.LastSeen, tigerData.Latitude, tigerData.Longitude)
		if err != nil {
			fmt.Println("failed to insert data", err)
			return err
		}
	}
	return nil
}

func (d *Database) InsertTigerSightingData(tableName string, tigerSightingData data.TigerSightingDetails) error {
	// exists, err := recordExists(d.Db, tableName, "tigername", tigerSightingData.TigerName)
	// if err != nil {
	// 	return err
	// }
	// if exists {
	// 	return errors.New("failed to insert data, duplicate")
	// } else {
	_, err := d.Db.Exec("INSERT INTO tigersighting (userName, tigerName, timestamp, latitude, longitude, uploadImage) VALUES (?, ?, ?, ?, ?, ?)", tigerSightingData.UserName, tigerSightingData.TigerName, tigerSightingData.TimeStamp, tigerSightingData.Latitude, tigerSightingData.Longitude, tigerSightingData.UploadImage)
	if err != nil {
		fmt.Println("failed to insert data", err)
		return err
	}

	return nil
}

func (d *Database) GetTigersData(tableName string) ([]data.TigerDetails, error) {
	var tigersDetails []data.TigerDetails
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY lastSeen", data.TigerInfoTableName)
	rows, err := d.Db.Query(query)
	if err != nil {
		fmt.Println("failed to get data from mysql table", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tigerDetails data.TigerDetails
		err := rows.Scan(&tigerDetails.Name, &tigerDetails.DateOfBirth, &tigerDetails.LastSeen, &tigerDetails.Latitude, &tigerDetails.Longitude)
		if err != nil {
			fmt.Println("failed to get data from mysql table", err)
			return nil, err
		}
		tigersDetails = append(tigersDetails, tigerDetails)
	}
	return tigersDetails, nil
}

func (d *Database) GetAllTigerSightingData(tableName string) ([]data.TigerSightingDetails, error) {
	var tigersSightingDetails []data.TigerSightingDetails
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY timestamp", data.TigerSightingInfoTableName)
	rows, err := d.Db.Query(query)
	if err != nil {
		fmt.Println("failed to get data from mysql table", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tigerSightDetails data.TigerSightingDetails
		err := rows.Scan(&tigerSightDetails.UserName, &tigerSightDetails.TigerName, &tigerSightDetails.Latitude, &tigerSightDetails.Longitude, &tigerSightDetails.TimeStamp, &tigerSightDetails.UploadImage)
		if err != nil {
			fmt.Println("failed to get data from mysql table", err)
			return nil, err
		}
		tigersSightingDetails = append(tigersSightingDetails, tigerSightDetails)
	}
	return tigersSightingDetails, nil
}

func (d *Database) GetTigerData(tableName string, tigerName string) (data.TigerDetails, error) {
	var tigerDetails data.TigerDetails
	query := fmt.Sprintf("SELECT * FROM %s WHERE name='%s'", tableName, tigerName)
	rows, err := d.Db.Query(query)
	if err != nil {
		fmt.Println("failed to get data from mysql table", err)
		return tigerDetails, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tigerDetails.Name, &tigerDetails.DateOfBirth, &tigerDetails.LastSeen, &tigerDetails.Latitude, &tigerDetails.Longitude)
		if err != nil {
			fmt.Println("failed to get data from mysql table", err)
			return tigerDetails, err
		}
	}
	return tigerDetails, nil
}

func recordExists(db *sql.DB, tableName, columnName, value string) (bool, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", tableName, columnName)
	var count int
	err := db.QueryRow(query, value).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d *Database) GetAllReportedUsers(tableName, tigerName string) ([]string, error) {
	var usernames []string
	query := fmt.Sprintf("SELECT username FROM %s where tigername='%s'", data.TigerSightingInfoTableName, tigerName)
	rows, err := d.Db.Query(query)
	if err != nil {
		fmt.Println("failed to get data from mysql table", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			fmt.Println("failed to get data from mysql table", err)
			return nil, err
		}
		usernames = append(usernames, username)
	}
	return usernames, nil
}
