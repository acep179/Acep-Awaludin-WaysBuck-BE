package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(ID int) (models.Profile, error)
	CreateProfile(models.Profile) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}

func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}
