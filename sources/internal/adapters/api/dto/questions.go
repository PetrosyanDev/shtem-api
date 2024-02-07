package dto

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/ports"
)

// CREATE
type CreateQuestionRequest struct {
	Bajin   int      `json:"bajin" binding:"required"`
	Mas     int      `json:"mas" binding:"required"`
	Number  int      `json:"number" binding:"required"`
	Text    string   `json:"text" binding:"required"`
	Options []string `json:"options" binding:"required"`
	Answers []int    `json:"answers" binding:"required"`
	ShtemId int      `json:"shtemaran" binding:"required"`
}

type CreateQuestionResponce struct {
	ID      int64    `json:"id"`
	Bajin   int      `json:"bajin"`
	Mas     int      `json:"mas"`
	Number  int      `json:"number"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answers []int    `json:"answers"`
	ShtemId int64    `json:"shtemaran"`
}

func (r *CreateQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Q_number = r.Number
	p.Text = r.Text
	p.Options = r.Options
	p.Answers = r.Answers
	p.ShtemId = int64(r.ShtemId)
	return nil
}

// UPDATE
type UpdateQuestionRequest struct {
	Bajin   int      `json:"bajin"`
	Mas     int      `json:"mas"`
	Number  int      `json:"number"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answers []int    `json:"answers"`
	ShtemId int64    `json:"shtemaran"`
}

type UpdateQuestionResponce struct {
	ID      int      `json:"id"`
	Bajin   int      `json:"bajin"`
	Mas     int      `json:"mas"`
	Number  int      `json:"number"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answers []int    `json:"answers"`
	ShtemId int64    `json:"shtemaran"`
}

func (r *UpdateQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.ShtemId = r.ShtemId
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Q_number = r.Number
	p.Text = r.Text
	p.Options = r.Options
	p.Answers = r.Answers
	return nil
}

// DELETE
type DeleteQuestionRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type DeleteQuestionResponce struct {
	ID int64 `json:"id"`
}

func (r *DeleteQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	p.Q_id = r.ID
	return nil
}

// FIND
type FindQuestionRequest struct {
	Bajin         int    `json:"bajin" binding:"required"`
	Mas           int    `json:"mas" binding:"required"`
	Number        int    `json:"number" binding:"required"`
	ShtemLinkName string `json:"shtemaran" binding:"required"`
}

type FindQuestionResponce struct {
	Bajin   int      `json:"bajin"`
	Mas     int      `json:"mas"`
	Number  int      `json:"number"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answers []int    `json:"answers"`
	ShtemId int64    `json:"shtemaran"`
}

func (r *FindQuestionRequest) ToDomain(p *domain.Question, s ports.ShtemsService) domain.Error {

	shtem, err := s.GetShtemByLinkName(r.ShtemLinkName)
	if err != nil {
		return domain.ErrBadRequest
	}

	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Q_number = r.Number
	p.ShtemId = shtem.Id

	return nil
}

// GLOBAL
// GLOBAL
// GLOBAL
type QuestionResponseData struct {
	ID      int64    `json:"id,omitempty"`
	Bajin   int      `json:"bajin,omitempty"`
	Mas     int      `json:"mas,omitempty"`
	Number  int      `json:"number,omitempty"`
	Text    string   `json:"text,omitempty"`
	Options []string `json:"options,omitempty"`
	Answers []int    `json:"answers,omitempty"`
	ShtemId int64    `json:"shtemaran,omitempty"`
}

type QuestionResponse struct {
	Response[QuestionResponseData]
}

type BajinResponse struct {
	Response[[]QuestionResponseData]
}

func (r *QuestionResponse) FromDomain(p *domain.Question) {
	r.Data = new(QuestionResponseData)
	r.Data.ID = p.Q_id
	r.Data.Bajin = p.Bajin
	r.Data.Mas = p.Mas
	r.Data.Number = p.Q_number
	r.Data.Text = p.Text
	r.Data.Options = p.Options
	r.Data.Answers = p.Answers
	r.Data.ShtemId = p.ShtemId
}

func (r *BajinResponse) SliceFromDomain(p []*domain.Question) {
	// Initialize r.Data as a pointer to a slice
	r.Data = new([]QuestionResponseData)

	// Initialize the underlying slice
	*r.Data = make([]QuestionResponseData, len(p))

	for index, q := range p {
		(*r.Data)[index] = QuestionResponseData{
			ID:      q.Q_id,
			Bajin:   q.Bajin,
			Mas:     q.Mas,
			Number:  q.Q_number,
			Text:    q.Text,
			Options: q.Options,
			Answers: q.Answers,
			ShtemId: q.Q_id,
		}
	}
}

func (r *ShtemsResponce) SliceFromDomain(p []string) {
	// Initialize r.Data as a pointer to a slice
	r.Data = new([]string)

	// Initialize the underlying slice
	*r.Data = make([]string, len(p))

	for index, q := range p {
		(*r.Data)[index] = q
	}
}
