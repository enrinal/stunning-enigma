package api

import (
	"encoding/json"
	"net/http"
	"rg-comp-api/model"
	"strconv"
)

func (api *API) FetchAllEmployee(writer http.ResponseWriter, request *http.Request) {
	employee, err := api.emplRepo.FetchAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employee)
}

func (api *API) StoreEmployee(writer http.ResponseWriter, request *http.Request) {
	var employee model.EmployeeAdd
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.emplRepo.Store(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employee)
}

func (api *API) UpdateEmployee(writer http.ResponseWriter, request *http.Request) {
	var employee model.EmployeeAdd
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.emplRepo.Update(&employee)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employee)

}

func (api *API) DeleteEmployee(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := request.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.emplRepo.Delete(idInt)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (api *API) FetchEmployeeByID(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	employee, err := api.emplRepo.FetchById(idInt)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(employee)
}
