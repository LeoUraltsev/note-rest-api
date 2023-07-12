package pgdb

import (
	"NoteRestApi/internal/entity"
	"NoteRestApi/pkg/postgres"
	"context"
	"fmt"
	"golang.org/x/exp/slog"
)

type NoteRepo struct {
	*postgres.Postgres
	*slog.Logger
}

func (r *NoteRepo) CreateNote(ctx context.Context, note entity.Note) (int, error) {

	r.Logger.Info("creating note")
	p := r.Pool

	q := `INSERT INTO note(title, description, created_at, modified_at) values ($1,$2,$3,$4) RETURNING id`

	err := p.QueryRow(ctx, q, note.Title, note.Description, note.CreatedAt, note.ModifiedAt).Scan(&note.ID)
	if err != nil {
		return 0, err
	}

	return note.ID, nil
}

func (r *NoteRepo) GetNoteById(ctx context.Context, id int) (entity.Note, error) {
	var n entity.Note

	p := r.Pool

	q := `SELECT id, title, description, created_at, modified_at FROM note WHERE id = $1`

	err := p.QueryRow(ctx, q, id).Scan(&n.ID, &n.Title, &n.Description, &n.CreatedAt, &n.ModifiedAt)
	if err != nil {
		return n, fmt.Errorf("query was not executed: query: %s, id: %d, error: %v", q, id, err)
	}

	return n, nil
}

func (r *NoteRepo) GetAllNote(ctx context.Context) ([]entity.Note, error) {
	var note = entity.Note{}
	var notes = make([]entity.Note, 10)
	var i int
	p := r.Pool

	q := `SELECT id, title, description, created_at, modified_at FROM note`

	query, err := p.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		if err := query.Scan(&note.ID, &note.Title, &note.Description, &note.CreatedAt, &note.ModifiedAt); err != nil {
			return nil, err
		}

		if i < len(notes) {
			notes[i] = note
		} else {
			notes = append(notes, note)
		}
		i++
	}

	return notes, nil
}

func (r *NoteRepo) UpdateNote(ctx context.Context, note entity.Note) error {
	p := r.Pool
	//goland:noinspection SqlDialectInspection
	q := `UPDATE note 
			SET 
			    title = $1,
			    description = $2,
			    modified_at = $3 
			WHERE note.id = $4`

	_, err := p.Exec(ctx, q, note.Title, note.Description, note.ModifiedAt, note.ID)
	if err != nil {
		return fmt.Errorf("update note error: %v", err)
	}

	return nil
}

func (r *NoteRepo) DeleteNote(ctx context.Context, id int) error {
	p := r.Pool
	q := `DELETE FROM note WHERE id=$1`

	exec, err := p.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %v", exec, err)
	}
	return nil
}

func NewNoteRepo(pg *postgres.Postgres, log *slog.Logger) *NoteRepo {
	return &NoteRepo{pg, log}
}
