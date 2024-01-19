package tiger

import (
	"Tiger-Kittens/data"
	"Tiger-Kittens/database"
	"Tiger-Kittens/distance"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// CreateTiger method for creating tiger
func CreateTiger(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	var tigerDetails data.TigerDetails
	err = json.Unmarshal(body, &tigerDetails)
	if err != nil {
		http.Error(w, "failed to unmarshal request body", http.StatusBadRequest)
		return
	}
	var dB database.Database
	err = dB.CreateTigerInfoTable(data.TigerInfoTableName)
	if err != nil {
		http.Error(w, "failed to create tigerdetails table", http.StatusInternalServerError)
		return
	}
	err = dB.InsertTigerData(data.TigerInfoTableName, tigerDetails)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to save data to tigerdetails Table", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(http.StatusCreated)
}

// GetAllTigerDetails method for getting all tiger details
func GetAllTigerDetails(w http.ResponseWriter, r *http.Request) {
	db := &database.Database{}
	mySql := db.InitDatabse()
	db.Db = mySql
	result, err := db.GetTigersData(data.TigerInfoTableName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to get all tiger data from tigerdetails Table", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// CreateTigerSighting method for creating tiger
func CreateTigerSighting(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}
	var tigerSightDetails data.TigerSightingDetails
	err = json.Unmarshal(body, &tigerSightDetails)
	if err != nil {
		http.Error(w, "failed to unmarshal request body", http.StatusBadRequest)
		return
	}
	var dB database.Database

	err = dB.CreateTigerSightingInfoTable(data.TigerSightingInfoTableName)
	if err != nil {
		http.Error(w, "failed to create tigersighting table", http.StatusInternalServerError)
		return
	}
	tigerData, err := dB.GetTigerData(data.TigerInfoTableName, tigerSightDetails.Name)
	fmt.Println("error of getting tiger data", err)
	fmt.Println("tiger data is,", tigerData)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to get tiger data from tigerdetails Table", http.StatusInternalServerError)
		return
	}
	if tigerData.Latitude != "" && tigerData.Longitude != "" {
		lat1, _ := strconv.ParseFloat(tigerSightDetails.Latitude, 64)
		lon1, _ := strconv.ParseFloat(tigerSightDetails.Longitude, 64)
		lat2, _ := strconv.ParseFloat(tigerData.Latitude, 64)
		lon2, _ := strconv.ParseFloat(tigerData.Longitude, 64)
		fmt.Println("latitude longitudes are:", lat1, lon1, lat2, lon2)
		dist := distance.CalculateDistance(lat1, lon1, lat2, lon2)
		if dist > 5 {
			err = dB.InsertTigerSightingData(data.TigerInfoTableName, tigerSightDetails)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "failed to save data to tigersighting Table", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "tiger sighting is within 5 kilometres of its previous sighting", http.StatusInternalServerError)
			return
		}

	}

	json.NewEncoder(w).Encode(http.StatusCreated)
}

// GetAllTigerSightingDetails method for getting all tiger sighting details
func GetAllTigerSightingDetails(w http.ResponseWriter, r *http.Request) {
	db := &database.Database{}
	mySql := db.InitDatabse()
	db.Db = mySql
	result, err := db.GetAllTigerSightingData(data.TigerSightingInfoTableName)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to get all tiger sighting data from tigersighting Table", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
