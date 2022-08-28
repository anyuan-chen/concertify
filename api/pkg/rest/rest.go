package rest

import (
	"github.com/anyuan-chen/concertify/api/pkg/concertify"
	"github.com/anyuan-chen/concertify/api/pkg/session_manager"
)

type ConcertifyAPI struct {
	Session_Manager *session_manager.Session_Manager
	ConcertifyCore  *concertify.ConcertifyCore
}

func NewConcertifyAPI(session_manager *session_manager.Session_Manager, concertifyCore *concertify.ConcertifyCore) ConcertifyAPI {
	return ConcertifyAPI{
		Session_Manager: session_manager,
		ConcertifyCore:  concertifyCore,
	}
}
