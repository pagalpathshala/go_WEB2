package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Record struct {
	Id      string `json :"id"`
	Name    string `json :"name"`
	City    string `json :"city"`
	IsAlive bool   `json :"isalive`
}

var Records []Record

func Get_All(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Here Is The List Of All Records In Our System")
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(Records)

}

func Get_Something(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	opt := mux.Vars(r)
	for _, v := range Records {
		if v.Id == opt["id"] {
			json.NewEncoder(w).Encode(v)
		}
	}
}

func Create_Something(w http.ResponseWriter, r *http.Request) {
	var record Record
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&record)
	Records = append(Records, record)
	fmt.Fprintf(w, "New Record Created Sucessfully.\n")
	json.NewEncoder(w).Encode(record)
	fmt.Fprintf(w, "\nHere Is Your Updated List Down Below\n")
	json.NewEncoder(w).Encode(Records)

}

func Update_Somthing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	opt := mux.Vars(r)
	for i, v := range Records {
		if v.Id == opt["id"] {
			fmt.Fprint(w, "\nYour Entered Id `Key` Is Matched With Our System \n")
			Records = append(Records[:i], Records[i+1:]...)
			fmt.Fprint(w, "\nDeleting Is In Progress !\n")
		}
	}
	fmt.Fprint(w, "\nRecord Deleted Successfully !\n")
	var record Record
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&record)
	Records = append(Records, record)
	fmt.Fprintf(w, "New Record Created Sucessfully.\n")
	json.NewEncoder(w).Encode(record)
	fmt.Fprintf(w, "\nHere Is Your Updated List Down Below\n")
	json.NewEncoder(w).Encode(Records)

}

func Delete_Something(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	opt := mux.Vars(r)
	for i, v := range Records {
		if v.Id == opt["id"] {
			fmt.Fprint(w, "\nYour Entered Id `Key` Is Matched With Our System \n")
			Records = append(Records[:i], Records[i+1:]...)
			fmt.Fprint(w, "\nDeleting Is In Progress !\n")
		}
	}
	fmt.Fprint(w, "\nRecord Deleted Successfully !\n")
	json.NewEncoder(w).Encode(Records)
}

func main() {
	Records = []Record{
		Record{
			Id:      "1",
			Name:    "Rampati",
			City:    "Rohtak",
			IsAlive: true,
		}, Record{
			Id:      "2",
			Name:    "Khushboo",
			City:    "Rohtak",
			IsAlive: true,
		}, Record{
			Id:      "3",
			Name:    "Vikash",
			City:    "Rohtak",
			IsAlive: true,
		}, Record{
			Id:      "4",
			Name:    "Ritika",
			City:    "Rohtak",
			IsAlive: true,
		}, Record{
			Id:      "5",
			Name:    "Niyati",
			City:    "Rohtak",
			IsAlive: true,
		}, Record{
			Id:      "6",
			Name:    "Bijender",
			City:    "Rohtak",
			IsAlive: false,
		}, Record{
			Id:      "7",
			Name:    "Ram_Niwas",
			City:    "Rohtak",
			IsAlive: false,
		}, Record{
			Id:      "8",
			Name:    "Bharpai_Devi",
			City:    "Rohtak",
			IsAlive: false,
		},
	}
	r := mux.NewRouter()
	r.HandleFunc("/records", Get_All).Methods("GET")
	r.HandleFunc("/records/{id}", Get_Something).Methods("GET")
	r.HandleFunc("/records", Create_Something).Methods("POST")
	r.HandleFunc("/records/{id}", Update_Somthing).Methods("PUT")
	r.HandleFunc("/records/{id}", Delete_Something).Methods("DELETE")
	http.ListenAndServe(":9090", r)
}
