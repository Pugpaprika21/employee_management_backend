package repository

import "github.com/Pugpaprika21/app/dto"

type IDepartmentRepository interface {
	Save(body *dto.DepartmentRequestBody) error
	Get() []*dto.DepartmentFetchRecord
	GetByID(id string) *dto.DepartmentFetchRecord
	UpdateByID(id string, body *dto.DepartmentRequestBody) error
	DeleteByID(id string) error
	CheckDepartmentNameIsAlreadyHaveIt(body *dto.DepartmentRequestBody) int64
}
