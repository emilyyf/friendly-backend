package services

import "gitlab.com/emilyyf/friendly-backend/internal/models"

type AuthService interface {
	models.SignUpInput()
}
