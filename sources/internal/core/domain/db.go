package domain

type Question struct {
	Q_id     int64
	Bajin    int
	Mas      int
	Q_number int
	Text     string
	Options  []string
	Answers  []int
	ShtemId  int64
}

type Shtemaran struct {
	Id          int64
	Name        string
	Description string
	Author      string
	LinkName    string
	Image       string
	PDF         string
	Keywords    []string
	Category    int64
	HasQuiz     bool
	HasPDF      bool
}

type ShtemBajin struct {
	Id      int64
	ShtemId int64
	Name    string
	Number  int
	IsReady bool
}

type Category struct {
	C_id        int64
	Name        string
	Description string
	LinkName    string
}

type Categories map[Category][]*Shtemaran
type CategoriesTpl map[Category][]Shtemaran
