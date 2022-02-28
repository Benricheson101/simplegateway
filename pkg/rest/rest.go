package rest

import "net/http"

type RestClient struct {
	token  string
	client *http.Client
}

func New(token string) *RestClient {
	return &RestClient{
		token:  token,
		client: http.DefaultClient,
	}
}

func (rc *RestClient) authed(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bot "+rc.token)
	return rc.client.Do(req)
}
