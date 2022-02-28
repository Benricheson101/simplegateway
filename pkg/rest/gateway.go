package rest

import (
	"encoding/json"
	"net/http"
)

type GatewayBot struct {
	Shards            int    `json:"shards"`
	URL               string `json:"url"`
	SessionStartLimit struct {
		MaxConcurrency int `json:"max_concurrency"`
		Remaining      int `json:"Remaining"`
		ResetAfter     int `json:"reset_after"`
		Total          int `json:"total"`
	} `json:"session_start_limit"`
}

func (r *RestClient) GetGatewayAuthed() *GatewayBot {
	req, _ := http.NewRequest("GET", "https://discord.com/api/v10/gateway/bot", nil)
	res, _ := r.authed(req)

	var body GatewayBot
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		panic(err)
	}

	return &body
}
