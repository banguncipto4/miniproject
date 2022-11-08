package service

import (
	"MINIPROJECT/models"
	"MINIPROJECT/repository"
)

type PublisherService struct {
	PublisherRepository repository.PublisherRepository
}

func NewPublisherService() PublisherService {
	return PublisherService{
		PublisherRepository: &repository.PublisherRepositoryImpl{},
	}
}

func (ps *PublisherService) GetAllPublisher() []models.Publisher {
	return ps.PublisherRepository.GetAllPublisher()
}

func (ps *PublisherService) GetByIdPublisher(id string) models.Publisher {
	return ps.PublisherRepository.GetByIdPublisher(id)
}

func (ps *PublisherService) CreatePublisher(input models.InputPublisher) models.Publisher {
	return ps.PublisherRepository.CreatePublisher(input)
}

func (ps *PublisherService) UpdatePublisher(id string, input models.InputPublisher) models.Publisher {
	return ps.PublisherRepository.UpdatePublisher(id, input)
}

func (ps *PublisherService) DeletePublisher(id string) bool {

	return ps.PublisherRepository.DeletePublisher(id)

}
