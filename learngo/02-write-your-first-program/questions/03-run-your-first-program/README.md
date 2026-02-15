🐹 Kuis Dasar Go: Kompilasi vs Runtime
1. Apa perbedaan antara go build dan go run?
   Jawaban: go run mengkompilasi sekaligus menjalankan program; 
   sedangkan go build hanya mengkompilasinya.

Penjelasan: Saat Anda menggunakan go run, Go akan mengkompilasi program Anda, meletakkannya di direktori sementara (temporary directory), lalu menjalankan program yang sudah dikompilasi tersebut dari sana.

2. Di mana Go menyimpan kode yang telah dikompilasi (saat menggunakan go build)?
   Jawaban: Di direktori yang sama tempat Anda menjalankan perintah go build.

Penjelasan Tambahan: * Direktori $GOPATH/src hanya digunakan untuk menyimpan file kode sumber Go.

Direktori $GOPATH/pkg digunakan oleh Go untuk menyimpan hasil kompilasi hanya ketika Anda memanggil perintah go install.

3. Manakah pernyataan yang benar tentang runtime (waktu berjalan)?
   Jawaban: Runtime terjadi ketika program Anda mulai berjalan di komputer.

4. Manakah pernyataan yang benar tentang compile-time (waktu kompilasi)?
   Jawaban: Compile-time terjadi selama program Anda sedng dalam proses kompilasi.

5. Kapan sebuah program Go dapat mencetak pesan ke konsol?
   Jawaban: Saat program sedang berjalan atau berada di fase runtime (setelah compile-time selesai).

Penjelasan: Pada tahap kompilasi (compile-time), program Anda pada dasarnya belum "hidup" sehingga tidak bisa mencetak pesan apa pun. Program baru bisa berinteraksi dengan komputer dan mencetak pesan ke konsol hanya setelah proses kompilasi selesai dan program mulai dijalankan.