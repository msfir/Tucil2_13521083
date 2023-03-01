# Tugas Kecil 2 Strategi Algoritma: Mencari Pasangan Titik Terdekat dengan Algoritma Divide and Conquer

## Penjelasan Singkat
Program ini adalah program yang mengimplementasikan algoritma Divide and Conquer untuk mencari pasangan titik di ruang 3D atau ruang vektor $\mathbf{R}^N, N \ge 1$ (Bonus 2), yang memiliki jarak Euclidean terdekat. Program ini juga mencari pasangan titik terdekat tersebut dengan algoritma Brute Force untuk perbandingan performa. Selain itu, program ini dapat memvisualisasikan titik-titik pada himpunan di bidang 2 dimensi dan ruang 3 dimensi dengan menggunakan program gnuplot.

## Requirements
Berikut beberapa hal yang harus diinstall di komputer Anda:
1. [Bahasa pemrograman Go](https://go.dev/dl/)
2. [Gnuplot](http://www.gnuplot.info/) untuk visualisasi

## Langkah Kompilasi
```shell
$ git clone https://github.com/msfir/Tucil2_13521083.git
$ cd Tucil2_13521083
$ make        # untuk Linux
$ .\build.bat # untuk Windows
```
Setelah itu executable file akan berada di dalam folder bin.

## Cara Menggunakan Program
1. Jalankan program
2. Program menanyakan dimensi vektor yang diinginkan (batasan: N >= 1)
3. Program akan menanyakan jumlah titik yang akan digenerate (batasan: n >= 2)
4. Selanjutnya, program akan menjalankan algoritma dan menampilkan pasangan titik terdekat serta hasil pengukuran performanya
5. Jika program **gnuplot** ada di dalam PATH, maka program akan memvisualisasikan titik-titik yang telah digenerate dan menampilkan pasangan titik terdekat dengan warna merah

Authored by: Moch. Sofyan Firdaus (13521083)
