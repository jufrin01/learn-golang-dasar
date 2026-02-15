## Di mana sebaiknya menyimpan file source code yang termasuk dalam sebuah package?
1. Setiap file harus dimasukkan ke dalam direktori yang berbeda
2. Dalam satu direktori tunggal *BENAR*



## Mengapa klausa package digunakan dalam file source-code Go?
1. Digunakan untuk mengimpor sebuah package
2. Digunakan untuk memberi tahu Go bahwa file tersebut adalah bagian dari sebuah package *BENAR*
3. Digunakan untuk mendeklarasikan fungsi baru

> **1:** Pernyataan `import` yang melakukan hal tersebut.
>
>
> **3:** Pernyataan `func` yang melakukan hal tersebut.
>
>


## Di mana sebaiknya kamu meletakkan `package clause` dalam file source-code Go?
1. Sebagai kode pertama dalam file source code Go *BENAR*
2. Sebagai kode terakhir dalam file source code Go
3. Kamu bisa meletakkannya di mana saja


## Berapa kali kamu bisa menggunakan `package clause` untuk satu file source code?
1. Satu kali *BENAR*
2. Tidak sama sekali
3. Beberapa kali


## Manakah penggunaan `package clause` yang benar?
1. `my package`
2. `package main`
3. `pkg main`


## Manakah yang benar?
1. Semua file yang berada dalam package yang sama tidak dapat memanggil fungsi satu sama lain
2. Semua file yang berada dalam package yang sama dapat memanggil fungsi satu sama lain *BENAR*


## Bagaimana cara menjalankan beberapa file Go?
1. go run *.*go
2. go build *go
3. go run go
4. go run *.go *BENAR*

> **4:** Kamu juga bisa memanggilnya seperti ini (asumsikan ada file1.go, file2.go, dan file3.go dalam direktori yang sama): go run file1.go file2.go file3.go
>
>