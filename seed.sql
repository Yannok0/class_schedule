INSERT INTO groups (name)
VALUES
    ('ИС-111'),
    ('ИС-222');

INSERT INTO teachers (full_name)
VALUES
    ('Иванов Иван Иванович'),
    ('Петров Петр Петрович');

INSERT INTO subjects (name)
VALUES
    ('Программирование'),
    ('Базы данных');

INSERT INTO lessons (
    group_id,
    teacher_id,
    subject_id,
    starts_at,
    ends_at,
    room,
    status
)
VALUES
    (
        1,
        1,
        1,
        '2026-09-24 09:00:00+03',
        '2026-09-24 10:30:00+03',
        '301',
        'planned'
    ),
    (
        2,
        2,
        2,
        '2026-09-24 11:00:00+03',
        '2026-09-24 12:30:00+03',
        '205',
        'planned'
    );