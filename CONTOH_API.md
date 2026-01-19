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
