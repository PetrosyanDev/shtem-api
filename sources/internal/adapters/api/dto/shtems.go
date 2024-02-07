package dto

import (
	"shtem-api/sources/internal/core/domain"
)

// CREATE
type CreateShtemRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name" binding:"required"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category" binding:"required"`
}

type CreateShtemResponce struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category"`
}

func (r *CreateShtemRequest) ToDomain(p *domain.Shtemaran) domain.Error {
	p.Name = r.Name
	p.Description = r.Description
	p.Author = r.Author
	p.LinkName = r.LinkName
	p.Image = r.Image
	p.PDF = r.PDF
	p.Keywords = r.Keywords
	p.Category = r.Category

	return nil
}

// FIND
type FindShtemRequest struct {
	LinkName string `json:"link-name" binding:"required"`
}

type FindShtemResponce struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category"`
}

func (r *FindShtemRequest) ToDomain(p *domain.Shtemaran) domain.Error {

	r.LinkName = p.LinkName

	return nil
}

// UPDATE
type UpdateShtemRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category"`
}

type UpdateShtemResponce struct {
	Id          int64    `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name" binding:"required"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category" binding:"required"`
}

func (r *UpdateShtemRequest) ToDomain(p *domain.Shtemaran, shtemaran *domain.Shtemaran) domain.Error {
	p.Name = r.Name
	p.Description = r.Description
	p.Author = r.Author
	p.LinkName = r.LinkName
	p.Image = r.Image
	p.PDF = r.PDF
	p.Keywords = r.Keywords
	p.Category = r.Category

	if p.Name == "" {
		p.Name = shtemaran.Name
	}
	if p.Description == "" {
		p.Description = shtemaran.Description
	}
	if p.Author == "" {
		p.Author = shtemaran.Author
	}
	if p.LinkName == "" {
		p.LinkName = shtemaran.LinkName
	}
	if p.Image == "" {
		p.Image = shtemaran.Image
	}
	if p.PDF == "" {
		p.PDF = shtemaran.PDF
	}
	if p.Keywords == nil {
		p.Keywords = shtemaran.Keywords
	}
	if p.Category == 0 {
		p.Category = shtemaran.Category
	}

	return nil
}

// UPDATE COVER
type UploadCoverShtemRequest struct {
	Data64 string `json:"data64" binding:"required"`
}

// GLOBAL
// GLOBAL
// GLOBAL

type ShtemResponceData struct {
	Id          int64    `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	LinkName    string   `json:"link-name" binding:"required"`
	Image       string   `json:"image"`
	PDF         string   `json:"pdf"`
	Keywords    []string `json:"keywords"`
	Category    int64    `json:"category" binding:"required"`
}

type ShtemResponse struct {
	Response[ShtemResponceData]
}

type ShtemsResponce struct {
	Response[[]string]
}

func (r *ShtemResponse) FromDomain(p *domain.Shtemaran) domain.Error {
	r.Data = new(ShtemResponceData)
	r.Data.Id = p.Id
	r.Data.Name = p.Name
	r.Data.Description = p.Description
	r.Data.Author = p.Author
	r.Data.LinkName = p.LinkName
	r.Data.Image = p.Image
	r.Data.PDF = p.PDF
	r.Data.Keywords = p.Keywords
	r.Data.Category = p.Category

	return nil
}
