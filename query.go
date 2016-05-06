package goscholar

type Query struct {
	keywords string
	author string
	title string
	cluster_id string
	info_id string
	after string
	before string
	num string
	start string
}


func (q *Query) SearchUrl() (url string) {
	return ""
}

func (q *Query) FindUrl() (url string) {
	return ""
}

func (q *Query) CiteUrl() (url string) {
	return ""
}

func (q *Query) CitePopUpQueryUrl() (url string) {
	return ""
}
