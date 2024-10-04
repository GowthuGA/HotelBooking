package services

import (
    "hotel-booking/models"
    "hotel-booking/repositories"
    "time"
)

type AdminService struct {
    repo *repositories.AdminRepository
}

func (as *AdminService) Authenticate(username, password string) (*models.Admin, *models.ResponseError) {
    return as.repo.Authenticate(username, password)
}
