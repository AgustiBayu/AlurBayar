# 🚀 AlurBayar - Go Midtrans x FakeStoreAPI Integration

Sistem backend **Go (Golang)** yang mengintegrasikan Payment Gateway Midtrans dengan katalog produk dinamis dari **FakeStoreAPI**. Proyek ini mendemonstrasikan alur transaksi lengkap dari pengambilan produk pihak ketiga hingga otomatisasi pembayaran menggunakan **Redis** dan **PostgreSQL**.
![GitHub Logo](https://cdn.prod.website-files.com/6100d0111a4ed76bc1b9fd54/62217e885f52b860da9f00cc_Apa%20Itu%20Golang%3F%20Apa%20Saja%20Fungsi%20Dan%20Keunggulannya%20-%20Binar%20Academy.jpeg)

## ✨ Fitur Utama
* **Dynamic Product Sourcing**: Mengambil data katalog real-time dari fakestoreapi.com.
* **Midtrans Snap Integration**: Pembuatan token transaksi otomatis untuk checkout yang aman.
* **Asynchronous Processing**: Menggunakan Redis sebagai menyimpan data produk dari FakeStoreAPI selama 24 jam untuk mengurangi beban network.
* **Automated Webhook:**: Menggunakan Localtunnel dalam Docker untuk menerima notifikasi pembayaran dari Midtrans ke localhost secara otomatis.
* **Database Persistence**: Menggunakan PostgreSQL untuk menyimpan data transaksi.
* **Dockerized**: Siap dijalankan di server mana pun hanya dengan satu perintah.

## 🛠️ Teknologi yang Digunakan
* **Language**: Go 1.24
* **Framework**: Gin Gonic
* **Cache & Queue**: Redis 7
* **Database**: PostgreSQL 14
* **Container**: Docker & Docker Compose

## 🚀 Cara Menjalankan (Quick Start)

Pastikan kamu sudah menginstall **Docker** dan **Docker Compose** di laptopmu.

1. **Clone Repositori**
   ```bash
   git clone [https://github.com/AgustiBayu/AlurBayar.git]
   cd AlurBayar
2. **Jalankan dengan Docker**
Cukup jalankan perintah berikut, Docker akan mengatur Database PostgreSQL dan Backend secara otomatis:
   ```bash
   docker-compose up --build
3. **Setup Webhook Midtrans**
Cek log container alurbayar-tunnel untuk mendapatkan URL publik (contoh: https://agusti-payment-dev.loca.lt). Masukkan URL tersebut di Dashboard Midtrans (Settings > Payment > HTTP Notification).
4. **Akses Aplikasi** akan berjalan pada: http://localhost:8080

## 🧪 Cara Pengujian (Testing)
1. **Mengambil Produk (Uji Coba Redis Cache)**
Mengambil data produk dari API luar. Hit pertama akan lambat (fetch API), hit kedua akan sangat cepat (fetch Redis).
* **Method:** GET
* **Url:**
   ``` bash 
   http://localhost:8080/api/v1/products/1
* **Json Body:**
   ``` bash 
   {
    "data": {
        "ID": "",
        "ProductID": ,
        "ProductName": "",
        "Amount": ,
        "Status": "PENDING",
        "SnapToken": "",
        "CreatedAt": ""
      },
    "message": "",
    "payment_url": ""
   }
2. **Membuat Pesanan (Checkout)**
Membuat transaksi baru dan mendapatkan snap_token untuk pembayaran.
* **Method:** POST
* **Url:**
   ``` bash 
   http://localhost:8080/api/v1/checkout
* **Json Body:**
   ``` bash 
   {
   "product_id": 1
   }
* **Simulasi Pembayaran:**
   1. Copy snap_token yang didapat dari response checkout.
   2. Buka Midtrans Simulator.
   3. Bayar menggunakan simulator, status di database PostgreSQL akan otomatis berubah menjadi settlement melalui webhook.

