package v1

import (
	"NoteRestApi/internal/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter(router chi.Router, services *service.NoteServices) {
	router.Route("/auth", func(r chi.Router) {

	})

	router.Route("/api/v1", func(r chi.Router) {
		newNoteRouter(r, services)
	})
}
