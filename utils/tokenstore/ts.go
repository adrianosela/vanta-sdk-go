package tokenstore

import "sync"

type TokenStore struct {
	sync.RWMutex // inherit read/write lock behavior
	tokenType    string
	token        string
}

func (ts *TokenStore) SetToken(tokenType string, token string) {
	ts.Lock()
	defer ts.Unlock()
	ts.tokenType = tokenType
	ts.token = token

}

func (ts *TokenStore) GetToken() (string, string) {
	ts.RLock()
	defer ts.RUnlock()
	return ts.tokenType, ts.token
}
