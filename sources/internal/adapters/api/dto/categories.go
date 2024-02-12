package dto

import "shtem-api/sources/internal/core/domain"

// CREATE
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	LinkName    string `json:"link-name" binding:"required"`
}

type CreateCategoryResponce struct {
	C_id        int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LinkName    string `json:"link-name"`
}

func (r *CreateCategoryRequest) ToDomain(p *domain.Category) domain.Error {
	p.Name = r.Name
	p.Description = r.Description
	p.LinkName = r.LinkName

	return nil
}

// FIND BY LINK NAME
type FindCategoryRequest struct {
	LinkName string `json:"link-name" binding:"required"`
}

type FindCategoryResponce struct {
	C_id        int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LinkName    string `json:"link-name"`
}

func (r *FindCategoryRequest) ToDomain(p *domain.Category) domain.Error {

	r.LinkName = p.LinkName

	return nil
}

// UPDATE
type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LinkName    string `json:"link-name"`
}

type UpdateCategoryResponce struct {
	C_id        string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LinkName    string `json:"link-name"`
}

func (r *UpdateCategoryRequest) ToDomain(p *domain.Category, current_category *domain.Category) domain.Error {
	p.Name = r.Name
	p.Description = r.Description
	p.LinkName = r.LinkName

	if p.Name == "" {
		p.Name = current_category.Name
	}
	if p.Description == "" {
		p.Description = current_category.Description
	}
	if p.LinkName == "" {
		p.LinkName = current_category.LinkName
	}

	return nil
}

// GLOBAL
// GLOBAL
// GLOBAL

type CategoryResponceData struct {
	C_id        int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LinkName    string `json:"link-name"`
}

type CategoryResponse struct {
	Response[CategoryResponceData]
}

type CategoriesResponse struct {
	Response[[]CategoryResponceData]
}

func (r *CategoryResponse) FromDomain(p *domain.Category) domain.Error {
	r.Data = new(CategoryResponceData)
	r.Data.C_id = p.C_id
	r.Data.Name = p.Name
	r.Data.Description = p.Description
	r.Data.LinkName = p.LinkName

	return nil
}

func (r *CategoriesResponse) SliceFromDomain(p []*domain.Category) {
	// Initialize r.Data as a pointer to a slice
	r.Data = new([]CategoryResponceData)

	// Initialize the underlying slice
	*r.Data = make([]CategoryResponceData, len(p))

	for index, q := range p {
		(*r.Data)[index] = CategoryResponceData{
			C_id:        q.C_id,
			Name:        q.Name,
			Description: q.Description,
			LinkName:    q.LinkName,
		}
	}
}
