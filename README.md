# CineReview

CineReview adalah aplikasi katalog dan rating film berbasis bahasa pemrograman Go yang digunakan untuk mendata daftar film favorit pengguna.

Project ini dibuat untuk memenuhi Tugas Besar Mata Kuliah Algoritma dan Pemrograman 2 dengan penerapan konsep dasar algoritma dan struktur data.

---

# Deskripsi Aplikasi

CineReview membantu pengguna untuk:

* Menambahkan data film
* Menampilkan daftar film
* Menghapus data film
* Mencari film berdasarkan judul
* Mengurutkan film berdasarkan rating

Data film yang disimpan meliputi:

* Judul film
* Genre
* Tahun rilis
* Deskripsi singkat
* Rating pengguna

---

# Konsep yang Digunakan

* Struct
* Function / Subprogram
* Array atau Slice
* Sequential Search
* Selection Sort
* Perulangan dan Percabangan

---

# Struktur Data

```go
type Film struct {
    Judul string
    Genre string
    Tahun int
    Deskripsi string
    Rating float64
}
```

---

# Cara Menjalankan Program

```bash
go run main.go
```

---

# Progress Saat Ini

## ✅ Fitur yang Sudah Selesai

* Tambah data film
* Menampilkan daftar film
* Hapus data film
* Sequential Search berdasarkan judul
* Selection Sort berdasarkan rating

## ⏳ Fitur yang Masih Dalam Pengembangan

* Edit data film
* Binary Search
* Insertion Sort
* Statistik film

---

# Teknologi

* Golang
* CLI / Terminal

---

# Tim Pengembang

* Adhara Faliya Utanti - 109082500033
* Shasa Olivia Rose - 109082500207
