package repository

import (
	"database/sql"
	"rg-comp-api/model"
)

type EmployeeRepository interface {
	FetchAll() ([]model.EmployeeFetch, error)
	FetchById(id int) (*model.EmployeeFetch, error)
	Store(empl *model.EmployeeAdd) error
	Update(empl *model.EmployeeAdd) error
	Delete(id int) error
}

// dependency Injection
type employeeRepoImpl struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *employeeRepoImpl {
	return &employeeRepoImpl{db}
}

func (e *employeeRepoImpl) FetchAll() ([]model.EmployeeFetch, error) {
	query := `
		SELECT
			employees.id AS id,
			employees.name AS name,
			employees.address AS address,
			employees.salary AS salary,
			departments.name AS department_name
		FROM
			employees
		INNER JOIN departments
		ON employees.department_id = departments.id
		ORDER BY employees.id ASC;
	`

	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}

	var empls []model.EmployeeFetch
	for rows.Next() {
		var empl model.EmployeeFetch
		err := rows.Scan(&empl.ID, &empl.Name, &empl.Address, &empl.Salary, &empl.DepartmentName)
		if err != nil {
			return nil, err
		}

		empls = append(empls, empl)
	}

	return empls, nil
}

func (e *employeeRepoImpl) FetchById(id int) (*model.EmployeeFetch, error) {
	query := `
		SELECT
			employees.id AS id,
			employees.name AS name,
			employees.address AS address,
			employees.salary AS salary,
			departments.name AS department_name
		FROM
			employees
		INNER JOIN departments
		ON employees.department_id = departments.id
		WHERE employees.id = $1;
	`

	row := e.db.QueryRow(query, id)
	var empl model.EmployeeFetch
	err := row.Scan(&empl.ID, &empl.Name, &empl.Address, &empl.Salary, &empl.DepartmentName)
	if err != nil {
		return nil, err
	}

	return &empl, nil
}

func (e *employeeRepoImpl) Store(empl *model.EmployeeAdd) error {
	query := `INSERT INTO employees (name, address, salary, department_id) VALUES ($1, $2, $3, $4)`

	_, err := e.db.Exec(query, empl.Name, empl.Address, empl.Salary, empl.DepartmentID)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepoImpl) Update(empl *model.EmployeeAdd) error {
	query := `
	UPDATE 
	    employees 
	SET 
	    name=$1, 
	    address=$2, 
	    salary=$3, 
	    department_id=$4
	WHERE id=$5`

	_, err := e.db.Exec(query, empl.Name, empl.Address, empl.Salary, empl.DepartmentID, empl.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepoImpl) Delete(id int) error {
	query := `DELETE FROM employees WHERE id=$1`

	_, err := e.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
