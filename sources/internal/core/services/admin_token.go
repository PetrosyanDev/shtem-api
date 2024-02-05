package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type adminTokenService struct {
	adminTokenRepository repositories.AdminTokenRepository
}

func (q *adminTokenService) GenerateToken() (*domain.AdminToken, domain.Error) {
	return q.adminTokenRepository.GenerateToken()
}
func (q *adminTokenService) GetToken(token string) (*domain.AdminToken, domain.Error) {
	return q.adminTokenRepository.GetToken(token)
}
func (q *adminTokenService) UpdateToken(t *domain.AdminToken) (*domain.AdminToken, domain.Error) {
	return q.adminTokenRepository.UpdateToken(t)
}

func NewAdminTokenService(adminTokenRepository repositories.AdminTokenRepository) *adminTokenService {
	return &adminTokenService{adminTokenRepository}
}
