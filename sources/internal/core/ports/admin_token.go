package ports

import "shtem-api/sources/internal/core/domain"

type AdminTokenService interface {
	GenerateToken(id int64) (*domain.AdminToken, domain.Error)
	GetToken(token string) (*domain.AdminToken, domain.Error)
	UpdateToken(t *domain.AdminToken) (*domain.AdminToken, domain.Error)
}
