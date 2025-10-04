# go-hack

Repositori untuk belajar Go (Golang) dengan fokus pada concurrency patterns dan praktik terbaik.

## 📖 Isi Repository

### Concurrency Examples

Kumpulan contoh lengkap tentang concurrency di Go, termasuk goroutine, channel, dan berbagai pattern concurrent programming.

**Lokasi:** `examples/`

#### Daftar Contoh:

1. **Goroutine** (`01_goroutine/`) - Basic goroutine dan concurrent execution
2. **Channel** (`02_channel/`) - Komunikasi antar goroutine
3. **Producer-Consumer** (`03_producer_consumer/`) - Pattern untuk task distribution
4. **Buffered vs Unbuffered** (`04_buffered_unbuffered/`) - Perbedaan blocking behavior
5. **Select** (`05_select/`) - Multiple channel handling dengan timeout
6. **Worker Pool** (`06_worker_pool/`) - Fixed worker pool pattern
7. **Fan-Out/Fan-In** (`07_fanout_fanin/`) - Parallel processing pattern

## 🚀 Cara Menggunakan

### Clone Repository
```bash
git clone https://github.com/zikazama/go-hack.git
cd go-hack
```

### Jalankan Contoh
```bash
# Goroutine
go run examples/01_goroutine/main.go

# Channel
go run examples/02_channel/main.go

# Producer-Consumer
go run examples/03_producer_consumer/main.go

# Buffered vs Unbuffered
go run examples/04_buffered_unbuffered/main.go

# Select
go run examples/05_select/main.go

# Worker Pool
go run examples/06_worker_pool/main.go

# Fan-Out/Fan-In
go run examples/07_fanout_fanin/main.go
```

### Jalankan Semua Contoh
```bash
for dir in examples/*/; do
    echo "Running ${dir}..."
    go run "${dir}main.go"
    echo ""
done
```

## 📚 Konsep yang Dipelajari

### Goroutine
- Lightweight concurrent threads
- Lebih efisien dari OS threads
- Dikelola oleh Go runtime

### Channel
- **Unbuffered Channel**: Blocking sampai ada receiver
- **Buffered Channel**: Non-blocking sampai buffer penuh
- Komunikasi type-safe antar goroutine

### Concurrency Patterns
- **Producer-Consumer**: Pemisahan produksi dan konsumsi data
- **Worker Pool**: Membatasi jumlah concurrent workers
- **Fan-Out/Fan-In**: Distribute work dan merge results
- **Select**: Handling multiple channel operations

## 🛠️ Requirements

- Go 1.16 atau lebih baru

## 📖 Resources

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Go by Example](https://gobyexample.com/)

## 📝 License

Free to use for learning purposes.

## 👤 Author

**Fauzi**
- GitHub: [@zikazama](https://github.com/zikazama)
