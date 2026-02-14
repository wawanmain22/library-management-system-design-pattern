# Panduan Penulisan Laporan Tugas Besar

## Struktur Laporan

Berikut adalah struktur laporan yang disarankan untuk tugas besar ini:

### 1. Pendahuluan

Jelaskan secara ringkas:
- Apa itu Design Pattern
- Mengapa penting mempelajari Design Pattern
- Tujuan tugas besar ini

### 2. Domain Kasus

Jelaskan Library Management System:
- Fungsi utama sistem perpustakaan
- Entities yang ada (Book, Category, User, etc.)
- Use cases utama (peminjaman, pengembalian, pencarian)
- Challenges yang perlu diatasi

### 3. Mapping Design Pattern

Buat table seperti berikut:

| No | Pattern | Kategori | Alasan Pemilihan |
|----|---------|-----------|-------------------|
| 1 | Builder | Creational | Buku memiliki banyak properti optional. Builder memungkinkan pembuatan objek dengan method chaining yang lebih readable |
| 2 | Prototype | Creational | Sering perlu membuat salinan buku untuk lokasi/edisi berbeda. Pattern ini efisien untuk cloning |
| 3 | Abstract Factory | Creational | Sistem perlu mendukung berbagai jenis media (Book, Magazine). Factory pattern memisahkan pembuatan objek berdasarkan jenis |
| 4 | Composite | Structural | Kategori membentuk hierarchy tree (Fiction > Romance > Historical). Pattern ini memungkinkan penanganan uniform |
| 5 | Decorator | Structural | Fitur tambahan seperti "Reserved" atau "Reference Only" ditambahkan secara dinamis tanpa modifikasi class asli |
| 6 | Proxy | Structural | Beberapa buku restricted untuk user tertentu. Pattern ini mengontrol access berdasarkan user tier |
| 7 | State | Behavioral | Buku memiliki berbagai state (Available, Borrowed, Overdue). Perilaku berubah berdasarkan state |
| 8 | Command | Behavioral | Operasi seperti Borrow/Return perlu di-log dan bisa di-undo. Pattern ini encapsulate actions sebagai objects |
| 9 | Strategy | Behavioral | Perlu berbagai metode pencarian (Title, Author). Pattern ini memungkinkan switch algoritma di runtime |

### 4. Implementasi Per Pattern

Untuk setiap pattern, jelaskan:

#### 4.1 Nama Pattern

**Alasan Pemilihan:**
Jelaskan mengapa pattern ini cocok untuk domain Library Management System

**Class Diagram:**
Gunakan ASCII art atau PlantUML format:
```
+----------------+         +----------------+
|  ClassName     |-------->|  OtherClass    |
+----------------+         +----------------+
| - field1       |         | - field1       |
| - field2       |         | - field2       |
+----------------+         +----------------+
| + Method1()    |         | + Method1()    |
| + Method2()    |         | + Method2()    |
+----------------+         +----------------+
```

**Penjelasan Implementasi:**
- Jelaskan file-file yang dibuat dan fungsinya
- Jelaskan struktur code (interface, struct, method)
- Berikan contoh code snippet penting

**Runutan Proses:**
Jelaskan alur execution dengan contoh:
1. Langkah 1: ...
2. Langkah 2: ...
3. Langkah 3: ...
4. ...

**Testing:**
- Jelaskan skenario test yang dilakukan
- Sertakan output program
- Jelaskan hasil dan verifikasi

### 5. Integrasi Pattern

Jelaskan bagaimana pattern-pattern bekerja bersama:
- Bagaimana Builder dan Prototype berinteraksi
- Bagaimana Composite dan State bekerja sama
- Bagaimana Proxy dan Command saling melengkapi
- Workflow lengkap dari pencarian sampai peminjaman

### 6. Hasil dan Pembahasan

#### 6.1 Kelebihan
- Code lebih modular dan maintainable
- Pattern memudahkan ekstensi tanpa modifikasi
- Testing lebih mudah dengan struktur yang jelas

#### 6.2 Tantangan
- Kompleksitas meningkat dengan banyak pattern
- Perlu pemahaman yang baik terhadap setiap pattern
- Balancing antara over-engineering dan flexibility

#### 6.3 Pembelajaran
- Memahami 9 design pattern berbeda
- Menerapkan pattern dalam konteks nyata
- Menggunakan Go idioms dalam implementasi pattern

### 7. Kesimpulan

Ringkas hasil tugas:
- Semua 9 pattern berhasil diimplementasikan
- Sistem Library Management berjalan dengan baik
- Pattern membantu menyelesaikan masalah domain

### 8. Referensi

- Design Patterns: Elements of Reusable Object-Oriented Software - GoF
- Refactoring.Guru - https://refactoring.guru/design-patterns
- Go Patterns - https://github.com/tmrts/go-patterns
- Clean Code - Robert C. Martin

## Tips Penulisan Laporan

### Formatting
- Gunakan font yang konsisten
- Gunakan bullet points untuk kejelasan
- Berikan spasi yang cukup antar section
- Gunakan bold untuk penekanan penting

### Class Diagram
- Gunakan tool seperti:
  - PlantUML: https://plantuml.com
  - Draw.io: https://app.diagrams.net
  - Mermaid: https://mermaid.live
- Pastikan diagram mudah dibaca
- Sertakan keterangan untuk setiap komponen

### Code Snippet
- Sertakan code yang relevan saja
- Jelaskan baris penting dengan komentar
- Format code dengan syntax highlighting
- Hindari code yang terlalu panjang

### Output Program
- Sertakan output dari demo program
- Berikan penjelasan untuk setiap output
- Highlight bagian penting dengan formatting
- Screenshot jika diperlukan

### Link GitHub/Gist
- Upload semua code ke GitHub
- Buat README yang jelas
- Sertakan link repository di laporan
- Pastikan repository public

## Checklist Penyelesaian

Sebelum submit, pastikan:

- [ ] Semua 9 pattern sudah diimplementasikan
- [ ] Setiap pattern memiliki class diagram
- [ ] Setiap pattern memiliki penjelasan implementasi
- [ ] Setiap pattern memiliki testing scenario
- [ ] Code sudah di-upload ke GitHub/Gist
- [ ] Laporan menggunakan bahasa Indonesia yang baku
- [ ] Tidak ada typo atau error gramatikal
- [ ] Formatting konsisten dan rapi
- [ ] Link GitHub/Gist sudah disertakan
- [ ] Nama dan NIM sudah tercantum

## Contoh Output Program untuk Laporan

Gunakan output dari menjalankan `go run main.go` sebagai bukti implementasi:

```
=== BUILDER PATTERN ===
Creating books using Builder pattern with method chaining
Created: Book{ID: BK-1, Title: 1984, Author: George Orwell, ISBN: 9780451524935}
Created: Book{ID: BK-2, Title: Clean Code, Author: Robert C. Martin, ISBN: 9780132350884}
Validation Error: title is required
Book without title: <nil>

=== STATE PATTERN ===
Book: The Alchemist (ISBN: 9780062315007)
  Current State: Available
Book is now borrowed
Book: The Alchemist (ISBN: 9780062315007)
  Current State: Available
Error: cannot mark an available book as overdue
...
```

## Waktu Estimasi

Berikut estimasi waktu untuk menyelesaikan laporan:

- Pendahuluan: 30 menit
- Domain Kasus: 45 menit
- Mapping Pattern: 30 menit
- Implementasi Pattern (9 x 20 menit): 3 jam
- Integrasi Pattern: 45 menit
- Hasil dan Pembahasan: 1 jam
- Kesimpulan: 30 menit
- Final Review: 30 menit

**Total estimasi: 7 jam**

## File yang Disediakan

Project ini menyediakan:
1. `main.go` - Demo program lengkap
2. `patterns/` - Implementasi semua 9 pattern
3. `doc/IMPLEMENTATION_NOTES.md` - Detail setiap pattern
4. `doc/PATTERN_SUMMARY.md` - Hubungan antar pattern
5. `README.md` - Dokumentasi project

Gunakan dokumentasi ini sebagai referensi dalam penulisan laporan!
