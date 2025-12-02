package handlers

import (
	// "fmt"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Student2 struct{
	Lastname string `json:"lastname"`
}

type RegisterData struct{
	Email string `json:"email"`
	Age int `json:age`
}

func GetHandler(w http.ResponseWriter, r *http.Request){

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(Student{"mayank", "pune"})
		
		fmt.Println(r.URL)
		fmt.Println(r.Method)
	}


func PostHandler(w http.ResponseWriter, r *http.Request){
	var student2 Student2

	err := json.NewDecoder(r.Body).Decode(&student2)

	if err != nil {
		return
	}

	fmt.Fprintf(w, "Received: %v", student2.Lastname)
}

func PathParams(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Fprintf(w, "%v", name)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var registerData RegisterData 

	err := json.NewDecoder(r.Body).Decode(&registerData)

	if err != nil {
		return 
	}

	if registerData.Age < 18{
		w.WriteHeader(400)
	}else{
		fmt.Fprintf(w, "%v", registerData.Email)
	}
}
