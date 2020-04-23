package twitter

type Twitter struct {
	URL       string
	Username  string
	Followers int
}


//Feed returns for Twitter post
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying people's work",
	}
}

//Fame returns the number of connections made
func (t *Twitter) Fame() int {
	return t.Followers
}