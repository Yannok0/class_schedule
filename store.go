package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type LessonStore struct {
	db *pgxpool.Pool
}

func NewLessonStore(db *pgxpool.Pool) *LessonStore {
	return &LessonStore{db: db}
}

func (s *LessonStore) List(ctx context.Context) ([]Lesson, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, group_id, teacher_id, subject_id, starts_at, ends_at, room, status
		 FROM lessons
		 ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lessons := []Lesson{}

	for rows.Next() {
		var lesson Lesson
		if err := rows.Scan(
			&lesson.ID,
			&lesson.GroupID,
			&lesson.TeacherID,
			&lesson.SubjectID,
			&lesson.StartsAt,
			&lesson.EndsAt,
			&lesson.Room,
			&lesson.Status,
		); err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	return lessons, rows.Err()
}

func (s *LessonStore) Get(ctx context.Context, id int64) (Lesson, error) {
	var lesson Lesson

	err := s.db.QueryRow(ctx,
		`SELECT id, group_id, teacher_id, subject_id, starts_at, ends_at, room, status
		 FROM lessons
		 WHERE id = $1`,
		id,
	).Scan(
		&lesson.ID,
		&lesson.GroupID,
		&lesson.TeacherID,
		&lesson.SubjectID,
		&lesson.StartsAt,
		&lesson.EndsAt,
		&lesson.Room,
		&lesson.Status,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return Lesson{}, ErrNotFound
	}

	if err != nil {
		return Lesson{}, err
	}

	return lesson, nil
}

func (s *LessonStore) Create(
	ctx context.Context,
	groupID int,
	teacherID int,
	subjectID int,
	startsAt string,
	endsAt string,
	room string,
) (Lesson, error) {
	var lesson Lesson

	err := s.db.QueryRow(ctx,
		`INSERT INTO lessons
			(group_id, teacher_id, subject_id, starts_at, ends_at, room)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, group_id, teacher_id, subject_id, starts_at, ends_at, room, status`,
		groupID,
		teacherID,
		subjectID,
		startsAt,
		endsAt,
		room,
	).Scan(
		&lesson.ID,
		&lesson.GroupID,
		&lesson.TeacherID,
		&lesson.SubjectID,
		&lesson.StartsAt,
		&lesson.EndsAt,
		&lesson.Room,
		&lesson.Status,
	)

	if err != nil {
		return Lesson{}, err
	}

	return lesson, nil
}

func (s *LessonStore) SetStatus(ctx context.Context, id int64, status string) error {
	tag, err := s.db.Exec(ctx,
		`UPDATE lessons SET status = $1 WHERE id = $2`,
		status,
		id,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *LessonStore) Delete(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx,
		`DELETE FROM lessons WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}