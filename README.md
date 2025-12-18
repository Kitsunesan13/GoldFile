# 🚢 GoldFile 🐴

> *"Library UI? Siapa butuh Library UI?! Kita rakit engine sendiri, Trainer!"* — Gold Ship

![Gold Ship Joget](https://cdn.cdnstep.com/S2avKgB9SUVfsUQxBqN2/cover.thumb256.webp)
---

## 🧐 Apa ini?!

**GoldFile** adalah File Manager berbasis TUI (Text User Interface) yang super ringan.

Apa yang bikin spesial? **KITA GAK PAKE LIBRARY UI APAPUN!**
Biasanya orang bikin TUI pake `bubbletea` atau `tview` yang beratnya minta ampun. GoldFile? **Nggak.** Kita nulis *Render Engine* sendiri pake **ANSI Escape Codes** murni. Hasilnya? Aplikasi yang ngebut kayak Gold Ship pas lagi *mood* lari! 🚀💨

## ✨ Fitur Unggulan (The Gold Specs)

* **🧠 Smart Recursive Search:**
    Bukan cuma cari nama file, tapi bisa cari sampai ke **Isi File-nya** (Content Grep) dan masuk ke sub-folder dalam-dalam!
* **🔍 Multi-Filter Logic:**
    Pake logika koma (`,`) buat filter berlapis.
    Contoh: `main, .go, import` 👉 Cari file `main.go` yang di dalamnya ada kata `import`. *Cerdas kan?* 🧐
* **⚡ Zero Bloat UI:**
    Tampilan antarmuka dibuat manual pixel-per-pixel (eh, char-per-char) di terminal. Tanpa dependensi aneh-aneh.
* **🛡️ Anti-Lag Protocol:**
    Pencarian otomatis menghindari "Folder Terkutuk" kayak `node_modules`, `.git`, atau `vendor`. Laptopmu gak bakal *hang*!
* **🖼 Image Preview (Waifu Support):**
    Preview gambar langsung di terminal via `chafa`.
* **🏎 Vim-Style Navigation:**
    Navigasi sat-set pake `j` dan `k`.

---

## 🚀 Cara Install

Pastikan kamu sudah install **Go** (Golang).

1.  **Clone & Masuk Folder:**
    ```bash
    git clone [https://github.com/username-kamu/goldfile.git](https://github.com/username-kamu/goldfile.git)
    cd goldfile
    ```

2.  **Build Binary:**
    ```bash
    go build -o goldfile
    ```

3.  **Jalankan:**
    ```bash
    ./goldfile
    ```

> **Catatan:** Install `chafa` dulu kalau mau liat gambar (`sudo apt install chafa`).

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

1.  **Cari Nama:** Ketik `json` -> Muncul semua file `.json`.
2.  **Cari Isi File:** Lupa nama file tapi ingat isinya? Ketik aja variabelnya, misal `StateAplikasi`. Ketemu deh!
3.  **Combo Maut:** Ketik `utils, .go` -> Cari semua file `.go` di dalam folder/subfolder `utils`.

---

## 📜 Lisensi

Project ini dilindungi oleh **MIT License**.
Artinya: Bebas kamu pake, kamu modif, kamu jual, atau kamu jadiin bungkus gorengan. Yang penting jangan lupa kredit ke pembuat aslinya (dan Gold Ship).

```text
MIT License
Copyright (c) 2025 [Nama Kamu / Trainer]
