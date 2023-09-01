package service

import (
	"context"
	"database/sql"
	"errors"
	db "github.com/adykaaa/online-notes/db/sqlc"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

var (
	ErrAlreadyExists     = errors.New("note already exists")
	ErrDBInternal        = errors.New("internal DB error during operation")
	ErrNotFound          = errors.New("requested note is not found")
	ErrUserAlreadyExists = errors.New("note already exists")
	ErrUserNotFound      = errors.New("requested note is not found")
)

type service struct {
	q db.Querier
}

func NewService(q db.Querier) *service {
	return &service{q}
}

func (s *service) RegisterUser(ctx context.Context, args *db.RegisterUserParams) (string, error) {
	userName, err := s.q.RegisterUser(ctx, args)

	switch {
	case err != nil:
		if err.(*pq.Error).Code.Name() == "unique_violation" {
			return "", ErrUserAlreadyExists
		}
		return "", ErrDBInternal
	default:
		return userName, nil
	}
}

func (s *service) GetUser(ctx context.Context, userName string) (db.User, error) {
	user, err := s.q.GetUser(ctx, userName)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return db.User{}, ErrUserNotFound
	case err != nil:
		return db.User{}, ErrDBInternal
	default:
		return user, nil
	}
}

func (s *service) CreateNote(ctx context.Context, title string, userName string, text string) (uuid.UUID, error) {
	retId, err := s.q.CreateNote(ctx, &db.CreateNoteParams{
		ID:        uuid.New(),
		Title:     title,
		Username:  userName,
		Text:      sql.NullString{String: text, Valid: true},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	switch {
	case err.(*pq.Error).Code.Name() == "unique_violation":
		return uuid.Nil, ErrAlreadyExists
	case err != nil:
		return uuid.Nil, ErrDBInternal
	default:
		return retId, nil
	}
}

func (s *service) GetAllNotesFromUser(ctx context.Context, userName string) ([]db.Note, error) {
	notes, err := s.q.GetAllNotesFromUser(ctx, userName)
	if err != nil {
		return nil, ErrDBInternal
	}
	return notes, nil
}

func (s *service) DeleteNote(ctx context.Context, reqId uuid.UUID) (uuid.UUID, error) {
	id, err := s.q.DeleteNote(ctx, reqId)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return uuid.Nil, ErrNotFound
	case err != nil:
		return uuid.Nil, ErrDBInternal
	default:
		return id, nil
	}
}

func (s *service) UpdateNote(ctx context.Context, reqId uuid.UUID, title string, isTextValid bool) (uuid.UUID, error) {
	id, err := s.q.UpdateNote(ctx, &db.UpdateNoteParams{
		ID:        reqId,
		Title:     sql.NullString{String: title, Valid: true},
		Text:      sql.NullString{String: title, Valid: isTextValid},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return uuid.Nil, ErrNotFound
	case err != nil:
		return uuid.Nil, ErrDBInternal
	default:
		return id, nil
	}
}
