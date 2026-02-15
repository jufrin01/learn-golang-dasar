🐹 Kuis Dasar Pemrograman Go (Golang)
1. Kata kunci untuk mendefinisikan package?
   Go
   package main

func main() {
}
Jawaban: package

Penjelasan: Kata kunci package memungkinkan Anda menentukan paket untuk sebuah file Go. func digunakan untuk fungsi, import untuk mengambil paket lain, dan fmt.Println adalah fungsi untuk mencetak teks.

2. Apa tujuan menggunakan package main?
   Jawaban: Untuk membuat program Go yang dapat dieksekusi (executable).

Penjelasan: Tanpa package main, Go akan menganggap kode tersebut sebagai library (pustaka) saja dan tidak bisa dijalankan sebagai aplikasi mandiri.

3. Apa kegunaan dari func main?
   Jawaban: Memungkinkan Go untuk mulai menjalankan program.

Penjelasan: Go secara otomatis mencari dan memanggil fungsi main sebagai titik awal (entry point) eksekusi program.

4. Apa tujuan dari import "fmt"?
   Jawaban: Mengimpor paket fmt agar Anda bisa menggunakan fungsionalitas di dalamnya.

Penjelasan: Setelah diimpor, Anda bisa menggunakan fungsi seperti Println untuk mencetak pesan ke konsol.

5. Kata kunci untuk mendeklarasikan fungsi baru?
   Jawaban: func

Penjelasan: func adalah singkatan dari function. Ini digunakan setiap kali Anda ingin membuat blok kode baru yang bisa dipanggil.

6. Apa itu fungsi?
   Jawaban: Semacam "program mini". Ini adalah blok kode yang dapat digunakan kembali dan dieksekusi.

Penjelasan: Fungsi membungkus logika tertentu sehingga Anda tidak perlu menulis kode yang sama berulang kali.

7. Apakah Anda harus memanggil fungsi main sendiri?
   Jawaban: Tidak, Go memanggil fungsi main secara otomatis.

Penjelasan: Sebagai titik masuk utama, runtime Go sudah diatur untuk mencari func main di dalam package main dan menjalankannya saat program dimulai.

8. Apakah Anda harus memanggil fungsi lain agar bisa dieksekusi?
   (Selain fungsi main)

Jawaban: Ya, Anda harus memanggilnya sendiri agar Go mengeksekusi fungsi tersebut.

Penjelasan: Go tidak akan menjalankan fungsi buatan Anda secara otomatis kecuali Anda memanggilnya di dalam main atau fungsi lain yang sedang berjalan.

9. Apa yang dicetak oleh program berikut?
   Go
   package main

func main() {
}
Jawaban: Program ini benar tetapi tidak mencetak apa pun.

Penjelasan: Program ini valid dan bisa dijalankan, namun karena tidak ada instruksi cetak (seperti fmt.Println), konsol akan kosong.

10. Apakah program ini benar?
    Go
    package main

func main() {
fmt.Println(Hi! I want to be a Gopher!)
}
Jawaban: Program ini Salah/Tidak Benar.

Penjelasan: Ada dua kesalahan utama:

Paket fmt tidak diimpor.

Teks di dalam Println tidak dibungkus dengan tanda kutip ganda ("").

11. Apa hasil cetak program ini?
    Go
    package main
    import "fmt"

func main() {
fmt.Println("Hi there!")
}
Jawaban: Hi there!

Penjelasan: Program ini sudah benar karena mengimpor paket yang diperlukan dan menggunakan sintaks string yang tepat