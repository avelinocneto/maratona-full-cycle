package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Driver struct {
	ID   string `json:"uuid"`
	Name string `json:"name"`
}

var drivers []Driver

func convertToObj(bytes []byte) {

	var obj Driver
	if err := json.Unmarshal(bytes, &obj); err != nil {
		panic(err)
	}

	drivers = append(drivers, obj)
}

func loadDrivers(file string) []byte {

	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func listDrivers(w http.ResponseWriter, r *http.Request) {
	//driverList := loadDrivers("drivers.json")
	//w.Write([]byte(driverList))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

func findDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range drivers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Driver{})
}

func main() {
	convertToObj(loadDrivers("drivers.json"))
	r := mux.NewRouter()
	r.HandleFunc("/drivers", listDrivers)
	r.HandleFunc("/drivers/{id}", findDriver)
	http.ListenAndServe(":8081", r)
}
