package repositories

import (
	"shtem-api/sources/internal/core/domain"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
)

type adminTokenRepository struct {
	db postgresrepository.AdminTokenDB
}

func (p *adminTokenRepository) GenerateToken(id int64) (*domain.AdminToken, domain.Error) {
	return p.db.GenerateToken(id)
}
func (p *adminTokenRepository) GetToken(token string) (*domain.AdminToken, domain.Error) {
	return p.db.GetToken(token)
}
func (p *adminTokenRepository) UpdateToken(t *domain.AdminToken) (*domain.AdminToken, domain.Error) {
	return p.db.UpdateToken(t)
}
func (p *adminTokenRepository) Delete(id int64) domain.Error {
	return p.db.Delete(id)
}

func NewAdminTokenRepository(db postgresrepository.AdminTokenDB) *adminTokenRepository {
	return &adminTokenRepository{db}
}
