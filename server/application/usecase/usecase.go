package usecase

import (
	"server/application/dto"
	"server/application/service"
	"server/domain/repository"
)

type MeetingsUseCase interface {
	StartNew(id string) (*dto.MeetingDto, error)
	GetAvailableCatalogItemsQuantity(ids []string) ([]*dto.MeetingDto, error)
	SeedDatabase() error
}

type meetingsUseCase struct {
	repo    repository.MeetingsRepository
	service *service.MeetingsService
}
