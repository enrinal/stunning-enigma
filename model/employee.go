package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type EmployeeFetch struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Salary         int    `json:"salary"`
	DepartmentName string `json:"department_name"`
}

type EmployeeAdd struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Salary       int    `json:"salary"`
	DepartmentID int    `json:"department_id"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
