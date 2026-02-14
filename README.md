# Library Management System - Design Pattern Implementation

Tugas Besar Mata Kuliah **Clean Code dan Design Pattern** - Universitas Komputer Indonesia (UNIKOM) 2025.

Implementasi 5 Design Pattern pada domain Library Management System menggunakan bahasa pemrograman Go.

## Design Patterns yang Diimplementasikan

### Creational Patterns (2)
| No | Pattern | Kegunaan |
|----|---------|----------|
| 1 | **Builder** | Membuat objek Book secara fleksibel dengan method chaining dan validasi field wajib |
| 2 | **Prototype** | Cloning book untuk variasi edisi/lokasi tanpa rebuild dari awal (deep copy) |

### Structural Patterns (1)
| No | Pattern | Kegunaan |
|----|---------|----------|
| 3 | **Decorator** | Menambahkan fitur Reserved dan Reference Only secara dinamis tanpa modifikasi class asli |

### Behavioral Patterns (2)
| No | Pattern | Kegunaan |
|----|---------|----------|
| 4 | **State** | Mengelola lifecycle buku (Available → Borrowed → Overdue) dengan aturan transisi yang ketat |
| 5 | **Strategy** | Pencarian buku dengan algoritma yang bisa diganti di runtime (Title Search, Author Search) |

## Struktur Project

```
library-management-system/
├── main.go                                    # Entry point & demo semua patterns
├── go.mod                                     # Go module definition
├── Laporan_Design_Pattern.docx                # Laporan tugas besar
├── doc/
│   ├── IMPLEMENTATION_NOTES.md                # Dokumentasi detail setiap pattern
│   ├── PATTERN_SUMMARY.md                     # Ringkasan hubungan antar pattern
│   ├── class-diagram-*.png                    # Class diagram per pattern
│   └── txt/                                   # Source code & output per pattern (txt)
│       ├── 1_builder_code.txt
│       ├── 1_builder_output.txt
│       ├── 2_prototype_code.txt
│       ├── 2_prototype_output.txt
│       ├── 3_decorator_code.txt
│       ├── 3_decorator_output.txt
│       ├── 4_state_code.txt
│       ├── 4_state_output.txt
│       ├── 5_strategy_code.txt
│       └── 5_strategy_output.txt
├── patterns/
│   ├── creational/
│   │   ├── builder/
│   │   │   ├── book.go                        # Book struct dengan properti lengkap
│   │   │   └── book_builder.go                # Builder dengan method chaining & validasi
│   │   └── prototype/
│   │       ├── book.go                        # Book dengan Clone() deep copy
│   │       └── prototype.go                   # PrototypeManager untuk registry
│   ├── structural/
│   │   └── decorator/
│   │       ├── book.go                        # Base Book (Concrete Component)
│   │       ├── book_decorator.go              # BookComponent interface & BaseDecorator
│   │       ├── reserved_decorator.go          # Reserved book decorator
│   │       └── reference_only_decorator.go    # Reference only decorator
│   └── behavioral/
│       ├── state/
│       │   ├── book_state.go                  # BookState interface
│       │   ├── book.go                        # Book context
│       │   ├── available_state.go             # Available state
│       │   ├── borrowed_state.go              # Borrowed state
│       │   └── overdue_state.go               # Overdue state
│       └── strategy/
│           ├── search_strategy.go             # SearchStrategy interface
│           ├── book.go                        # Book data model
│           ├── title_search.go                # Title search strategy
│           ├── author_search.go               # Author search strategy
│           └── catalog.go                     # Catalog context
└── README.md
```

## Cara Menjalankan

### Prerequisites
- Go 1.25 atau lebih tinggi

### Run Program

```bash
# Clone repository
git clone https://github.com/wawanmain22/library-management-system-design-pattern.git
cd library-management-system-design-pattern

# Install dependencies
go mod tidy

# Jalankan program
go run main.go
```

### Build Binary

```bash
go build -o library-system main.go
./library-system
```

## Output Program

Program menampilkan demo untuk setiap pattern secara sequential:

### 1. Builder Pattern
- Pembuatan buku lengkap dengan semua field dan tags
- Pembuatan buku partial (field wajib saja)
- Validasi error ketika field wajib kosong
- Verifikasi ID unik per buku

### 2. Prototype Pattern
- Registrasi prototype ke PrototypeManager
- Cloning dan modifikasi clone (ISBN, Price, Stock)
- Verifikasi original tidak terpengaruh (deep copy)

### 3. Decorator Pattern
- Base book dengan CanBorrow() = true
- Reserved decorator: CanBorrow() = false
- Reference Only decorator: CanBorrow() = false, CanReadInLibrary() = true
- GetDetails() menampilkan info tambahan dari decorator

### 4. State Pattern
- State awal: Available
- Borrow → state berubah ke Borrowed
- Borrow lagi → error (sudah dipinjam)
- MarkOverdue → state berubah ke Overdue
- Return → state kembali ke Available
- Return lagi → error (sudah available)

### 5. Strategy Pattern
- Title Search: query "Clean" → 2 hasil
- Author Search: query "Robert" → 2 hasil
- Switch strategy ke Title Search: query "Design" → 1 hasil

## Teknologi

- **Bahasa**: Go (Golang) 1.25+
- **Dependencies**: Standard library only (tidak ada external dependencies)
- **Paradigma**: Interface-based design, composition over inheritance

## Clean Code Principles

Project ini mengikuti prinsip SOLID:
- **Single Responsibility**: Setiap file dan struct memiliki satu tanggung jawab
- **Open/Closed**: Pattern memungkinkan ekstensi tanpa modifikasi kode existing
- **Liskov Substitution**: Concrete types dapat menggantikan interface
- **Interface Segregation**: Interface terfokus dan minimal
- **Dependency Inversion**: Bergantung pada abstraksi (interface), bukan concrete

## Author

**Ridwan Syarif Abidin** (10122053)
Teknik Informatika - Universitas Komputer Indonesia
Mata Kuliah: Clean Code dan Design Pattern
Dosen Pengampu: Alif Finandhita, S.Kom, M.T

## License

Project ini dibuat untuk keperluan Tugas Besar mata kuliah Clean Code dan Design Pattern, UNIKOM 2025.
