package repository

import "github.com/Pugpaprika21/app/dto"

type IPositionRepository interface {
	Save(body *dto.PositionRequestBody) error
	Get() []*dto.PositionFetchRecord
	GetByID(id string) *dto.PositionFetchRecord
	UpdateByID(id string, body *dto.PositionRequestBody) error
	DeleteByID(id string) error
	CheckPositionNameIsAlreadyHaveIt(body *dto.PositionRequestBody) int64
}
