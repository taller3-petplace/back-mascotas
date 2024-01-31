package requester

type Requester struct {
	clientHTTP HttpClient
	baseUrl    string
}

func NewRequester(client HttpClient, url string) *Requester {
	var requester Requester
	requester.clientHTTP = client
	requester.baseUrl = url
	return &requester
}
