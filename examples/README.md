# Go Concurrency Examples

Kumpulan contoh lengkap tentang concurrency di Go, termasuk goroutine, channel, dan berbagai pattern concurrent programming.

## 📚 Daftar Contoh

### 1. Goroutine
**File:** `01_goroutine/main.go`

Contoh dasar penggunaan goroutine untuk menjalankan fungsi secara concurrent.

```bash
go run examples/01_goroutine/main.go
```

**Yang Dipelajari:**
- Basic goroutine dengan anonymous function
- Goroutine dengan parameter
- Multiple goroutine workers

---

### 2. Channel
**File:** `02_channel/main.go`

Komunikasi antar goroutine menggunakan channel.

```bash
go run examples/02_channel/main.go
```

**Yang Dipelajari:**
- Channel dasar (send & receive)
- Channel dengan loop
- Close channel
- Directional channel (send-only, receive-only)

---

### 3. Producer-Consumer Pattern
**File:** `03_producer_consumer/main.go`

Pattern klasik untuk task distribution dengan multiple workers.

```bash
go run examples/03_producer_consumer/main.go
```

**Yang Dipelajari:**
- Producer menghasilkan jobs
- Multiple consumers memproses jobs secara parallel
- Channel untuk jobs dan results

---

### 4. Buffered vs Unbuffered Channel
**File:** `04_buffered_unbuffered/main.go`

Perbedaan perilaku blocking antara buffered dan unbuffered channel.

```bash
go run examples/04_buffered_unbuffered/main.go
```

**Yang Dipelajari:**
- Unbuffered channel (blocking send/receive)
- Buffered channel (non-blocking sampai buffer penuh)
- Kapan menggunakan masing-masing

---

### 5. Select Statement
**File:** `05_select/main.go`

Menangani multiple channel operations dengan select.

```bash
go run examples/05_select/main.go
```

**Yang Dipelajari:**
- Basic select (multiple channels)
- Select dengan timeout
- Select dengan default (non-blocking)

---

### 6. Worker Pool Pattern
**File:** `06_worker_pool/main.go`

Pattern untuk membatasi jumlah concurrent workers.

```bash
go run examples/06_worker_pool/main.go
```

**Yang Dipelajari:**
- Fixed number of workers
- Job distribution
- Result collection
- Efficient resource management

---

### 7. Fan-Out, Fan-In Pattern
**File:** `07_fanout_fanin/main.go`

Pattern untuk parallel processing dengan multiple workers dan merge results.

```bash
go run examples/07_fanout_fanin/main.go
```

**Yang Dipelajari:**
- Fan-out: distribute work ke multiple workers
- Fan-in: merge results dari multiple channels
- sync.WaitGroup untuk koordinasi

---

## 🚀 Cara Menjalankan

### Jalankan Satu Contoh
```bash
go run examples/01_goroutine/main.go
```

### Jalankan Semua Contoh Sekaligus
```bash
for dir in examples/*/; do
    echo "Running ${dir}..."
    go run "${dir}main.go"
    echo ""
done
```

### Jalankan Dari Root Directory
```bash
go run ./examples/01_goroutine/main.go
go run ./examples/02_channel/main.go
# dst...
```

---

## 📖 Konsep Penting

### Goroutine
- Lightweight thread yang dikelola oleh Go runtime
- Dibuat dengan keyword `go`
- Lebih efisien dari OS thread (ribuan goroutine dalam satu program)

### Channel
- Pipe untuk komunikasi antar goroutine
- **Unbuffered**: Blocking sampai ada receiver
- **Buffered**: Non-blocking sampai buffer penuh
- Tutup channel dengan `close(ch)` setelah selesai

### Pattern Umum

#### Producer-Consumer
```go
jobs := make(chan int)
go producer(jobs)
go consumer(jobs)
```

#### Worker Pool
```go
for w := 1; w <= numWorkers; w++ {
    go worker(w, jobs, results)
}
```

#### Select
```go
select {
case msg := <-ch1:
    // handle ch1
case msg := <-ch2:
    // handle ch2
case <-time.After(1 * time.Second):
    // timeout
}
```

---

## ⚠️ Best Practices

1. **Selalu close channel** setelah selesai mengirim data
2. **Gunakan buffered channel** jika tahu jumlah data yang akan dikirim
3. **Worker pool** untuk membatasi resource usage
4. **Context** untuk cancellation dan timeout (advanced)
5. **sync.WaitGroup** untuk menunggu multiple goroutines selesai
6. **Hindari race condition** dengan proper synchronization

---

## 🔗 Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)

---

## 📝 License

Free to use for learning purposes.
