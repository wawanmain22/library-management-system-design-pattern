# Library Management System - Design Pattern Implementation Notes

## Domain Overview

Library Management System adalah sistem yang digunakan untuk mengelola operasi perpustakaan seperti pengelolaan buku, peminjaman, dan pengembalian. Sistem ini diimplementasikan menggunakan bahasa pemrograman Go dengan menerapkan 5 design pattern dari 3 kategori: Creational, Structural, dan Behavioral.

## Design Patterns yang Diimplementasikan

### 1. BUILDER PATTERN (Creational)

**Alasan Pemilihan:**
Buku memiliki banyak properti yang bisa diisi secara optional. Builder pattern memungkinkan pembuatan objek buku dengan method chaining yang lebih readable dan flexible dibandingkan constructor dengan banyak parameter.

**Class Diagram:**
```
+----------------+         +----------------+
|  BookBuilder   |-------->|      Book      |
+----------------+         +----------------+
| - title         |         | - id           |
| - author        |         | - title        |
| - isbn          |         | - author       |
| - publisher     |         | - isbn         |
| - category      |         | - publisher    |
| - pageCount     |         | - category     |
| - language      |         | - pageCount    |
| - published     |         | - language     |
| - tags          |         | - published    |
+----------------+         | - tags         |
| + SetTitle()    |         +----------------+
| + SetAuthor()   |         | + String()     |
| + SetISBN()     |
| + SetPublisher()|
| + SetCategory() |
| + SetPageCount()|
| + SetLanguage() |
| + SetPublished()|
| + AddTag()     |
| + Build() Book |
+----------------+
```

**Implementasi:**
- `book.go`: Mendefinisikan struct Book dengan semua properti yang diperlukan
- `book_builder.go`: Builder dengan method chaining untuk set properti satu per satu
- Validasi field wajib (title dan author) di method Build()
- Menggunakan counter global untuk generate ID unik

**Testing Scenario:**
- Membuat buku lengkap dengan semua field
- Membuat buku partial (hanya title dan author)
- Membuat buku tanpa title (harus error)
- Verifikasi ID unik untuk setiap buku

---

### 2. PROTOTYPE PATTERN (Creational)

**Alasan Pemilihan:**
Dalam sistem perpustakaan, seringkali perlu membuat salinan buku yang sama (misalnya untuk edisi berbeda atau lokasi berbeda). Prototype pattern memungkinkan cloning objek tanpa harus mendefinisikan dari awal.

**Class Diagram:**
```
+-----------------+         +----------------+
|   Prototype     |         |     Book       |
+-----------------+         +----------------+
| + Clone()       |         | - id           |
| + GetDetails()  |-------->| - title        |
+-----------------+         | - author       |
        ^                  | - isbn         |
        |                  | - publisher    |
        |                  | - category      |
+-----------------+        | - price        |
|PrototypeManager|        | - stock        |
+-----------------+        +----------------+
| - prototypes    |        | + Clone()      |
|                |        | + GetDetails() |
| + Register()   |        | + UpdateISBN() |
| + Get()        |        | + UpdatePrice()|
| + List()       |        | + UpdateStock()|
+-----------------+        +----------------+
```

**Implementasi:**
- `book.go`: Book struct dengan Clone() method untuk deep copy
- `prototype.go`: PrototypeManager untuk mengelola prototype collection
- Clone method membuat book baru dengan data sama tapi ID berbeda
- PrototypeManager memungkinkan register dan retrieve prototype dengan key

**Testing Scenario:**
- Register book sebagai prototype
- Clone prototype dan modifikasi beberapa field
- Verify original tidak terpengaruh oleh modifikasi clone
- Clone multiple kali untuk variasi berbeda

---

### 3. DECORATOR PATTERN (Structural)

**Alasan Pemilihan:**
Fitur tambahan pada buku seperti "Reserved" atau "Reference Only" dapat ditambahkan secara dinamis tanpa memodifikasi class Book original. Decorator pattern memungkinkan penambahan behavior secara fleksibel.

**Class Diagram:**
```
+-------------------+
|  BookComponent    |  <<interface>>
+-------------------+
| + GetTitle()      |
| + GetAuthor()     |
| + GetISBN()       |
| + GetDetails()    |
| + CanBorrow()     |
| + CanReadInLib()  |
| + Borrow()        |
| + Return()        |
+-------------------+
        ^
        |
        | +---------------------------+
        |                           |
+-------------------+  +------------------------+
|      Book         |  |  BaseBookDecorator     |
+-------------------+  +------------------------+
| - title           |  | - book BookComponent   |
| - author          |  +------------------------+
| - isbn            |  | + GetTitle()           |
| - publisher       |  | + GetAuthor()          |
| - copies          |  | + GetISBN()            |
+-------------------+  | + GetDetails()         |
| + GetTitle()      |  | + CanBorrow()          |
| + GetAuthor()     |  | + CanReadInLib()       |
| + GetISBN()       |  | + Borrow()             |
| + GetDetails()    |  | + Return()             |
| + CanBorrow()     |  +------------------------+
| + CanReadInLib()  |           ^
| + Borrow()        |           |
| + Return()        |     +-----+-----+
+-------------------+     |           |
              +-----------------------+  +-----------------------+
              |ReservedBookDecorator |  |ReferenceOnlyDecorator |
              +-----------------------+  +-----------------------+
              | - reservedBy         |  |                       |
              +-----------------------+  +-----------------------+
              | + GetDetails()       |  | + GetDetails()        |
              | + CanBorrow()        |  | + CanBorrow()         |
              +-----------------------+  +-----------------------+
```

**Implementasi:**
- `book.go`: Concrete Component - base Book dengan basic functionality
- `book_decorator.go`: BookComponent interface + BaseBookDecorator untuk DRY delegation
- `reserved_decorator.go`: Concrete Decorator - embeds BaseBookDecorator, override GetDetails() dan CanBorrow()
- `reference_only_decorator.go`: Concrete Decorator - embeds BaseBookDecorator, override GetDetails() dan CanBorrow()
- Decorator bisa di-stack (multiple decorators pada satu book)

**Testing Scenario:**
- Decorate book dengan ReservedDecorator
- Decorate book dengan ReferenceOnlyDecorator
- Verify decorator tidak memodifikasi original book
- Stack multiple decorators
- Test CanBorrow() dan CanReadInLibrary() berdasarkan decorator

---

### 4. STATE PATTERN (Behavioral)

**Alasan Pemilihan:**
Buku memiliki berbagai state: Available, Borrowed, Overdue. Perilaku buku berubah berdasarkan state (tidak bisa borrow jika Overdue). State pattern memungkinkan perubahan behavior dinamis berdasarkan internal state.

**Class Diagram:**
```
+-----------------+
|   BookState     |
+-----------------+
| + Borrow()      |
| + Return()      |
| + MarkOverdue() |
| + GetName()     |
+-----------------+
        ^
        |
        | +------------+------------+------------+
        |            |            |            |
+----------------+  +----------------+  +----------------+
|AvailableState  | |BorrowedState  | |OverdueState   |
+----------------+  +----------------+  +----------------+
| - book          |  | - book          |  | - book          |
+----------------+  +----------------+  +----------------+
| + Borrow() OK   |  | + Borrow() FAIL |  | + Borrow() FAIL |
| + Return() FAIL |  | + Return() OK   |  | + Return() OK   |
| + Mark() FAIL   |  | + Mark() OK    |  | + Mark() FAIL   |
| + Name()        |  | + Name()        |  | + Name()        |
+----------------+  +----------------+  +----------------+
        ^              ^              ^
        |              |              |
        +--------------+--------------+
                       |
                +----------------+
                |     Book       |
                +----------------+
                | - title        |
                | - isbn         |
                | - state        |
                +----------------+
                | + SetState()   |
                | + Borrow()     |
                | + Return()     |
                | + Mark()       |
                | + GetName()    |
                +----------------+
```

**Implementasi:**
- `book_state.go`: Interface untuk semua states
- `available_state.go`: State ketika book available (bisa di-borrow)
- `borrowed_state.go`: State ketika book borrowed (bisa di-return atau di-mark overdue)
- `overdue_state.go`: State ketika book overdue (bisa di-return)
- `book.go`: Context yang menyimpan current state dan delegates ke state
- State transitions: Available -> Borrowed -> Overdue -> Available

**Testing Scenario:**
- Borrow available book (sukses, state jadi Borrowed)
- Borrow borrowed book (gagal, state tetap Borrowed)
- Return borrowed book (sukses, state jadi Available)
- Mark available book as overdue (gagal)
- Mark borrowed book as overdue (sukses, state jadi Overdue)
- Return overdue book (sukses, state jadi Available)

---

### 5. STRATEGY PATTERN (Behavioral)

**Alasan Pemilihan:**
Catalog perlu mendukung berbagai metode pencarian: by title, by author, by ISBN, by category. Strategy pattern memungkinkan ganti algoritma pencarian di runtime tanpa memodifikasi client code.

**Class Diagram:**
```
+-------------------+
| SearchStrategy    |
+-------------------+
| + Search()        |
| + GetName()       |
+-------------------+
        ^
        |
        | +----------------+----------------+
        |                 |                |
+----------------+  +----------------+  +----------------+
|TitleSearchStrat|  |AuthorSearchStrat|  |ISBNSearchStrat |
+----------------+  +----------------+  +----------------+
| + Search()      |  | + Search()      |  | + Search()      |
| + GetName()     |  | + GetName()     |  | + GetName()     |
+----------------+  +----------------+  +----------------+


+----------------+         +----------------+
|   Catalog      |         |     Book       |
+----------------+         +----------------+
| - books[]      |<--------| - title        |
| - strategy     |         | - author       |
+----------------+         | - isbn         |
| + SetStrategy()|         | - category     |
| + Find()       |         +----------------+
| + AddBook()    |
| + GetAll()     |
| + GetCount()   |
| + Display()    |
+----------------+
```

**Implementasi:**
- `search_strategy.go`: Interface SearchStrategy dengan Search() method
- `title_search.go`: Strategy untuk search by title (partial match, case-insensitive)
- `author_search.go`: Strategy untuk search by author (partial match, case-insensitive)
- `catalog.go`: Catalog yang menggunakan strategy untuk mencari books
- Strategy bisa diganti di runtime menggunakan SetStrategy()

**Testing Scenario:**
- Add multiple books ke catalog
- Search by title using TitleSearchStrategy
- Switch to AuthorSearchStrategy dan search same query
- Display hasil berbeda berdasarkan strategy
- Switch strategy lagi di runtime
- Verify results sesuai dengan strategy yang aktif

---

## Summary

Sistem Library Management System ini telah mengimplementasikan 5 design pattern:

### Creational Patterns (2)
1. Builder Pattern - Membuat complex Book objects dengan method chaining
2. Prototype Pattern - Cloning book instances untuk multiple copies

### Structural Patterns (1)
3. Decorator Pattern - Menambahkan fitur ke book secara dinamis

### Behavioral Patterns (2)
4. State Pattern - Mengelola state transitions book (Available, Borrowed, Overdue)
5. Strategy Pattern - Multiple search strategies (Title, Author)

Setiap pattern diimplementasikan dengan:
- Clear interface definitions
- Proper Go idioms (implicit interfaces, error handling)
- Comprehensive demo scenarios
- Clean code principles

Semua pattern terintegrasi dalam satu sistem yang fungsional dan dapat di-run dengan `go run main.go`.
