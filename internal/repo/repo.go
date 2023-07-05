package repo

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/internal/repo/pgdb"
	"NoteRestApi/pkg/postgres"
	"context"
)

type Note interface {
	CreateNote(ctx context.Context, note entity.Note) error
	GetNoteById(ctx context.Context, id int) (entity.Note, error)
	GetAllNote(ctx context.Context) ([]entity.Note, error)
	UpdateNote(ctx context.Context, note entity.Note) error
	DeleteNote(ctx context.Context, note entity.Note) error
}

type Repositories struct {
	Note
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{Note: pgdb.NewNoteRepo(pg)}
}
