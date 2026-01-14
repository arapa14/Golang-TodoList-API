Todo List API (Golang)

Todo List API adalah RESTful backend service monolith yang dibangun menggunakan Golang + PostgreSQL untuk mengelola data todo (task) tanpa autentikasi.

Project ini difokuskan pada:

Struktur backend yang rapi & scalable

Best practice response API

Pagination otomatis

Penanganan data nullable yang benar

Cocok sebagai fondasi sebelum berkembang ke microservice

🎯 Tujuan Aplikasi

Menyediakan API CRUD untuk todo

Menerapkan pola monolith yang terstruktur

Menjadi baseline backend Golang sebelum:

menambahkan autentikasi

memisahkan ke microservice

menambahkan domain lain

📦 Fitur Utama

✅ Create Todo
✅ Get Todo List (pagination otomatis)
✅ Get Todo by ID (planned)
✅ Update Todo (partial update / PATCH) (planned)
✅ Delete Todo (soft delete) (planned)
✅ Response API konsisten (success & error)
✅ Pagination meta otomatis
✅ Nullable field handling (NULL → pointer)
✅ Validation dasar request

❌ Tidak ada autentikasi
❌ Tidak ada user management
❌ Belum microservice

🧱 Arsitektur

Project ini MASIH monolith, namun sudah disiapkan agar:

Mudah dipisah ke microservice

Logic tidak tercampur (helper, pagination, response)

TODO-LIST-API/
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── shared/
│   │   ├── response.go      # RespondSuccess & RespondError
│   │   ├── pagination.go   # pagination helper
│   │   └── db.go            # helper DB generic (CountRows)
│   │
│   └── todo/                # (akan dipisah nanti)
│
├── infrastructure/
│   └── database/
│       └── postgres.go
│
├── go.mod
└── README.md

🗄️ Database Design
Table: todos_tb
Column	Type	Nullable	Description
id	BIGSERIAL	❌	Primary key
title	VARCHAR	❌	Judul todo
description	TEXT	✅	Deskripsi
is_completed	BOOLEAN	❌	Status
priority	SMALLINT	❌	1=low, 2=medium, 3=high
due_date	TIMESTAMP	✅	Deadline
created_at	TIMESTAMP	❌	Created time
updated_at	TIMESTAMP	❌	Updated time
deleted_at	TIMESTAMP	✅	Soft delete
📦 Data Model (API)
Todo Response
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Clean architecture",
  "is_completed": false,
  "priority": 2,
  "due_date": "2026-01-20",
  "created_at": "2026-01-14T09:00:00Z",
  "updated_at": "2026-01-14T09:00:00Z"
}


Field nullable tidak dikirim jika NULL (menggunakan pointer + omitempty)

🌐 API Endpoint

Base URL:

/api/v1

▶️ Get Todo List
GET /api/v1/todos?page=1&limit=10


Response:

{
  "status": "success",
  "code": "OK",
  "message": "Get todo",
  "data": {
    "items": [...],
    "meta": {
      "page": 1,
      "limit": 10,
      "total_items": 100,
      "total_pages": 10
    }
  }
}

▶️ Create Todo
POST /api/v1/todos


Request body:

{
  "title": "Belajar Golang",
  "description": "API monolith",
  "priority": 2,
  "due_date": "2026-01-20"
}


Response:

{
  "status": "success",
  "code": "CREATED",
  "message": "Post todo",
  "data": {
    "items": {
      "id": 1,
      "title": "Belajar Golang",
      "priority": 2
    }
  }
}

🧪 Validation Rules

title wajib dan tidak boleh kosong

priority harus valid

JSON harus valid

Error response konsisten

🧠 Response Format (Standar)
Success
{
  "status": "success",
  "code": "OK",
  "message": "Message",
  "data": {
    "items": {},
    "meta": {}
  }
}

Error
{
  "status": "error",
  "code": "BAD_REQUEST",
  "message": "Error message"
}

🚀 Development Plan (Next)

 Get Todo by ID

 Update Todo (PATCH)

 Soft delete

 Repository & usecase separation

 Unit testing

 Transition monolith → microservice

🏁 Catatan

Project ini sengaja tidak overengineering, namun:

Sudah mengikuti best practice Golang API

Mudah dikembangkan

Siap direfactor ke microservice