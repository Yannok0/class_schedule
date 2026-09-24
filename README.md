# class_schedule

## API

| Метод | Путь | Тело запроса | Успех | Ошибки |
|---|---|---|---|---|
| GET | /lessons | — | 200, массив | 500 |
| GET | /lessons/{id} | — | 200, объект | 400, 404 |
| POST | /lessons | {"group_id":1,"teacher_id":1,"subject_id":1,"starts_at":"2026-09-24T09:00:00","ends_at":"2026-09-24T10:30:00","room":"305"} | 201, объект | 400, 422 |
| PATCH | /lessons/{id} | {"status":"completed"} | 200, объект | 400, 404, 422 |
| DELETE | /lessons/{id} | — | 204 | 400, 404 |

### Объект Lesson

```json
{
  "id": 1,
  "group_id": 1,
  "teacher_id": 1,
  "subject_id": 1,
  "starts_at": "2026-09-24T09:00:00Z",
  "ends_at": "2026-09-24T10:30:00Z",
  "room": "305",
  "status": "scheduled"
}