package service

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/internal/repo"
	"context"
)

type NoteServices struct {
	noteRepo repo.Note
}

func (s *NoteServices) UpdateNote(ctx context.Context, note entity.Note) error {
	return s.noteRepo.UpdateNote(ctx, note)
}

func (s *NoteServices) DeleteNote(ctx context.Context, id int) error {
	return s.noteRepo.DeleteNote(ctx, id)
}

func (s *NoteServices) GetAllNote(ctx context.Context) ([]entity.Note, error) {
	return s.noteRepo.GetAllNote(ctx)
}

func (s *NoteServices) CreateNote(ctx context.Context, note entity.Note) (int, error) {
	id, err := s.noteRepo.CreateNote(ctx, note)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *NoteServices) GetNoteByID(ctx context.Context, noteID int) (entity.Note, error) {

	note, err := s.noteRepo.GetNoteById(ctx, noteID)
	if err != nil {
		return entity.Note{}, err
	}

	return note, nil
}

func NewNoteServices(noteRepo repo.Note) *NoteServices {
	return &NoteServices{noteRepo: noteRepo}
}
