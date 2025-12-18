# 🚢 GoldFile

> *"Hah? Pake library bawaan? Nggak asik! Kita bikin logika sendiri, Trainer!"* — Gold Ship (probably)

[<img width="256" height="256" alt="image" src="https://github.com/user-attachments/assets/1e5a681a-bb7e-422d-bdb1-69271d58e672" />

](https://cdn.cdnstep.com/S2avKgB9SUVfsUQxBqN2/cover.thumb256.webp)
---

## 🧐 Apa ini?!

**GoldFile** adalah File Manager berbasis TUI (Text User Interface) yang dibuat dengan keringat, air mata, dan sedikit *chaos*. 

Apa yang bikin spesial? **KITA GAK PAKE LIBRARY `path/filepath` ATAU `strings` UNTUK LOGIKA UTAMA!** Ya, kamu gak salah baca. Fitur search, navigasi, sampai manipulasi string dibuat secara **MANUAL** (Handmade) menggunakan logika *byte array* murni. Kenapa? Karena biar *hardcore* kayak balapan di medan berlumpur! 🚜💨

## ✨ Fitur Unggulan (The Gold Specs)

* **🕵️‍♂️ Deep Search Mode (Recursive):** Cari file sampai ke akar-akarnya! Mendukung filter ganda.
    Contoh: Ketik `main, .go, import` -> Dia bakal nyari file `main.go` yang isinya ada kata `import`. *Mindblowing!* 🤯
* **🛠 Handmade Logic:**
    Tidak ada `strings.Contains` atau `filepath.WalkDir`. Kita nulis fungsi rekursif sendiri. *Zero bloat, 100% pure Go logic.*
* **🏎 Vim-Style Navigation:**
    Pake `j` dan `k` buat geser. Mouse? Apa itu mouse? Kita keyboard warrior!
* **🖼 Image Preview (Waifu Support):**
    Bisa liat preview gambar langsung di terminal (berkat bantuan `chafa`).
* **🐍 Anti-Lag Technology:**
    Folder berat kayak `node_modules` atau `.git` otomatis di-skip. Kita gak mau lemot!

---

## 🚀 Cara Install (Gampang Banget)

Pastikan kamu sudah install **Go** (Golang) dan terminal yang support warna-warni.

1.  **Clone repo ini (atau copy kodenya):**
    ```bash
    git clone [https://github.com/username-kamu/goldfile.git](https://github.com/username-kamu/goldfile.git)
    cd goldfile
    ```

2.  **Build Binary-nya:**
    ```bash
    go build -o goldfile
    ```

3.  **Jalankan!**
    ```bash/zsh
    ./goldfile
    ```

> **Catatan Penting:** Agar fitur preview gambar jalan, kamu perlu install **`chafa`** di Linux/Mac kamu (`sudo apt install chafa` atau `brew install chafa`). Kalau gak ada, ya gak muncul gambarnya.

---

## 🎮 Kontrol (Cheatsheet)

| Tombol | Fungsi | Komentar Gold Ship |
| :--- | :--- | :--- |
| `j` / `↓` | Turun | Gas pol ke bawah! |
| `k` / `↑` | Naik | Rem dikit, naik lagi! |
| `/` | **SEARCH MODE** | *Serious Mode On* 🧐 |
| `Enter` | Buka Folder/File | Masuk, Pak Eko! |
| `Backspace` | Kembali (Parent Dir) | Mundur alon-alon. |
| `Tab` | Preview File | Intip isinya dikit. |
| `q` | Keluar | Bye bye, Trainer! 👋 |

---

## 🔍 Cara Pakai Search (Rahasia Dapur)

Saat kamu tekan `/`, bar kuning akan muncul di bawah.

* **Cari Nama:** Ketik `json` -> Muncul semua file `.json`.
* **Cari Isi File:** Ketik nama variabel codinganmu, misal `StateAplikasi`. Dia bakal nyari file yang *isinya* ada kata itu.
* **Combo Maut (Koma):** Ketik: `utils, .go, func`
    Artinya: Cari file yang path-nya ada "utils", ekstensinya ".go", DAN isinya ada kata "func".

---

## 📜 Lisensi

Project ini dilindungi oleh **MIT License**.
Artinya: Bebas kamu pake, kamu modif, kamu jual, atau kamu jadiin bungkus gorengan. Yang penting jangan lupa kredit ke pembuat aslinya (dan Gold Ship).

```text
MIT License
Copyright (c) 2025 [Kitsunesan13]
