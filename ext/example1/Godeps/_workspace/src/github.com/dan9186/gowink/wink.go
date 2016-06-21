package gowink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dan9186/gowink/oauth"
)

const (
	// API URI
	BaseURI = "https://api.wink.com/"
)

type Wink struct {
	baseUri string
	client  http.Client
	creds   *oauth.BearerToken
}

func New(baseUri string) *Wink {
	return &Wink{baseUri, http.Client{}, &oauth.BearerToken{}}
}

func (w *Wink) DoRequest(verb string, uri *url.URL, payload []byte) (*http.Response, error) {
	verb = strings.ToUpper(verb)
	uri.RawQuery = uri.Query().Encode()

	req, err := http.NewRequest(verb, uri.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	w.creds.SignRequest(req)

	return w.client.Do(req)
}

func (w *Wink) GetUri(endpoint string, params url.Values) (*url.URL, error) {
	uri, err := url.Parse(w.baseUri)
	if err != nil {
		return nil, err
	}

	uri.Path = endpoint
	uri.RawQuery = params.Encode()

	return uri, nil
}

func (w *Wink) SignIn(clientId, clientSecret, username, password string) error {
	uri, err := w.GetUri("/oauth2/token", nil)
	if err != nil {
		return err
	}

	oauthCreds := oauth.Credentials{clientId, clientSecret, username, password, "password"}
	b, err := json.Marshal(oauthCreds)
	if err != nil {
		return err
	}

	resp, err := w.DoRequest("POST", uri, b)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad response from Wink: (%v) %v", resp.StatusCode, string(bodyBytes))
	}

	var bt *oauth.BearerToken
	err = json.Unmarshal(bodyBytes, &bt)
	if err != nil {
		return err
	}

	w.creds = bt

	return nil
}
