package dto

import (
	"shtem-api/sources/internal/core/domain"
	"time"
)

// CREATE
type AdminCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminCreateResponce struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     struct {
		ID        time.Time `json:"id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Expiry    time.Time `json:"expiry"`
	} `json:"token"`
}

func (r *AdminCreateRequest) ToDomain(p *domain.Admin) domain.Error {
	p.Username = r.Username
	p.Password = r.Password
	return nil
}

// UPDATE
type AdminUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminUpdateResponce struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

func (r *AdminUpdateRequest) ToDomain(p *domain.Admin) domain.Error {
	p.Username = r.Username
	p.Password = r.Password
	return nil
}

// DELETE
type AdminDeleteRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type AdminDeleteResponce struct {
	ID int64 `json:"id"`
}

func (r *AdminDeleteRequest) ToDomain(p *domain.Admin) domain.Error {
	p.ID = r.ID
	return nil
}

// GLOBAL
// GLOBAL
// GLOBAL
type AdminResponseData struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     struct {
		ID        int64     `json:"id"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Expiry    time.Time `json:"expiry"`
	} `json:"token"`
}

type AdminResponse struct {
	Response[AdminResponseData]
}

func (r *AdminResponse) FromDomain(p *domain.Admin) {
	r.Data = new(AdminResponseData)
	r.Data.ID = p.ID
	r.Data.CreatedAt = p.CreatedAt
	r.Data.UpdatedAt = p.UpdatedAt
	r.Data.Username = p.Username
	r.Data.Password = p.Password
	r.Data.Token.ID = p.Token.Id
	r.Data.Token.Token = p.Token.Token
	r.Data.Token.CreatedAt = p.Token.CreatedAt
	r.Data.Token.UpdatedAt = p.Token.UpdatedAt
	r.Data.Token.Expiry = p.Token.Expiry
}
