package oauth

import (
	"fmt"
	"net/http"
)

const (
	// API Headers
	AccessHeader = "Authorization"
)

type Credentials struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
}

type BearerToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scopes       string `json:"scopes"`
	TokenType    string `json:"token_type"`
}

func (b *BearerToken) SignRequest(r *http.Request) {
	if b.AccessToken != "" {
		r.Header.Add(AccessHeader, fmt.Sprintf("Bearer %v", b.AccessToken))
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
}
