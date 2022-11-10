package repository

import (
	"MINIPROJECT/database"
	"MINIPROJECT/models"
)

type PublisherRepositoryImpl struct{}

func (pr *PublisherRepositoryImpl) GetAllPublisher() []models.Publisher {
	var publishers []models.Publisher

	database.DB.Find(&publishers)

	return publishers

}

func (pr *PublisherRepositoryImpl) GetByIdPublisher(id string) models.Publisher {
	var publisher models.Publisher

	database.DB.First(&publisher, "id_publisher =?", id)
	return publisher
}

func (pr *PublisherRepositoryImpl) CreatePublisher(input models.InputPublisher) models.Publisher {
	var newPublisher models.Publisher = models.Publisher{
		Publisher_name: input.Publisher_name,
		Publisher_desc: input.Publisher_desc,
	}

	var createPublisher models.Publisher = models.Publisher{}

	result := database.DB.Create(&newPublisher)

	result.Last(&createPublisher)

	return createPublisher

}

func (pr *PublisherRepositoryImpl) UpdatePublisher(id string, input models.InputPublisher) models.Publisher {

	var publisher models.Publisher = pr.GetByIdPublisher(id)

	publisher.Publisher_name = input.Publisher_name
	publisher.Publisher_desc = input.Publisher_desc

	database.DB.Save(&publisher)

	return publisher
}

func (pr *PublisherRepositoryImpl) DeletePublisher(id string) bool {
	var publisher models.Publisher = pr.GetByIdPublisher(id)

	result := database.DB.Delete(&publisher)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}
