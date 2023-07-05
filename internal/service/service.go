package service

import (
	"NoteRestApi/internal/entity"
	"context"
)

type Note interface {
	CreateNote(ctx context.Context) error
	GetNoteByID(ctx context.Context, noteID int) (entity.Note, error)
}
