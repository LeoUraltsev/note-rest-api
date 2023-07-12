package v1

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

type noteRouter struct {
	noteService service.Note
}

func newNoteRouter(router chi.Router, noteService service.Note) {
	r := &noteRouter{noteService: noteService}

	router.Get("/notes/{noteID}", r.note)
	router.Get("/notes/", r.allNotes)
	router.Post("/notes/", r.createNote)
	router.Put("/notes/{noteID}", r.updateNote)
	router.Delete("/notes/{noteID}", r.deleteNote)

}

type errAnswer struct {
	ErrCode string
	ErrMsg  string
}

type noteInput struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,"`
}

// Метод для получения заметки по её ID.
func (r *noteRouter) note(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	noteID, _ := strconv.Atoi(chi.URLParam(request, "noteID"))

	note, err := r.noteService.GetNoteByID(context.Background(), noteID)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)

		msg, _ := json.Marshal(errAnswer{"404", fmt.Sprintf("Note by id = %d not found", noteID)})

		if _, err := writer.Write(msg); err != nil {
			return
		}
		return
	}

	resp, err := json.Marshal(note)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(resp); err != nil {
		return
	}
}

func (r *noteRouter) allNotes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	notes, err := r.noteService.GetAllNote(context.Background())
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	marshal, err := json.Marshal(notes)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(marshal); err != nil {
		return
	}
}

func (r *noteRouter) createNote(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	body := request.Body
	defer body.Close()

	ni := noteInput{}

	if err := json.NewDecoder(body).Decode(&ni); err != nil {
		return
	}

	note := entity.Note{
		Title:       ni.Title,
		Description: ni.Description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	id, err := r.noteService.CreateNote(context.Background(), note)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("Note created. ID = %d", id)))
}

func (r *noteRouter) updateNote(writer http.ResponseWriter, request *http.Request) {

	note := entity.Note{
		ID:          100,
		Title:       "Update Title",
		Description: "Update Description",
		ModifiedAt:  time.Now(),
	}

	err := r.noteService.UpdateNote(context.Background(), note)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Success update"))

}

func (r *noteRouter) deleteNote(writer http.ResponseWriter, request *http.Request) {
	var id int

	id, err := strconv.Atoi(chi.URLParam(request, "noteID"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = r.noteService.DeleteNote(context.Background(), id)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf("Note by id=%d success delete", id)))

}
