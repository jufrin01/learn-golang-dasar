1. **Print your name and your best friend's name** using Println twice. [Check out this exercise here](https://github.com/inancgumus/learngo/tree/master/02-write-your-first-program/exercises/01-print-names).

2. **Say hello to yourself.**

    1. Build your program using `go build`

    2. **Send it to your friend**

       S/he should use be using the same operating system.

       For example, if you're using Windows, then hers/his should be Windows as well.

    3. **Send your program to a friend with a different operating system**

       So, you should compile y🚀 Latihan Dasar Go (Langkah demi Langkah)
1. Mencetak Nama (Menggunakan Println)
   Buat file bernama main.go dan masukkan kode berikut:

Go
package main

import "fmt"

func main() {
fmt.Println("Nama Saya: Gopher")
fmt.Println("Nama Teman: Budi")
}

2. Menyapa Diri Sendiri & Kompilasi Silang (Cross-Compilation)
   Ubah kode di atas untuk menyapa diri sendiri, lalu ikuti langkah ini:

A. Build untuk OS Sendiri:
Jalankan di terminal:
go build -o sapa_diriku (Ini akan menghasilkan file biner/exe).

B. Kompilasi untuk OS Berbeda:
Salah satu keunggulan Go adalah kemampuannya membuat file aplikasi untuk OS lain tanpa pindah komputer.

Untuk Mac (OSX): GOOS=darwin GOARCH=amd64 go build -o sapa_mac

Untuk Windows: GOOS=windows GOARCH=amd64 go build -o sapa_windows.exe

Untuk Linux: GOOS=linux GOARCH=amd64 go build -o sapa_linux

3. Perbedaan Print vs Println
   Coba ubah kodenya menjadi:

Go
fmt.Print("Halo")
fmt.Print("Dunia")
Hasilnya: Teks akan menyambung (HaloDunia) karena Print tidak menambahkan baris baru (newline) di akhir, sedangkan Println menambahkannya otomatis.

4. Menggunakan Banyak Nilai (Koma)
   Anda bisa mencetak banyak hal sekaligus:

Go
fmt.Println("Umur saya:", 25, "tahun.")
Go akan otomatis memberikan spasi di antara nilai-nilai tersebut.

5. Eksperimen: Menghapus Tanda Kutip
   Jika Anda menulis fmt.Println(Halo), Go akan menganggap Halo sebagai variabel, bukan teks (string). Karena variabel Halo belum didefinisikan, program akan Error saat dijalankan (undefined: Halo).

6. Eksperimen: Memindah package dan import ke Bawah
   Jika Anda memindahkan package main ke baris paling bawah:
   Hasilnya: Error. Di Go, package harus selalu menjadi baris kode pertama yang bukan komentar. Ini adalah aturan wajib.

7. Membaca Dokumentasi (pkg.go.dev)
   Membaca dokumentasi adalah skill terpenting seorang programmer Go.

fmt: Untuk urusan input/output teks.

os: Untuk interaksi dengan sistem operasi (file, argumen terminal).

math: Untuk fungsi matematika tingkat lanjut.

8. Ikuti Go Tour
   Sangat disarankan untuk meluangkan waktu 15-30 menit di tour.golang.org. Ini adalah taman bermain interaktif di mana Anda bisa mengetik kode Go langsung di browser.our program for her operating system.

       **Create an OSX executable:**
       `GOOS=darwin GOARCH=386 go build`

       **Create a Windows executable:**
       `GOOS=windows GOARCH=386 go build`

       **Create a Linux executable:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **You can find the full list in here:**
       https://golang.org/doc/install/source#environment

       **NOTE:** If you're using the command prompt or the PowerShell, you may need to use it like this:
       `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

3. **Call [Print](https://golang.org/pkg/fmt/#Print) instead of [Println](https://golang.org/pkg/fmt/#Println)** to see what happens.

4. **Call [Println](https://golang.org/pkg/fmt/#Println) or [Print](https://golang.org/pkg/fmt/#Print) with multiple values** by separating them using commas.

5. **Remove the double quotes from a string literal** and see what happens.

6. **Move the package and import statement** to the bottom of the file and see what happens.

7. **[Read Go online documentation](https://golang.org/pkg)**.

    1. Take a quick look at the packages and read what they do.

    2. Look at their source-code by clicking on their titles.

    3. You don't have to understand everything, just do it. This will warm you up for the upcoming lectures.

8. Also, **take a tour on**: https://tour.golang.org/

    1. Have a quick look. Check out the language features.

    2. We're going to talk all about them soon.
