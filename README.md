# AnekaZoo API

## Prerequisites
- Go
- MySQL 

## How to run

1. Clone the repository
git clone https://github.com/username/anekazoo-api.git
cd anekazoo-api

2. Buat database baru di MySQL untuk digunakan oleh API:
mysql -u root -p -e "CREATE DATABASE anekazoo_db;"
Gantilah root dengan user MySQL Anda jika berbeda.

3. Impor file SQL ke dalam database yang baru dibuat
mysql -u root -p anekazoo_db < anekazoo-api/database/anekazoo_db.sql

4. Akses API di `localhost:8080/v1/animals`

## API Endpoints
- `GET /v1/animals` - Mendapatkan daftar hewan
- `GET /v1/animals/{id}` - Mendapatkan detail hewan berdasarkan ID
- `POST /v1/animals` - Menambahkan hewan baru
- `PUT /v1/animals` - Mengupdate hewan
- `DELETE /v1/animals/{id}` - Menghapus hewan
