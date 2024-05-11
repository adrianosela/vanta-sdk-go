package vanta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/adrianosela/vanta-sdk-go/utils/tokenstore"
)

const (
	vantaAPIBaseURL = "https://api.vanta.com"

	ScopeAllRead  = "vanta-api.all:read"
	ScopeAllWrite = "vanta-api.all:write"
)

type vanta struct {
	httpClient *http.Client

	baseURL string

	clientID     string
	clientSecret string
	clientScopes []string

	tokenStore *tokenstore.TokenStore
}

type Option func(*vanta)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(v *vanta) { v.httpClient = httpClient }
}

func WithScopes(scopes ...string) Option {
	return func(v *vanta) { v.clientScopes = scopes }
}

func WithBaseURL(url string) Option {
	return func(v *vanta) { v.baseURL = url }
}

func WithOAuthCredentials(clientID, clientSecret string) Option {
	return func(v *vanta) { v.clientID, v.clientSecret = clientID, clientSecret }
}

func WithToken(token string) Option {
	return func(v *vanta) { ts := new(tokenstore.TokenStore); ts.SetToken("Bearer", token) }
}

func New(ctx context.Context, opts ...Option) (Vanta, error) {
	v := &vanta{
		httpClient:   http.DefaultClient,
		baseURL:      vantaAPIBaseURL,
		clientScopes: []string{ScopeAllRead, ScopeAllWrite}, // all read/write by default
	}
	for _, opt := range opts {
		opt(v)
	}
	if v.tokenStore == nil {
		v.tokenStore = new(tokenstore.TokenStore)
		if err := v.refreshToken(ctx); err != nil {
			return nil, fmt.Errorf("failed to acquire auth token with oauth credentials: %v", err)
		}
	}
	return v, nil
}

func (v *vanta) refreshToken(ctx context.Context) error {
	if v.clientID == "" {
		return errors.New("empty oauth client id")
	}
	if v.clientSecret == "" {
		return errors.New("empty oauth client secret")
	}

	bodyBytes, err := json.Marshal(&GetOauthTokenInput{
		ClientID:     v.clientID,
		ClientSecret: v.clientSecret,
		Scope:        strings.Join(v.clientScopes, " "),
		GrantType:    "client_credentials",
	})
	if err != nil {
		return fmt.Errorf("failed to JSON-encode token request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oauth/token", v.baseURL), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to build http request: %v", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute http request: %v", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read http response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 http response status code (%d), body: %s", resp.StatusCode, string(respBodyBytes))
	}

	var oauthTokenOutput *GetOauthTokenOutput
	if err = json.Unmarshal(respBodyBytes, &oauthTokenOutput); err != nil {
		return fmt.Errorf("failed to JSON-decode token response body: %v", err)
	}

	v.tokenStore.SetToken(oauthTokenOutput.TokenType, oauthTokenOutput.AccessToken)

	return nil
}
