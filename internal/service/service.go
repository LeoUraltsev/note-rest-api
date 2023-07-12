package service

import (
	"NoteRestApi/internal/entity"
	"context"
)

type Note interface {
	CreateNote(ctx context.Context, note entity.Note) (int, error)
	GetNoteByID(ctx context.Context, noteID int) (entity.Note, error)
	GetAllNote(ctx context.Context) ([]entity.Note, error)
	DeleteNote(ctx context.Context, id int) error
	UpdateNote(ctx context.Context, note entity.Note) error
}
