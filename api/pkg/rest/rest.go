package rest

import "github.com/anyuan-chen/concertify/api/pkg/session_manager"

type ConcertifyAPI struct {
	Session_Manager *session_manager.Session_Manager
}

func NewConcertifyAPI(session_manager *session_manager.Session_Manager) ConcertifyAPI {
	return ConcertifyAPI{
		Session_Manager: session_manager,
	}
}
