package service

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/internal/repo"
	"context"
	"time"
)

type NoteServices struct {
	noteRepo repo.Note
}

func (s *NoteServices) CreateNote(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *NoteServices) GetNoteByID(ctx context.Context, noteID int) (entity.Note, error) {
	return entity.Note{
		ID:          noteID,
		Title:       "Title 1",
		Description: "Description 2",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}, nil
}

func NewNoteServices(noteRepo repo.Note) *NoteServices {
	return &NoteServices{noteRepo: noteRepo}
}
