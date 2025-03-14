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

## TODO

- [ ] Membuat kontrak untuk service user
  - [ ] `Create` : Membuat user baru
    - Masukkan setidakya semua data wajib non-default dari skema
    - Validasi data
    - Hash password
    - Save ke database

  - [ ] `Delete` : Hapus akun berdasarkan ID
    - Cari user berdasarkan id
    - Verifikasi apakah user merupakan pemilik akun / admin
    - Soft delete user dengan menambahkan deleted_at

  - [ ] `List` : Daftar user dengan filter
  - [ ] `GetOther` : Mengambil informasi akun lain
  - [ ] `GetSelf` : Mengambil informasi akun sendiri
  - [ ] `Update` : Update data umum user
- [ ] Membuat kontrak untuk service auth
  - [ ] `ChangeEmail` : Mengubah email user
  - [ ] `ChangePassword` : Mengubah password dari user
  - [ ] `ChangeUsername` : Mengubah username dari user
  - [ ] `LogIn` : Membuat akses dan refresh JWT token
  - [ ] `VerifyToken`: Mengubah jwt ke data kredensial user
