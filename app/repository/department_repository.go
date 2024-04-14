package repository

import (
	"strings"

	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/models"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{
		db: db,
	}
}

func (d *DepartmentRepository) Save(body *dto.DepartmentRequestBody) error {
	data := &models.DepartmentMaster{
		Name:     strings.TrimSpace(body.Name),
		Location: body.Location,
	}

	tx := d.db.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (d *DepartmentRepository) Get() []*dto.DepartmentFetchRecord {
	var records []*dto.DepartmentFetchRecord
	d.db.Model(&models.DepartmentMaster{}).Where("deleted_at IS NULL").Order("id DESC").Find(&records)
	return records
}

func (d *DepartmentRepository) GetByID(id string) *dto.DepartmentFetchRecord {
	var record *dto.DepartmentFetchRecord
	d.db.Model(&models.DepartmentMaster{}).Where("deleted_at IS NULL AND id = ?", id).First(&record)
	return record
}

func (d *DepartmentRepository) UpdateByID(id string, body *dto.DepartmentRequestBody) error {
	data := &models.DepartmentMaster{
		Name:     strings.TrimSpace(body.Name),
		Location: body.Location,
	}

	if err := d.db.Model(&models.DepartmentMaster{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) DeleteByID(id string) error {
	dept := &models.DepartmentMaster{}
	if err := d.db.Where("id = ?", id).Delete(dept).Error; err != nil {
		return err
	}
	return nil
}

func (d *DepartmentRepository) CheckDepartmentNameIsAlreadyHaveIt(body *dto.DepartmentRequestBody) int64 {
	var records int64
	d.db.Model(&models.DepartmentMaster{}).Where("name", strings.TrimSpace(body.Name)).Count(&records)
	return records
}
