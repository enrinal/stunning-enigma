package api

import (
	"fmt"
	"net/http"
	"rg-comp-api/repository"
)

type API struct {
	emplRepo repository.EmployeeRepository
	mux      *http.ServeMux
}

func NewApi(emplRepo repository.EmployeeRepository) API {
	mux := http.NewServeMux()
	api := API{
		emplRepo,
		mux,
	}

	mux.Handle("/employee/get-all", api.Get(http.HandlerFunc(api.FetchAllEmployee)))
	mux.Handle("/employee/get", api.Get(http.HandlerFunc(api.FetchEmployeeByID)))
	mux.Handle("/employee/add", api.Post(http.HandlerFunc(api.StoreEmployee)))
	mux.Handle("/employee/update", api.Put(http.HandlerFunc(api.UpdateEmployee)))
	mux.Handle("/employee/delete", api.Delete(http.HandlerFunc(api.DeleteEmployee)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8088", api.Handler())
}
