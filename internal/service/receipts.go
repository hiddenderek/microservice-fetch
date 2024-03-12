package service

import (
	"fetch-service/internal/database"
	"fetch-service/internal/models"

	"github.com/google/uuid"
)

func PostReceipt(item models.Receipt) models.IdResponse {
	newId := uuid.New().String()
	database.PostData[models.Receipt](newId, item)
	idResponse := models.IdResponse{
		Id: newId,
	}
	return idResponse
}
