package dto

import "shtem-api/sources/internal/core/domain"

// CREATE
type CreateQuestionRequest struct {
	ShtemName string   `json:"shtemName" binding:"required"`
	Bajin     int      `json:"bajin" binding:"required"`
	Mas       int      `json:"mas" binding:"required"`
	Number    int      `json:"number" binding:"required"`
	Text      string   `json:"text" binding:"required"`
	Options   []string `json:"options" binding:"required"`
	Answers   []int    `json:"answers" binding:"required"`
}

type CreateQuestionResponce struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

func (r *CreateQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.ShtemName = r.ShtemName
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Number = r.Number
	p.Text = r.Text
	p.Options = r.Options
	p.Answers = r.Answers
	return nil
}

// UPDATE
type UpdateQuestionRequest struct {
	ID        int      `json:"id" binding:"required"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

type UpdateQuestionResponce struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

func (r *UpdateQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.ID = r.ID
	p.ShtemName = r.ShtemName
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Number = r.Number
	p.Text = r.Text
	p.Options = r.Options
	p.Answers = r.Answers
	return nil
}

// DELETE
type DeleteQuestionRequest struct {
	ID        int    `json:"id" binding:"required"`
	ShtemName string `json:"shtemName" binding:"required"`
}

type DeleteQuestionResponce struct {
	ID        int    `json:"id"`
	ShtemName string `json:"shtemName"`
}

func (r *DeleteQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.ID = r.ID
	p.ShtemName = r.ShtemName
	return nil
}

// FIND
type FindQuestionRequest struct {
	ShtemName string `json:"shtemName" binding:"required"`
	Bajin     int    `json:"bajin" binding:"required"`
	Mas       int    `json:"mas" binding:"required"`
}

type FindQuestionResponce struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

func (r *FindQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.ShtemName = r.ShtemName
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	return nil
}

// Globals
type QuestionResponseData struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

type QuestionResponse struct {
	Response[QuestionResponseData]
}

func (r *QuestionResponse) FromDomain(p *domain.Question) {
	r.Data = new(QuestionResponseData)
	r.Data.ShtemName = p.ShtemName
	r.Data.Bajin = p.Bajin
	r.Data.Mas = p.Mas
	r.Data.Number = p.Number
	r.Data.Text = p.Text
	r.Data.Options = p.Options
	r.Data.Answers = p.Answers
}
