package repositories

import (
	"shtem-api/sources/internal/core/domain"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
)

type adminRepository struct {
	db postgresrepository.AdminDB
}

func (p *adminRepository) Create(username, password string) (*domain.Admin, domain.Error) {
	return p.db.Create(username, password)
}

func (p *adminRepository) GetByToken(token string) (*domain.Admin, domain.Error) {
	return p.db.GetByToken(token)
}

func (p *adminRepository) Update(adm *domain.Admin) domain.Error {
	return p.db.Update(adm)
}

func (p *adminRepository) Delete(id int64) domain.Error {
	return p.db.Delete(id)
}

func NewAdminRepository(db postgresrepository.AdminDB) *adminRepository {
	return &adminRepository{db}
}
