package classdiagram

func GetNotesText(notes []*Note) string {
	s := ""
	for _, n := range notes {
		s += n.Text
	}
	return s
}
