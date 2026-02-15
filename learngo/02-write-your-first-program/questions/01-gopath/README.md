🐹 Dasar-Dasar Lingkungan Go (GOPATH)
Berikut adalah panduan tanya-jawab mengenai pengaturan ruang kerja di Bahasa Pemrograman Go.

1. Lokasi Penyimpanan Kode Sumber
   Pertanyaan: Di mana seharusnya Anda menyimpan kode sumber Go?

Jawaban: Di bawah $GOPATH/src

Penjelasan: Secara struktur tradisional (GOPATH mode), direktori src adalah tempat utama untuk menyimpan semua file .go dan paket-paket yang Anda unduh.

2. Definisi $GOPATH
   Pertanyaan: Apa arti dari $GOPATH?

Jawaban: Tempat menyimpan file kode sumber Go dan paket-paket yang telah dikompilasi.

Penjelasan: $GOPATH adalah root dari ruang kerja Anda. Di dalamnya terdapat tiga sub-direktori penting:

src: Untuk kode sumber.

pkg: Untuk objek paket yang terkompilasi.

bin: Untuk file biner (aplikasi jadi) yang bisa langsung dijalankan.

3. Pengaturan $GOPATH
   Pertanyaan: Apakah Anda perlu mengatur $GOPATH secara manual?

Jawaban: Tidak, secara otomatis tersimpan di bawah jalur pengguna (user path).

Penjelasan: Sejak Go versi 1.8, Go sudah memiliki nilai default.

Windows: %USERPROFILE%\go

Unix/Linux/macOS: ~/go

4. Cara Cek Nilai $GOPATH
   Pertanyaan: Bagaimana cara mencetak atau melihat isi $GOPATH?

Jawaban: Menggunakan perintah go env GOPATH.

Penjelasan: Perintah ini akan menampilkan alamat folder yang sedang digunakan oleh sistem Go Anda saat ini