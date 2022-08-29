package session_manager

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Session_Manager struct {
	sessions map[string]map[string]*oauth2.Token
}

func CreateManager() *Session_Manager {
	return &Session_Manager{
		sessions: make(map[string]map[string]*oauth2.Token),
	}
}

func (m *Session_Manager) GetSpotifySession(code string) (*oauth2.Token, error) {
	if m.sessions[code]["spotify"] == nil {
		return nil, errors.New("not available")
	}
	return m.sessions[code]["spotify"], nil
}
func (m *Session_Manager) GetYoutubeSession(code string) (*oauth2.Token, error) {
	if m.sessions[code]["youtube"] == nil {
		return nil, errors.New("not available")
	}
	return m.sessions[code]["youtube"], nil
}

func (m *Session_Manager) CreateSession() (string, error) {
	var uid string
	for {
		uid = uuid.New().String()
		_, err := m.GetSpotifySession(uid)
		if err != nil {
			m.sessions[uid] = make(map[string]*oauth2.Token)
			return uid, nil
		}
	}
}

func (m *Session_Manager) SetSpotifySession(code string, session *oauth2.Token) error {
	m.sessions[code]["spotify"] = session
	return nil
}
func (m *Session_Manager) SetYoutubeSession(code string, session *oauth2.Token) error {
	m.sessions[code]["youtube"] = session
	return nil
}
