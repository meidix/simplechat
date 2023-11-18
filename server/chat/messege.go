package chat


type Messege struct {
	sender string
	text string
}

func (m Messege) toString() string {
	return m.sender + " > " + m.text
}