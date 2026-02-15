## Apa itu scope?
* Blok kode yang dapat dieksekusi
* Visibilitas dari nama-nama yang dideklarasikan **BENAR**
* Menentukan apa yang harus dijalankan
* 
```go
package awesome

import "fmt"

var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```

## Nama manakah di bawah ini yang memiliki scope package?
1. awesome
2. fmt
3. enabled **CORRECT**
4. counter

> **3:** That's right. `enabled` is out of any functions, so it's a package scoped name. `block()` function is also package scoped; it's out of any function too.
>
>


## Nama manakah di bawah ini yang memiliki scope file?
1. awesome
2. fmt **BENAR**
3. enabled
4. block()
5. counter

> **2:** Benar. Nama package yang diimpor memiliki scope file. Nama tersebut hanya dapat digunakan di dalam file yang sama.
>
>


## Nama manakah di bawah ini yang berada dalam scope fungsi block()?
1. awesome
2. fmt
3. enabled
4. block()
5. counter **BENAR**

> **5:** Benar. `counter` dideklarasikan di dalam fungsi `block()`, jadi ia berada dalam scope fungsi block. Di luar fungsi `block()`, kode lain tidak dapat melihatnya.
>
>


## Bisakah `block()` melihat nama `enabled`?
1. Ya: Karena berada dalam scope package **BENAR**
2. Tidak: Karena berada dalam scope file
3. Tidak: Karena berada dalam scope blok dari block()

> **1:** Semua kode di dalam package yang sama dapat melihat semua nama lain yang dideklarasikan di tingkat package.
>
>


## Bisakah file lain dalam package `awesome` melihat nama `counter`?
1. Ya
2. Tidak: Karena berada dalam scope package
3. Tidak: Karena berada dalam scope file
4. Tidak: Karena berada dalam scope blok dari block() **BENAR**

> **4:** Benar. Tidak ada kode lain yang bisa melihat nama-nama di dalam fungsi `block()`. Hanya kode di dalam fungsi `block()` yang bisa melihatnya (Hanya sampai batas tertentu. Contoh: Di dalam blok, kode hanya bisa melihat variabel yang dideklarasikan sebelumnya.)
>
>


## Bisakah file lain dalam package `awesome` melihat nama `fmt`?
1. Ya
2. Tidak: Karena berada dalam scope package
3. Tidak: Karena berada dalam scope file **BENAR**
4. Tidak: Karena berada dalam scope blok dari block()

> **3:** Hanya file yang sama yang dapat melihat package yang diimpor, bukan file lain meskipun mereka berada dalam package yang sama.
>
>


## Apa yang terjadi jika kamu mendeklarasikan nama yang sama dalam scope yang sama seperti ini:
```go
package awesome

import "fmt"

// declared twice in the package scope
var enabled bool
var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```
1. Nama yang baru dideklarasikan akan menimpa yang sebelumnya.
2. Saya tidak bisa melakukannya. Nama tersebut sudah dideklarasikan di scope package. *BENAR*
3. Saya tidak bisa melakukannya. Nama tersebut sudah dideklarasikan di scope file.

> **2:** Benar. Kamu tidak bisa mendeklarasikan nama yang sama dalam scope yang sama. Jika kamu bisa melakukannya, lalu bagaimana kamu akan mengaksesnya kembali? Atau merujuk ke yang mana?
>
>


## Apa yang terjadi jika kamu mendeklarasikan nama yang sama di scope lain seperti ini:
```go
package awesome

import "fmt"

// declared at the package scope
var enabled bool

func block() {
    // also declared in the block scope
    var enabled bool

    var counter int
    fmt.Println(counter)
}
```
1. Nama yang baru dideklarasikan akan menimpa (shadowing) yang sebelumnya. *BENAR*
2. Saya tidak bisa melakukannya. Nama tersebut sudah dideklarasikan di scope package.
3. Saya tidak bisa melakukannya. Nama tersebut sudah dideklarasikan di scope file.

> **1:** Sebenarnya, kamu bisa mendeklarasikan nama yang sama di dalam scope yang lebih dalam (inner scope) seperti ini. Scope milik `block()` berada di dalam package-nya. Ini berarti fungsi tersebut dapat mengakses scope package-nya (tetapi tidak sebaliknya). Jadi, scope `block()` berada di bawah scope package. Ini berarti kamu bisa mendeklarasikan nama yang sama lagi. Nama baru ini akan menimpa (override/shadow) nama yang ada di scope induknya. Keduanya bisa ada secara bersamaan. Silakan cek contoh di repositori kursus untuk mengetahuinya lebih lanjut.
>
>
