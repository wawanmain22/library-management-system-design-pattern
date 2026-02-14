# Design Pattern Summary - Library Management System

## Overview

Dokumentasi ini menyediakan ringkasan singkat tentang 5 design pattern yang diimplementasikan dalam Library Management System, termasuk penggunaannya bersama-sama dan alur kerja sistem secara keseluruhan.

## Pattern Interdependency

### Bagaimana Pattern Bekerja Bersama

1. **Builder -> Prototype**
   - Builder pattern digunakan untuk membuat Book objects
   - Prototype kemudian bisa clone Book yang sudah dibuat
   - Contoh: Builder buat buku master, Prototype buat banyak copy

2. **Decorator -> State**
   - Decorator menambahkan fitur ke book (Reserved, Reference Only)
   - State mengelola lifecycle book tersebut
   - Contoh: Book di-decorate dengan Reserved, state mengatur Available/Borrowed

3. **Strategy -> Builder**
   - Strategy mencari books di catalog
   - Builder membuat book objects yang disimpan di catalog
   - Contoh: Builder buat books, Strategy cari berdasarkan title/author

## Domain Workflow Lengkap

### Alur Peminjaman Buku

1. **Pencarian Buku (Strategy Pattern)**
   - User memilih strategy pencarian (Title/Author)
   - Catalog menggunakan strategy untuk mencari buku

2. **Cek Ketersediaan (State Pattern)**
   - Book dicek apakah Available, Borrowed, atau Overdue
   - Jika Available, proses lanjut
   - Jika tidak, error ditampilkan

3. **Peminjaman (State Pattern)**
   - Book.Borrow() dipanggil
   - State mengelola transisi Available -> Borrowed

4. **Penambahan Fitur (Decorator Pattern)**
   - Jika perlu, book di-decorate
   - Contoh: Reserved, Reference Only
   - Decorator mengubah behavior tanpa modifikasi original

### Alur Pengembalian Buku

1. **Pengembalian (State Pattern)**
   - Book.Return() dipanggil
   - State transisi dari Borrowed/Overdue -> Available

### Alur Manajemen Koleksi

1. **Pengaturan Properti (Builder Pattern)**
   - Book dibuat dengan properti kompleks
   - Method chaining untuk readability
   - Validasi field wajib

2. **Duplikasi (Prototype Pattern)**
   - Jika perlu, book bisa di-cloning
   - Berguna untuk multiple copies
   - Efisien pembuatan similar objects

## Kategori Pattern

### Creational Patterns (Object Creation)

| Pattern | Fokus | Benefit |
|----------|--------|----------|
| Builder | Pembuatan objek kompleks | Method chaining, validasi |
| Prototype | Cloning objek | Efisiensi untuk multiple copies |

### Structural Patterns (Object Composition)

| Pattern | Fokus | Benefit |
|----------|--------|----------|
| Decorator | Dynamic behavior | Add features tanpa modifikasi |

### Behavioral Patterns (Object Interaction)

| Pattern | Fokus | Benefit |
|----------|--------|----------|
| State | State transitions | Encapsulate state-specific behavior |
| Strategy | Algorithm selection | Runtime algorithm switching |

## Complexity Analysis

### Tingkat Kompleksitas

**Low Complexity (Mudah Dipahami)**
- Builder: Method chaining, intuitive
- Prototype: Simple cloning
- Decorator: Wrapper pattern sederhana

**Medium Complexity (Moderate)**
- Strategy: Runtime algorithm switching

**High Complexity (Complex)**
- State: Multiple states dan transitions

### Mapping ke Domain Library Management

| Pattern | Relevance | Implementation Complexity | Testing Complexity |
|----------|-------------|------------------------|-------------------|
| Builder | Tinggi | Rendah | Rendah |
| Prototype | Sedang | Rendah | Sedang |
| Decorator | Tinggi | Rendah | Sedang |
| State | Tinggi | Tinggi | Tinggi |
| Strategy | Tinggi | Sedang | Sedang |

## Best Practices yang Diimplementasikan

1. **Interface over Implementation**
   - Semua pattern menggunakan interface
   - Memungkinkan dependency injection
   - Membuat testing mudah

2. **Error Handling**
   - Proper error messages
   - Validation di constructor/builder
   - Graceful failure handling

3. **Clean Code**
   - Clear naming conventions
   - Single responsibility per type
   - Minimal coupling antar komponen

4. **Go Idioms**
   - Implicit interfaces
   - Pointer receivers untuk mutasi
   - Error handling pattern
   - Composisi over inheritance

## Testing Strategy

### Unit Testing Per Pattern

Setiap pattern diuji dengan:

1. **Happy Path**
   - Skenario normal yang berjalan benar
   - Output sesuai ekspektasi

2. **Error Path**
   - Skenario edge cases
   - Error handling teruji
   - System tidak crash

3. **Integration**
   - Pattern bekerja dengan pattern lain
   - Interdependencies teruji

### Output Verification

Program demo menghasilkan output yang jelas:
- Setiap pattern memiliki section terpisah
- Step-by-step execution
- Clear status updates
- Error messages informatif

## Kesimpulan

Lima design pattern yang diimplementasikan menyediakan:

1. **Fleksibilitas**: Pattern memungkinkan sistem beradaptasi dengan perubahan
2. **Maintainability**: Code lebih mudah dimaintain karena struktur jelas
3. **Extensibility**: Fitur baru dapat ditambahkan tanpa memodifikasi existing code
4. **Reusability**: Pattern dapat digunakan di domain lain
5. **Testability**: Setiap pattern dapat diuji secara independen

Domain Library Management System memanfaatkan pattern ini secara efektif karena:
- Banyak objek dengan kompleksitas berbeda
- Berbagai workflows (peminjaman, pengembalian, pencarian)
- Requirement untuk berbagai algoritma pencarian
- Kebutuhan menambahkan fitur dinamis pada buku

Implementasi ini menunjukkan bagaimana design pattern dapat dipilih dan diimplementasikan untuk menyelesaikan masalah nyata dalam sistem perangkat lunak.
