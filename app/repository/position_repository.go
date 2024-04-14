package repository

import (
	"strings"

	"github.com/Pugpaprika21/app/dto"
	"github.com/Pugpaprika21/app/models"
	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{
		db: db,
	}
}

func (p *PositionRepository) Save(body *dto.PositionRequestBody) error {
	data := &models.PositionMaster{
		Name:             strings.TrimSpace(body.Name),
		Location:         body.Location,
		Description:      body.Description,
		Salary:           body.Salary,
		Status:           "1",
		Responsibilities: body.Responsibilities,
	}

	tx := p.db.Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (p *PositionRepository) Get() []*dto.PositionFetchRecord {
	var records []*dto.PositionFetchRecord
	p.db.Model(&models.PositionMaster{}).Where("deleted_at IS NULL").Order("id DESC").Find(&records)
	return records
}

func (p *PositionRepository) GetByID(id string) *dto.PositionFetchRecord {
	var record *dto.PositionFetchRecord
	p.db.Model(&models.PositionMaster{}).Where("deleted_at IS NULL AND id = ?", id).First(&record)
	return record
}

func (p *PositionRepository) UpdateByID(id string, body *dto.PositionRequestBody) error {
	data := &models.PositionMaster{
		Name:             strings.TrimSpace(body.Name),
		Location:         body.Location,
		Description:      body.Description,
		Salary:           body.Salary,
		Status:           "1",
		Responsibilities: body.Responsibilities,
	}

	if err := p.db.Model(&models.PositionMaster{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (p *PositionRepository) DeleteByID(id string) error {
	pos := &models.PositionMaster{}
	if err := p.db.Where("id = ?", id).Delete(pos).Error; err != nil {
		return err
	}
	return nil
}

func (d *PositionRepository) CheckPositionNameIsAlreadyHaveIt(body *dto.PositionRequestBody) int64 {
	var records int64
	d.db.Model(&models.PositionMaster{}).Where("name", strings.TrimSpace(body.Name)).Count(&records)
	return records
}
