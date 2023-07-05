package pgdb

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/pkg/postgres"
	"context"
)

type NoteRepo struct {
	*postgres.Postgres
}

func (r *NoteRepo) CreateNote(ctx context.Context, note entity.Note) error {
	//TODO implement me
	panic("implement me")
}

func (r *NoteRepo) GetNoteById(ctx context.Context, id int) (entity.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (r *NoteRepo) GetAllNote(ctx context.Context) ([]entity.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (r *NoteRepo) UpdateNote(ctx context.Context, note entity.Note) error {
	//TODO implement me
	panic("implement me")
}

func (r *NoteRepo) DeleteNote(ctx context.Context, note entity.Note) error {
	//TODO implement me
	panic("implement me")
}

func NewNoteRepo(pg *postgres.Postgres) *NoteRepo {
	return &NoteRepo{pg}
}
