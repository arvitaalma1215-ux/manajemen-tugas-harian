# Aplikasi Manajemen Tugas Harian (Juni 2026)

Aplikasi ini adalah program berbasis bahasa pemrograman **Go** yang dirancang untuk membantu mahasiswa atau pengguna dalam mengelola daftar tugas harian secara terstruktur. Proyek ini dibuat sebagai implementasi konsep **Struktur Data dan Algoritma** dengan menerapkan berbagai teknik pencarian dan pengurutan data.

## 📌 Fitur Utama

### 📁 Manajemen Data Tugas
* **Tambah Tugas**: Menambahkan tugas baru (Mata Kuliah, Kategori, Deadline, Jam) ke dalam sistem.
* **Tampilkan Semua Tugas**: Melihat daftar lengkap tugas yang telah disimpan.
* **Ubah Tugas**: Memperbarui detail tugas yang sudah ada.
* **Hapus Tugas**: Menghapus data tugas dari sistem.
* **Ubah Status Tugas**: Menandai tugas menjadi "Selesai" atau "Belum Selesai".

### 🔍 Fitur Pencarian
* **Kategori**: Mencari tugas berdasarkan kategori tertentu (*Sequential Search*).
* **Deadline**: Mencari tugas berdasarkan tanggal deadline (*Binary Search*).
* **Spesifik**: Mencari tugas berdasarkan kombinasi Tanggal dan Jam (*Sequential Search*).

### 📊 Pengurutan Data
* **Berdasarkan Deadline**: Menggunakan *Selection Sort* (Ascending/Descending).
* **Berdasarkan Status**: Menggunakan *Insertion Sort* (Ascending/Descending).

### 💡 Analisis & Feedback
* **Validasi**: Batasan input khusus untuk bulan Juni 2026.
* **Feedback Otomatis**: Memberikan rekomendasi atau peringatan berdasarkan sisa hari menuju deadline.

---

## 🛠️ Algoritma yang Digunakan

| Fitur | Algoritma |
| :--- | :--- |
| Cari Kategori | Sequential Search |
| Cari Deadline | Binary Search |
| Urutkan Deadline | Selection Sort |
| Urutkan Status | Insertion Sort |

---

## 🚀 Cara Menjalankan

1. **Clone repository**
   ```bash
   git clone [https://github.com/arvitaalma1215-ux/manajemen-tugas-harian.git](https://github.com/arvitaalma1215-ux/manajemen-tugas-harian.git)
2. ```bash
   cd manajemen-tugas-harian
3. ```bash
   go run main.go
