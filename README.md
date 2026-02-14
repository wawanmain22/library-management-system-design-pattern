# Library Management System - Design Pattern Implementation

Project ini adalah implementasi sistem manajemen perpustakaan yang menggabungkan 9 design pattern dari 3 kategori: Creational, Structural, dan Behavioral. Project ini dibuat sebagai tugas besar mata kuliah Clean Code dan Design Pattern.

## Design Patterns yang Diimplementasikan

### Creational Patterns (3)
1. **Builder Pattern** - Membuat complex Book objects dengan method chaining
2. **Prototype Pattern** - Cloning book instances untuk multiple copies
3. **Abstract Factory Pattern** - Membuat berbagai jenis media (Book, Magazine)

### Structural Patterns (3)
4. **Composite Pattern** - Mengelola category hierarchy
5. **Decorator Pattern** - Menambahkan fitur ke book secara dinamis
6. **Proxy Pattern** - Control access ke restricted books

### Behavioral Patterns (3)
7. **State Pattern** - Mengelola state transitions book (Available, Borrowed, Overdue)
8. **Command Pattern** - Encapsulate dan undo operations (Borrow, Return)
9. **Strategy Pattern** - Multiple search strategies (Title, Author)

## Struktur Project

```
library-management-system/
├── main.go                              # Demo program untuk semua patterns
├── go.mod                                # Go module definition
├── doc/
│   └── IMPLEMENTATION_NOTES.md            # Dokumentasi detail setiap pattern
├── patterns/
│   ├── creational/
│   │   ├── builder/
│   │   │   ├── book.go                   # Book struct
│   │   │   └── book_builder.go           # Builder dengan method chaining
│   │   ├── prototype/
│   │   │   ├── book.go                  # Book dengan Clone() method
│   │   │   └── prototype.go             # PrototypeManager
│   │   └── abstract_factory/
│   │       ├── media_item.go             # MediaItem interface
│   │       ├── media_factory.go          # FactoryProducer
│   │       ├── book_factory.go           # Concrete factory untuk Book
│   │       └── magazine_factory.go       # Concrete factory untuk Magazine
│   ├── structural/
│   │   ├── composite/
│   │   │   ├── library_item.go          # Interface untuk composite
│   │   │   ├── book.go                 # Leaf node
│   │   │   └── category.go             # Composite node
│   │   ├── decorator/
│   │   │   ├── book.go                 # Base Book
│   │   │   ├── book_decorator.go        # Decorator interface
│   │   │   ├── reserved_decorator.go     # Reserved flag decorator
│   │   │   └── reference_only_decorator.go # Reference only decorator
│   │   └── proxy/
│   │       ├── book.go                 # Real Book
│   │       ├── book_proxy.go           # Proxy dengan access control
│   │       └── user.go                # User dengan tier
│   └── behavioral/
│       ├── state/
│       │   ├── book_state.go           # State interface
│       │   ├── available_state.go      # Available state
│       │   ├── borrowed_state.go       # Borrowed state
│       │   ├── overdue_state.go        # Overdue state
│       │   └── book.go               # Context
│       ├── command/
│       │   ├── command.go             # Command interface
│       │   ├── borrow_command.go      # Borrow command
│       │   ├── return_command.go      # Return command
│       │   ├── invoker.go            # Command invoker
│       │   └── receiver.go           # Library receiver
│       └── strategy/
│           ├── search_strategy.go      # Strategy interface
│           ├── title_search.go        # Search by title
│           ├── author_search.go       # Search by author
│           ├── catalog.go            # Catalog dengan strategy
│           └── book.go              # Book struct
└── README.md                        # File ini
```

## Cara Menjalankan

### Prerequisites
- Go 1.25 atau lebih tinggi

### Running the Program

1. Clone atau download project
2. Masuk ke directory project:
   ```bash
   cd library-management-system
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the program:
   ```bash
   go run main.go
   ```

Program akan menampilkan demo untuk setiap design pattern secara sequential.

## Output Program

Program menghasilkan output yang mencakup:

1. **Builder Pattern Demo**
   - Pembuatan buku lengkap dengan semua field
   - Pembuatan buku partial
   - Validasi error untuk field wajib

2. **Prototype Pattern Demo**
   - Registrasi prototype
   - Cloning book dan modifikasi
   - Verifikasi original tidak terpengaruh

3. **Abstract Factory Demo**
   - Pembuatan Book menggunakan BookFactory
   - Pembuatan Magazine menggunakan MagazineFactory
   - Switch factory di runtime

4. **Composite Pattern Demo**
   - Membuat category hierarchy
   - Display hierarchy tree
   - Total count calculation
   - Search dalam hierarchy

5. **Decorator Pattern Demo**
   - Base book functionality
   - Decorate dengan Reserved
   - Decorate dengan Reference Only
   - Behavior changes

6. **Proxy Pattern Demo**
   - Student akses public book
   - Student akses restricted book
   - Faculty akses restricted book
   - Access log display

7. **State Pattern Demo**
   - Borrow available book
   - State transition Available -> Borrowed
   - Mark overdue
   - Return book
   - Invalid state transitions

8. **Command Pattern Demo**
   - Execute borrow commands
   - Execute return command
   - Undo operation
   - Command history

9. **Strategy Pattern Demo**
   - Search by title
   - Switch ke author search
   - Display hasil berbeda
   - Multiple search queries

## Dokumentasi

Untuk dokumentasi detail setiap pattern, class diagram, dan penjelasan implementasi, lihat:
- `doc/IMPLEMENTATION_NOTES.md`

Dokumentasi mencakup:
- Alasan pemilihan setiap pattern
- Class diagram (ASCII art)
- Detail implementasi
- Testing scenarios
- Hubungan antar pattern

## Teknologi

- **Language**: Go (Golang) 1.25+
- **No External Dependencies**: Menggunakan hanya standard library Go

## Clean Code Principles

Project ini mengikuti prinsip clean code:
- **Single Responsibility**: Setiap struct memiliki satu tanggung jawab
- **Open/Closed**: Pattern memungkinkan ekstensi tanpa modifikasi
- **Liskov Substitution**: Interface bisa digantikan dengan implementasi
- **Interface Segregation**: Interface terfokus dan minimal
- **Dependency Inversion**: Depend pada abstraksi (interface), bukan concrete

## Design Pattern Selection Rationale

Pattern dipilih berdasarkan:
- **Relevance**: Sesuai dengan domain Library Management System
- **Practicality**: Memiliki manfaat nyata dalam sistem
- **Educational Value**: Memberikan pemahaman berbagai kategori pattern
- **Go Idiomatic**: Cocok dengan pendekatan Go (interface-based design)

## Author

Nama: [ISI NAMA ANDA]
Kelas: 2025/2026 Gasal - Clean Code Dan Design Pattern
Dosen: Alif Finandhita, S.Kom, M.T

## License

Project ini dibuat untuk keperluan tugas kuliah.
