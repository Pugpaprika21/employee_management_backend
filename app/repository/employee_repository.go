package repository

import (
	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (emp *EmployeeRepository) Save(body *dto.EmployeeRequestBody) error {
	data := &models.Employee{
		Name:         body.Name,
		Surname:      body.Surname,
		PositionID:   body.PositionID,
		DepartmentID: body.DepartmentID,
		Salary:       body.Salary,
		Address:      body.Address,
		PhoneNumber:  body.PhoneNumber,
		Email:        body.Email,
		Gender:       body.Gender,
	}

	tx := emp.db.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (emp *EmployeeRepository) Get() []*dto.EmployeeFetchRecord {
	var records []*dto.EmployeeFetchRecord
	emp.db.Model(&models.Employee{}).Where("deleted_at IS NULL").Order("id DESC").Find(&records)
	return records
}

func (emp *EmployeeRepository) GetByID(id string) *dto.EmployeeFetchRecord {
	var record *dto.EmployeeFetchRecord
	emp.db.Model(&models.Employee{}).Where("deleted_at IS NULL AND id = ?", id).Order("id DESC").Find(&record)
	return record
}

func (emp *EmployeeRepository) UpdateByID(id string, body *dto.EmployeeRequestBody) error {
	data := &models.Employee{
		Name:         body.Name,
		Surname:      body.Surname,
		PositionID:   body.PositionID,
		DepartmentID: body.DepartmentID,
		Salary:       body.Salary,
		Address:      body.Address,
		PhoneNumber:  body.PhoneNumber,
		Email:        body.Email,
		Gender:       body.Gender,
	}

	if err := emp.db.Model(&models.Employee{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (emp *EmployeeRepository) DeleteByID(id string) error {
	dept := &models.Employee{}
	if err := emp.db.Where("id = ?", id).Delete(dept).Error; err != nil {
		return err
	}
	return nil
}

func (emp *EmployeeRepository) GetEmployeePositonByID(posID uint) *dto.PositionFetchRecord {
	var record *dto.PositionFetchRecord
	emp.db.Model(&models.PositionMaster{}).Where("deleted_at IS NULL AND id = ?", posID).First(&record)
	return record
}

func (emp *EmployeeRepository) GetEmployeeDepartmentByID(dmpID uint) *dto.DepartmentFetchRecord {
	var record *dto.DepartmentFetchRecord
	emp.db.Model(&models.DepartmentMaster{}).Where("deleted_at IS NULL AND id = ?", dmpID).First(&record)
	return record
}
