package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Structure Definitions

//ResponsePayload defines a successful search response
type ResponsePayload struct {
	Term   string  `json:"searchTerm"`
	Result string  `json:"matchedValue"`
	Data   AppData `json:"matchedApp"`
}

//Maintainer defines a single maintainer nested in JSON
type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//AppData maps app meta data from JSON
type AppData struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Version     string       `json:"version"`
	Maintainers []Maintainer `json:"maintainers"`
	Company     string       `json:"company"`
	Website     string       `json:"website"`
	Source      string       `json:"source"`
	License     string       `json:"license"`
	Description string       `json:"description"`
}

//Globals

//Persists app meta data
var apps []AppData

//Temporarily search results
var results []ResponsePayload

//Validators

//Validate AppData struct
func (a AppData) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required),
		validation.Field(&a.Version, validation.Required),
		validation.Field(&a.Maintainers, validation.Required),
		validation.Field(&a.Source, validation.Required, is.URL),
		validation.Field(&a.Company, validation.Required),
		validation.Field(&a.License, validation.Required),
		validation.Field(&a.Description, validation.Required),
		validation.Field(&a.Website, validation.Required, is.URL),
	)
}

//Validate maintainer struct
func (a Maintainer) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

//CRUD Operators --abridged

//AddApps adds App metadata to the apps object - POST
func AddApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Individual records stored/ queried by ID
	var appID = uuid.Must(uuid.NewRandom()).String()
	var app AppData
	_ = json.NewDecoder(r.Body).Decode(&app)
	app.ID = appID
	//payload validation
	appErr := app.Validate()
	if appErr == nil {
		apps = append(apps, app)
		json.NewEncoder(w).Encode(app)
	} else {
		json.NewEncoder(w).Encode(appErr)
	}
}

//SearchApps is the search feature.  I intentionally shifted
// the search functionality to the Golang layer to spend more time with the language.
// the frontend sends the query string that is passed to the 'api/v1/search'
// endpoint and parses the result from this function.
func SearchApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()["q"]
	for j, s := range apps {
		var appSearch interface{} = s
		app := reflect.ValueOf(appSearch)
		for _, q := range query {
			for i := 0; i < app.NumField(); i++ {
				appItem := app.Field(i).String()
				if strings.Contains(strings.ToLower(appItem), strings.ToLower(q)) {
					results = append(results, ResponsePayload{q, appItem, apps[j]})
				}
			}
		}
	}
	if results == nil {
		json.NewEncoder(w).Encode("no results")
	} else {
		json.NewEncoder(w).Encode(results)
		results = nil
	}
}

//GetApps fetches all app metadata - GET
func GetApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apps)
}

//GetApp fetches one App by appID (/{id} param) -  GET
func GetApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, app := range apps {
		err := errors.New("ID not found")
		if err != nil {
			fmt.Print(err)
			json.NewEncoder(w).Encode(err)

		} else {
			return
		}
		if app.ID == params["id"] {
			json.NewEncoder(w).Encode(app)
			return
		}
	}
	json.NewEncoder(w).Encode(&AppData{})
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/apps", AddApps).Methods("POST")
	r.HandleFunc("/api/v1/apps", GetApps).Methods("GET")
	r.HandleFunc("/api/v1/apps/{id}", GetApp).Methods("GET")
	r.HandleFunc("/api/v1/search", SearchApps).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
