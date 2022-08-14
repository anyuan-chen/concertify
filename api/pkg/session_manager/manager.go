package session_manager

import (
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Session_Manager struct {
	sessions map[string]*oauth2.Token
}

func CreateManager() *Session_Manager {
	return &Session_Manager{
		sessions: make(map[string]*oauth2.Token),
	}
}

func (m *Session_Manager) GetSession(code string) (*oauth2.Token, error) {
	return m.sessions[code], nil
}

func (m *Session_Manager) SetSession(session *oauth2.Token) (string, error) {
	var uid string
	for {
		uid = uuid.New().String()
		session, err := m.GetSession(uid)
		if err != nil {
			return "", err
		} else if session == nil {
			m.sessions[uid] = session
			break;
		}
	}
	return uid, nil
}
