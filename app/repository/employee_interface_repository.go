package repository

import "github.com/Pugpaprika21/app/dto"

type IEmployeeRepository interface {
	Save(body *dto.EmployeeRequestBody) error
	Get() []*dto.EmployeeFetchRecord
	GetByID(id string) *dto.EmployeeFetchRecord
	UpdateByID(id string, body *dto.EmployeeRequestBody) error
	DeleteByID(id string) error
	GetEmployeePositonByID(posID uint) *dto.PositionFetchRecord
	GetEmployeeDepartmentByID(dmpID uint) *dto.DepartmentFetchRecord
}
