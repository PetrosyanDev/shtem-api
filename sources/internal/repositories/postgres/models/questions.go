// Erik Petrosyan ©
package models

import "shtem-api/sources/internal/core/domain"

type Question struct {
	ShtemName string
	ID        int
	Bajin     int
	Mas       int
	Number    int
	Text      string
	Options   []string
	Answers   []int
}

func (s *Question) ToDomain() *domain.Question {
	return &domain.Question{
		ShtemName: s.ShtemName,
	}
}

func (s *Question) FromDomain(d *domain.Question) error {
	s.ShtemName = d.ShtemName
	return nil
}
