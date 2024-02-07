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

func (r *UpdateShtemRequest) ToDomain(p *domain.Shtemaran) domain.Error {
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

// DELETE
