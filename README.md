# User Service Proto Deklarasi dan Kontrak Go

## Skema User

| Nama Kolom | Tipe Data | Keterangan      | Nilai Default     |
| ---------- | --------- | --------------- | ----------------- |
| id         | UUID      | PRIMARY KEY     | gen_random_uuid() |
| email      | VARCHAR   | UNIQUE NOT NULL | -                 |
| username   | VARCHAR   | UNIQUE NOT NULL | -                 |
| password   | VARCHAR   | NOT NULL        | -                 |
| first_name | VARCHAR   | NOT NULL        | -                 |
| last_name  | VARCHAR   | NULL            | -                 |
| created_at | TIMESTAMP | NOT NULL        | CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | NOT NULL        | CURRENT_TIMESTAMP |
| deleted_at | TIMESTAMP | NULL            | -                 |



