package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NicolasDeveloper/go-rest-api/dtos"
	"github.com/NicolasDeveloper/go-rest-api/repositories/leads"
)

/* Post: insert lead in database */
func Post(w http.ResponseWriter, r *http.Request) {
	newLead, err := getLeadFromBody(w, r)

	if err != nil {
		fmt.Print("formato de lead inválido")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Print(newLead)
	leads.Save(*newLead)
	w.WriteHeader(http.StatusOK)
}

/* Get: insert lead in database */
func Get(w http.ResponseWriter, r *http.Request) {
}

/* GetById: insert lead in database */
func GetById(w http.ResponseWriter, r *http.Request) {
}

/* Put: insert lead in database */
func Put(w http.ResponseWriter, r *http.Request) {
}

func getLeadFromBody(w http.ResponseWriter, r *http.Request) (*dtos.Lead, error) {
	var newLead dtos.Lead
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "não foi possível resgatar o lead")
	}

	err = json.Unmarshal(reqBody, &newLead)

	if err != nil {
		fmt.Fprintf(w, "não foi possível resgatar o lead")
		return nil, errors.New(" ao resgatar lead")
	}

	return &newLead, nil
}
