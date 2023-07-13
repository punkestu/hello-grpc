# GRPC
## Review
gRPC merupakan salah satu metode fetch and post data yang menggunakan prosedur yang telah disediakan dalam sebuah file kontrak yang disebut proto dimana prosedur dan parameter yang dikirim harus sesuai dengan yang ada dalam file kontrak tersebut.

## Alur Kerja
Pada dasarnya gRPC sangat bergantung pada file kontrak / proto yang disediakan pada bagian backend karena hanya prosedur prosedur tersebutlah yang bisa dipanggil oleh client ketika akses servis menggunakan gRPC. Sehingga alur kerjanya dapat dituliskan sebagai:
1. Server menyediakan dan mempublish service berdasarkan file proto
2. Client mengakses prosedur sesuai dengan yang tertera pada file proto

## Alur Pembuatan
1. Buat file proto. file proto akan berisi:
   1. Service yang akan berisi prosedur yang dapat dipanggil oleh client (diawali dengan keyword rpc)
   2. Message yaitu struktur data yang dapat diterima dan dikirim oleh prosedur pada service (terdiri dari "tipe_data nama = index")
2. Compile file proto sesuai dengan bahasa pemrograman yang digunakan pada backend
3. Buat implementasi dari interface atau kontrak yang dihasilkan dari proses compile sebelumnya
4. Buka network untuk handler gRPC
5. Daftarkan implementasi yang sudah kita buat kedalam register gRPC
6. Ekspos gRPC ke network yang sudah kita buat sebelumnya