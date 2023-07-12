package repo

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/internal/repo/pgdb"
	"NoteRestApi/pkg/postgres"
	"context"
	"golang.org/x/exp/slog"
)

type Note interface {
	CreateNote(ctx context.Context, note entity.Note) (int, error)
	GetNoteById(ctx context.Context, id int) (entity.Note, error)
	GetAllNote(ctx context.Context) ([]entity.Note, error)
	UpdateNote(ctx context.Context, note entity.Note) error
	DeleteNote(ctx context.Context, id int) error
}

type Repositories struct {
	Note
}

func NewRepositories(pg *postgres.Postgres, log *slog.Logger) *Repositories {
	return &Repositories{Note: pgdb.NewNoteRepo(pg, log)}
}
