package facebook

//Facebook returns 
type Facebook struct {
	URL      string
	Username string
	Friends  int
}

//Feed to new post
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

//Fame returns the number of connections made
func (f *Facebook) Fame() int {
	return f.Friends
}
