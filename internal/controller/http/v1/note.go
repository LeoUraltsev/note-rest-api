package v1

import (
	"NoteRestApi/internal/service"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type noteRouter struct {
	noteService service.Note
}

func newNoteRouter(router chi.Router, noteService service.Note) {
	r := &noteRouter{noteService: noteService}

	router.Get("/{noteID}", r.getNote)
}

func (r *noteRouter) getNote(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	note, err := r.noteService.GetNoteByID(context.Background(), 123)
	if err != nil {
		log.Println(err)
	}

	resp, err := json.Marshal(note)
	if err != nil {
		return
	}

	writer.Write(resp)
}
