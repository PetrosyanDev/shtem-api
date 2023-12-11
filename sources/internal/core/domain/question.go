package domain

type Question struct {
	ID        int
	ShtemName string
	Bajin     int
	Mas       int
	Number    int
	Text      string
	Options   []string
	Answers   []int
}
