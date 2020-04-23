package linkedin

//Linkedin struct
type Linkedin struct {
	URL         string
	Username    string
	Connections int
}


//Feed returns the linkedin post
func (l *Linkedin) Feed() []string {
	return []string{
		"Linkedin feeds",
		"Hey, I just started a new position at Hotels.ng",
	}
}

//Fame returns the number of connections made
func (l *Linkedin) Fame() int {
	return l.Connections
}

