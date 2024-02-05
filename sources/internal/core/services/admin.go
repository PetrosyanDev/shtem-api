package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type adminService struct {
	adminRepository repositories.AdminRepository
}

func (q *adminService) Create(username, password string) (*domain.Admin, domain.Error) {
	return q.adminRepository.Create(username, password)
}

func (q *adminService) GetByToken(token string) (*domain.Admin, domain.Error) {
	return q.adminRepository.GetByToken(token)
}

func (q *adminService) Update(adm *domain.Admin) domain.Error {
	return q.adminRepository.Update(adm)
}

func (q *adminService) Delete(id int64) domain.Error {
	return q.adminRepository.Delete(id)
}
func (q *adminService) GetByUsername(username string) (*domain.Admin, domain.Error) {
	return q.adminRepository.GetByUsername(username)
}
func (q *adminService) GetAdmins() (*[]*domain.Admin, domain.Error) {
	return q.adminRepository.GetAdmins()
}

func NewAdminService(adminRepository repositories.AdminRepository) *adminService {
	return &adminService{adminRepository}
}
