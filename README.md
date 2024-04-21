# DOKUMENTASI Kuis-3-golang

### Bagian 1
1. path = bangun-datar/segitiga-sama-sisi?hitung=luas&alas=7&tinggi=10
Merupakan path yang digunakan untuk menghitung luas segitiga. Nilai alas dan tinggi dapat disesuaikan.
hitung=luas dapat diganti menjadi hitung=keliling untuk mendapatkan keliling segitiga
   
2. path = bangun-datar/persegi?hitung=luas&sisi=8
Merupakan path yang digunakan untuk menghitung luas persegi. Nilai sisi dapat disesuaikan.
hitung=luas dapat diganti menjadi hitung=keliling untuk mendapatkan keliling persegi.

3. bangun-datar/persegi-panjang?hitung=luas&panjang=8&lebar=10
Merupakan path yang digunakan untuk menghitung luas persegi panjang. Nilai panjang dan lebar dapat disesuaikan.
hitung=luas dapat diganti menjadi hitung=keliling untuk mendapatkan keliling persegi panjang.

7. bangun-datar/lingkaran?hitung=luas&jarijari=9
Merupakan path yang digunakan untuk menghitung luas Lingkaran. Nilai jari-jari dapat disesuaikan.
hitung=luas dapat diganti menjadi hitung=keliling untuk mendapatkan keliling lingkaran.

### Bagian 2
1. path = /categories
Dengan menggunakan method GET maka akan menampilkan kategory yang ada dalam database
Dengan menggunakan method POST dan menambahkan json akan menambahkan data kedalam database. method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password

2. path = /categories/:id
Dengan menggunakan methode PUT dan menyesuaikan nilai id ex:1, 2, 3,.... dsb maka akan mengubah nilai pada database sesuai json yang diinput method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password
Dengan menggunakan methode DELETE dan menyesuaikan nilai id ex:1, 2, 3,.... dsb maka akan menghapus nilai pada database sesuai id yang di input method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password

3. path = /categories:/id/books
dengan menggunakan method GET akan mengembalikan nilai dari database berupa semua buku dengan id kategori yang dipilih

4. path = /books
Dengan menggunakan method GET maka akan menampilkan buku yang ada dalam database.
Dengan menggunakan method POST dan menambahkan json akan menambahkan data buku kedalam database. method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password

5. path = /books/:id
Dengan menggunakan methode PUT dan menyesuaikan nilai id ex:1, 2, 3,.... dsb maka akan mengubah nilai buku pada database sesuai json yang diinput method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password
Dengan menggunakan methode DELETE dan menyesuaikan nilai id ex:1, 2, 3,.... dsb maka akan menghapus nilai buku pada database sesuai id yang di input method ini menggunakan Authentikasi basic auth sehingga harus memasukkan username dan password

### Authorisasi
Terdapat 2 user yang diberikan kuasa dalam mengakses fungsi tertentu.
username = admin, password = password
username = editor, password = secret

### Deploy
Deploy menggunakan Railway yang dapat diakses pada: https://kuis-3-golang-production.up.railway.app

Selamat Mencoba!
