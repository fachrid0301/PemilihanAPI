# Contoh Penggunaan API Kandidat

## Base URL
```
http://localhost:8080
```

---

## 1. POST /kandidat - Membuat Kandidat Baru

### Request
**URL:** `POST http://localhost:8080/kandidat`

**Headers:**
```
Content-Type: application/json
```

**Body (JSON):**
```json
{
  "nomor_urut": 1,
  "nama_ketua": "Ahmad Hidayat",
  "nama_wakil": "Siti Nurhaliza",
  "visi": "Mewujudkan kampus yang maju, inovatif, dan berkarakter",
  "misi": "1. Meningkatkan kualitas pendidikan\n2. Mengembangkan teknologi informasi\n3. Membangun karakter mahasiswa",
  "foto": "https://example.com/foto/kandidat1.jpg"
}
```

**Body (Form Data):**
```
nomor_urut=1
nama_ketua=Ahmad Hidayat
nama_wakil=Siti Nurhaliza
visi=Mewujudkan kampus yang maju, inovatif, dan berkarakter
misi=1. Meningkatkan kualitas pendidikan\n2. Mengembangkan teknologi informasi\n3. Membangun karakter mahasiswa
foto=https://example.com/foto/kandidat1.jpg
```

**Body (tanpa foto - opsional):**
```json
{
  "nomor_urut": 2,
  "nama_ketua": "Budi Santoso",
  "nama_wakil": "Dewi Lestari",
  "visi": "Kampus unggul dan berdaya saing global",
  "misi": "1. Peningkatan kualitas SDM\n2. Pengembangan riset\n3. Kerjasama internasional"
}
```

### Response Success (201 Created)
```json
{
  "message": "Kandidat berhasil dibuat",
  "data": {
    "id_kandidat": 1,
    "nomor_urut": 1,
    "nama_ketua": "Ahmad Hidayat",
    "nama_wakil": "Siti Nurhaliza",
    "visi": "Mewujudkan kampus yang maju, inovatif, dan berkarakter",
    "misi": "1. Meningkatkan kualitas pendidikan\n2. Mengembangkan teknologi informasi\n3. Membangun karakter mahasiswa",
    "foto": "https://example.com/foto/kandidat1.jpg",
    "created_at": "2026-01-19 09:44:15"
  }
}
```

### Response Error (400 Bad Request)
```json
{
  "message": "nomor urut sudah digunakan"
}
```

atau

```json
{
  "message": "Data tidak valid. Nomor urut, nama ketua, nama wakil, visi, dan misi wajib diisi"
}
```

---

## 2. PUT /kandidat/:id - Update Kandidat

### Request
**URL:** `PUT http://localhost:8080/kandidat/1`

**Headers:**
```
Content-Type: application/json
```

**Body (Update semua field):**
```json
{
  "nomor_urut": 1,
  "nama_ketua": "Ahmad Hidayat Updated",
  "nama_wakil": "Siti Nurhaliza Updated",
  "visi": "Visi yang diperbarui",
  "misi": "Misi yang diperbarui",
  "foto": "https://example.com/foto/kandidat1-updated.jpg"
}
```

**Body (Update sebagian field - partial update):**
```json
{
  "nama_ketua": "Ahmad Hidayat Baru",
  "visi": "Visi baru yang lebih baik"
}
```

**Body (Hapus foto - set foto menjadi NULL):**
```json
{
  "foto": ""
}
```

**Body (Update tanpa mengubah foto - tidak kirim field foto):**
```json
{
  "nama_ketua": "Nama Baru",
  "nama_wakil": "Nama Wakil Baru"
}
```

### Response Success (200 OK)
```json
{
  "message": "Kandidat berhasil diupdate",
  "data": {
    "id_kandidat": 1,
    "nomor_urut": 1,
    "nama_ketua": "Ahmad Hidayat Updated",
    "nama_wakil": "Siti Nurhaliza Updated",
    "visi": "Visi yang diperbarui",
    "misi": "Misi yang diperbarui",
    "foto": "https://example.com/foto/kandidat1-updated.jpg",
    "created_at": "2026-01-19 09:44:15"
  }
}
```

### Response Error (400 Bad Request)
```json
{
  "message": "nomor urut sudah digunakan"
}
```

atau

```json
{
  "message": "tidak ada data yang diupdate"
}
```

### Response Error (404 Not Found)
```json
{
  "message": "kandidat tidak ditemukan"
}
```

---

## 3. DELETE /kandidat/:id - Hapus Kandidat

### Request
**URL:** `DELETE http://localhost:8080/kandidat/1`

**Headers:**
```
(tidak perlu body)
```

### Response Success (200 OK)
```json
{
  "message": "Kandidat berhasil dihapus"
}
```

### Response Error (400 Bad Request)
```json
{
  "message": "kandidat tidak ditemukan"
}
```

---

## 4. GET /kandidat - Get All Kandidat

### Request
**URL:** `GET http://localhost:8080/kandidat`

### Response Success (200 OK)
```json
{
  "message": "Berhasil mengambil data kandidat",
  "data": [
    {
      "id_kandidat": 1,
      "nomor_urut": 1,
      "nama_ketua": "Ahmad Hidayat",
      "nama_wakil": "Siti Nurhaliza",
      "visi": "Mewujudkan kampus yang maju",
      "misi": "Misi kandidat 1",
      "foto": "https://example.com/foto/kandidat1.jpg",
      "created_at": "2026-01-19 09:44:15"
    },
    {
      "id_kandidat": 2,
      "nomor_urut": 2,
      "nama_ketua": "Budi Santoso",
      "nama_wakil": "Dewi Lestari",
      "visi": "Kampus unggul",
      "misi": "Misi kandidat 2",
      "foto": "",
      "created_at": "2026-01-19 09:45:20"
    }
  ]
}
```

---

## 5. GET /kandidat/:id - Get Kandidat by ID

### Request
**URL:** `GET http://localhost:8080/kandidat/1`

### Response Success (200 OK)
```json
{
  "message": "Berhasil mengambil data kandidat",
  "data": {
    "id_kandidat": 1,
    "nomor_urut": 1,
    "nama_ketua": "Ahmad Hidayat",
    "nama_wakil": "Siti Nurhaliza",
    "visi": "Mewujudkan kampus yang maju",
    "misi": "Misi kandidat 1",
    "foto": "https://example.com/foto/kandidat1.jpg",
    "created_at": "2026-01-19 09:44:15"
  }
}
```

### Response Error (404 Not Found)
```json
{
  "message": "kandidat tidak ditemukan"
}
```

---

## Contoh Menggunakan cURL

### POST - Create Kandidat
```bash
curl -X POST http://localhost:8080/kandidat \
  -H "Content-Type: application/json" \
  -d '{
    "nomor_urut": 1,
    "nama_ketua": "Ahmad Hidayat",
    "nama_wakil": "Siti Nurhaliza",
    "visi": "Mewujudkan kampus yang maju",
    "misi": "Misi kandidat 1",
    "foto": "https://example.com/foto/kandidat1.jpg"
  }'
```

### PUT - Update Kandidat
```bash
curl -X PUT http://localhost:8080/kandidat/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nama_ketua": "Ahmad Hidayat Updated",
    "visi": "Visi yang diperbarui"
  }'
```

### DELETE - Hapus Kandidat
```bash
curl -X DELETE http://localhost:8080/kandidat/1
```

### GET - Get All Kandidat
```bash
curl -X GET http://localhost:8080/kandidat
```

### GET - Get Kandidat by ID
```bash
curl -X GET http://localhost:8080/kandidat/1
```

---

## Contoh Menggunakan JavaScript (Fetch API)

### POST - Create Kandidat
```javascript
fetch('http://localhost:8080/kandidat', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    nomor_urut: 1,
    nama_ketua: 'Ahmad Hidayat',
    nama_wakil: 'Siti Nurhaliza',
    visi: 'Mewujudkan kampus yang maju',
    misi: 'Misi kandidat 1',
    foto: 'https://example.com/foto/kandidat1.jpg'
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

### PUT - Update Kandidat
```javascript
fetch('http://localhost:8080/kandidat/1', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    nama_ketua: 'Ahmad Hidayat Updated',
    visi: 'Visi yang diperbarui'
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

### DELETE - Hapus Kandidat
```javascript
fetch('http://localhost:8080/kandidat/1', {
  method: 'DELETE'
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## Catatan Penting

1. **Nomor Urut**: Harus unik, tidak boleh duplikat
2. **Foto**: Opsional, bisa NULL atau string kosong
3. **Update Partial**: Bisa update hanya beberapa field saja
4. **Hapus Foto**: Kirim `"foto": ""` untuk menghapus foto (set menjadi NULL)
5. **ID**: `id_kandidat` adalah auto-increment, tidak perlu dikirim saat create

---

## 6. POST /voting - User Melakukan Voting

### Request
**URL:** `POST http://localhost:8080/voting`

**Headers:**
```
Content-Type: application/json
```

**Body (JSON):**
```json
{
  "id_user": 3,
  "id_kandidat": 2
}
```

### Response Success (201 Created)
```json
{
  "message": "Voting berhasil",
  "data": {
    "id_voting": 1,
    "id_user": 3,
    "id_kandidat": 2,
    "waktu_voting": "2026-01-20 10:15:00"
  }
}
```

### Response Error (400 Bad Request)

**User sudah voting:**
```json
{
  "message": "user sudah melakukan voting"
}
```

**User tidak aktif:**
```json
{
  "message": "user tidak aktif"
}
```

**Hanya user yang bisa voting (admin tidak bisa):**
```json
{
  "message": "hanya user yang dapat melakukan voting"
}
```

**Kandidat tidak ditemukan:**
```json
{
  "message": "kandidat tidak ditemukan"
}
```

**ID tidak valid:**
```json
{
  "message": "id_user wajib diisi dan harus berupa angka yang valid"
}
```

### Contoh Menggunakan cURL
```bash
curl -X POST http://localhost:8080/voting \
  -H "Content-Type: application/json" \
  -d "{\"id_user\":3,\"id_kandidat\":2}"
```

### Contoh Menggunakan JavaScript (Fetch API)
```javascript
fetch('http://localhost:8080/voting', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    id_user: 3,
    id_kandidat: 2
  })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
```

---

## 7. GET /voting - Get All Voting (Admin)

### Request
**URL:** `GET http://localhost:8080/voting`

### Response Success (200 OK)
```json
{
  "message": "Berhasil mengambil data voting",
  "data": [
    {
      "id_voting": 1,
      "id_user": 3,
      "id_kandidat": 2,
      "waktu_voting": "2026-01-20 10:15:00"
    },
    {
      "id_voting": 2,
      "id_user": 4,
      "id_kandidat": 1,
      "waktu_voting": "2026-01-20 10:17:22"
    }
  ]
}
```

---

## 8. GET /voting/result - Get Hasil Voting

### Request
**URL:** `GET http://localhost:8080/voting/result`

### Response Success (200 OK)
```json
{
  "message": "Berhasil mengambil hasil voting",
  "data": [
    {
      "id_kandidat": 1,
      "nomor_urut": 1,
      "nama_ketua": "Calon 1",
      "nama_wakil": "Wakil 1",
      "jumlah_suara": 10
    },
    {
      "id_kandidat": 2,
      "nomor_urut": 2,
      "nama_ketua": "Calon 2",
      "nama_wakil": "Wakil 2",
      "jumlah_suara": 7
    }
  ]
}
```

---

## Cara Test Voting Lengkap

### Langkah 1: Pastikan User Siap
Di database, pastikan ada user dengan:
- `role = 'user'` (bukan admin)
- `status = 'aktif'` (bukan sudah_voting atau nonaktif)

Contoh query:
```sql
SELECT id, username, role, status FROM users WHERE role = 'user' AND status = 'aktif';
```

### Langkah 2: Pastikan Ada Kandidat
```sql
SELECT id_kandidat, nomor_urut, nama_ketua FROM kandidat;
```

### Langkah 3: Test Voting
**Request:**
```json
POST http://localhost:8080/voting
Content-Type: application/json

{
  "id_user": 3,
  "id_kandidat": 2
}
```

**Jika berhasil:**
- Response: `"Voting berhasil"`
- Di database: 
  - Tabel `voting` ada record baru
  - Tabel `users` → `status` user berubah jadi `'sudah_voting'`

**Jika error 400:**
- Cek log di terminal server untuk detail error
- Pastikan `id_user` dan `id_kandidat` adalah angka valid
- Pastikan user `role = 'user'` dan `status = 'aktif'`
- Pastikan user belum pernah voting sebelumnya

### Langkah 4: Test Voting Kedua (Harus Gagal)
Coba voting lagi dengan user yang sama → harus dapat error `"user sudah melakukan voting"`

### Langkah 5: Cek Hasil Voting
```bash
GET http://localhost:8080/voting/result
```

---

## Troubleshooting Error 400 Bad Request

1. **Pastikan format JSON benar:**
   ```json
   {
     "id_user": 3,        // angka, bukan string
     "id_kandidat": 2    // angka, bukan string
   }
   ```

2. **Pastikan Content-Type header:**
   ```
   Content-Type: application/json
   ```

3. **Cek log server** untuk melihat error detail:
   - Buka terminal tempat server berjalan
   - Lihat log yang muncul saat request voting

4. **Pastikan user ada dan valid:**
   - User harus ada di database
   - `role` harus `'user'` (bukan `'admin'`)
   - `status` harus `'aktif'` (bukan `'sudah_voting'` atau `'nonaktif'`)

5. **Pastikan kandidat ada:**
   - `id_kandidat` harus ada di tabel `kandidat`
