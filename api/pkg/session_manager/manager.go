package session_manager

import (
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
	return m.sessions[code]["spotify"], nil
}
func (m *Session_Manager) GetYoutubeSession(code string) (*oauth2.Token, error) {
	return m.sessions[code]["youtube"], nil
}

func (m *Session_Manager) SetSpotifySession(session *oauth2.Token) (string, error) {
	var uid string
	for {
		uid = uuid.New().String()
		session, err := m.GetSpotifySession(uid)
		if err != nil {
			return "", err
		} else if session == nil {
			m.sessions[uid]["spotify"] = session
			break
		}
	}
	return uid, nil
}
func (m *Session_Manager) SetYoutubeSession(code string, session *oauth2.Token) error {
	m.sessions[code]["youtube"] = session
	return nil
}
