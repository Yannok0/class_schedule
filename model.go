package main

import "time"

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Teacher struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type Subject struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Lesson struct {
	ID        int64     `json:"id"`
	GroupID   int64     `json:"group_id"`
	TeacherID int64     `json:"teacher_id"`
	SubjectID int64     `json:"subject_id"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
	Room      string    `json:"room"`
	Status    string    `json:"status"`
}