package domain

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
