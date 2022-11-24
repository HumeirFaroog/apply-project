package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Objects struct {
// 	ProjectDescriptions []ProjectDescriptions `json:"ProjectDescriptions"`
// 	TeamMembers         []TeamMembers         `json:"TeamMembers"`
// 	RoadMap             []RoadMap             `json:"RoadMap"`
// }

type Objects struct {
	ProjectName             string    `json:"projectName"`
	ProjectDescription      string    `json:"ProjectDescription"`
	TokenAmount             int       `json:"tokenAmount"`
	AddTeamMember           []string  `json:"addTeamMember"`
	TeamMembersRole         []string  `json:"teamMembersRole"`
	TeamMemebersDescription string    `json:"teamMemebersDescription"`
	Options                 [2]string `json:"Description"`
	Percentages             [2]int    `json:"percentages"`
}

// type ProjectDescriptions struct {
// 	projectName        string `json:"projectName"`
// 	ProjectDescription string `json:"ProjectDescription"`
// 	tokenAmount        int    `json:"tokenAmount"`
// }

// type TeamMembers struct {
// 	addTeamMember           string `json:"addTeamMember"`
// 	teamMembersRole         string `json:"teamMembersRole"`
// 	teamMemebersDescription string `json:"teamMemebersDescription"`
// }
// type RoadMap struct {
// 	options     string `json:"Description"`
// 	percentages int    `json:"percentages"`
// }

//MY DATABASE
var objects []Objects

// //MY  HEKPER
// func (c *Objects) IsEmpty() bool {
// 	return c.projectName == ""

// }

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = ""
// 	dbname   = "calhounio_demo"
// )

// func homePage(w http.ResponseWriter, r *http.Request) { //  r: GET  THE VALUE FROM WHOVER SENDING THE REQUEST  && w:  WHEN YOU WRITE THE RESPONSE FOR THE VALUE
// 	w.Write([]byte("<h1>Our Home Page </h1>"))
// }

// func createProject(w http.ResponseWriter, r *http.Request) {

// }

// func checkErorr(err error) {
// 	if err != nil {
// 		panic(err)

// 	}

// }

func main() {
	fmt.Println(" my second task")

	r := mux.NewRouter()

	objects = append(objects, Objects{
		ProjectName:             "prj1",
		ProjectDescription:      "dechhvh1",
		TokenAmount:             568,
		AddTeamMember:           []string{"dbgfnfgng", "eee"},
		TeamMembersRole:         []string{"fdgfhghgf", "coding"},
		TeamMemebersDescription: "here some de",
		Options:                 [2]string{"444", "333"},
		Percentages:             [2]int{45, 46},

		// ProjectDescriptions: []ProjectDescriptions{
		// 	{
		// 		projectName:        "prj1",
		// 		ProjectDescription: "dec1",
		// 		tokenAmount:        568,
		// 	},
		// },
		// TeamMembers: []TeamMembers{
		// 	{
		// 		addTeamMember:           "dbgfnfgng",
		// 		teamMembersRole:         "fdgfhghgf",
		// 		teamMemebersDescription: "here some de",
		// 	},
		// },
		// RoadMap: []RoadMap{
		// 	{
		// 		options:     "444",
		// 		percentages: 45,
		// 	},
		// 	{
		// 		options:     "222",
		// 		percentages: 55,
		// 	},
		// },
	})

	// }

	// 	TeamMembers: &TeamMembers{
	// 		addTeamMember:           "me1",
	// 		teamMembersRole:         "me2",
	// 		teamMemebersDescription: "m3",
	// 	},
	// 	RoadMap: &RoadMap{
	// 		options:     "op1",
	// 		percentages: 45,
	// 	},
	// })

	//routing
	r.HandleFunc("/", servHome).Methods("GET")
	r.HandleFunc("/object", getOject).Methods("GET")
	r.HandleFunc("/objetss", createObjects).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))

}

//MY  CONTROLLER  SERVE HOME ROUTE

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>MY SECOND TASK</h1>"))
}

func getOject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all object")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}

// func getOneObject(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get one  object")
// 	w.Header().Set("Content-Type", "application/json")

// }

func createObjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one object")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Plese insert some data ")
	}

	var object Objects
	_ = json.NewDecoder(r.Body).Decode(&object)

	objects = append(objects, object)
	json.NewEncoder(w).Encode(object)

}
