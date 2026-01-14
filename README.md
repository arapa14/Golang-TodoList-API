# Todo List API (Golang)

Todo List API adalah RESTful backend service sederhana untuk mengelola todo (task) tanpa fitur autentikasi.

Project ini cocok untuk:
- Fondasi backend CRUD
- Dasar sebelum menambahkan auth atau microservice
- Sistem monolith (bukan microservice)

---

## 🎯 Tujuan Aplikasi

- Menyediakan API untuk membuat, membaca, mengubah, dan menghapus todo
- Membuat logika dasar CRUD Golang dengan sistem monolith

---

## 📦 Fitur Utama

- Create Todo
- Get Todo List (pagination, filter, sorting)
- Get Todo by ID
- Update Todo (partial update)
- Delete Todo (soft delete)
- Konsistensi response success & error
- Validation terpusat

❌ Tidak ada autentikasi  
❌ Tidak ada user management  
❌ Tidak microservice 

---

## 🗄️ Database Design

### Table: `todos_tb`

| Column | Type | Description |
|------|------|------------|
| id | UUID / BIGSERIAL | Primary key |
| title | VARCHAR | Judul todo (wajib) |
| description | TEXT | Deskripsi (opsional) |
| is_completed | BOOLEAN | Status todo |
| priority | SMALLINT | 1=low, 2=medium, 3=high |
| due_date | TIMESTAMP | Deadline (opsional) |
| created_at | TIMESTAMP | Waktu dibuat |
| updated_at | TIMESTAMP | Waktu update |
| deleted_at | TIMESTAMP | Soft delete |

---

## 🌐 API Endpoint

Base URL: