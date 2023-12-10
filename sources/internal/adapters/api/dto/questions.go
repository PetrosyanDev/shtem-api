package dto

import "shtem-api/sources/internal/core/domain"

type CreateQuestionRequest struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

type UpdateQuestionRequest struct {
	ID        int      `json:"id"`
	ShtemName string   `json:"shtemName"`
	Bajin     int      `json:"bajin"`
	Mas       int      `json:"mas"`
	Number    int      `json:"number"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	Answers   []int    `json:"answers"`
}

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

func (r *CreateQuestionRequest) ToDomain(p *domain.Question) domain.Error {
	if r.ShtemName == "" {
		return domain.ErrBadRequest
	}
	p.ShtemName = r.ShtemName
	p.Bajin = r.Bajin
	p.Mas = r.Mas
	p.Number = r.Number
	p.Text = r.Text
	p.Options = r.Options
	p.Answers = r.Answers
	return nil
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
