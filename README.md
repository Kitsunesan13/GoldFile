# GoldFile
🛸 GOLD FILE: The "Goru-shi" Edition 🐴

(Bayangkan gambar ini bergerak di terminalmu!)

    "Oi, Trainer! Kenapa pakai File Manager yang membosankan kalau kamu bisa pakai yang EMAS? Ini bukan sekadar alat, ini adalah KARYA SENI yang dibuat dengan keringat, air mata, dan sedikit Yakisoba!" — Gold Ship

🧐 Apa ini?!

GoldFile adalah File Manager berbasis TUI (Text User Interface) yang ditulis dengan bahasa Go.

Tapi tunggu dulu... ini bukan File Manager biasa. Ini adalah "The Hardcore Edition". Kenapa? Karena kami menolak menggunakan library bawaan yang memanjakan seperti strings atau filepath. Kami menulis ulang logika pemecah string, pencarian path, dan rekursi DARI NOL (FROM SCRATCH).

Kenapa? KARENA KITA BISA! 😤
✨ Fitur Unggulan (Yang Bikin Tetangga Iri)

    ⚡ Super Cepat & Ringan: Navigasi secepat Gold Ship lari di trek rumput (kalau lagi mood).

    🔍 Deep Search System (Manual Logic):

        Cari file sampai ke lubang semut (Sub-folder).

        Bisa filter pakai Koma Logic (Contoh: main, .go, import).

        Membaca isi file? BISA! (Kami baca 10KB pertama biar gak meledak).

    🚫 No Helper Libraries: strings.Split? Gak kenal. filepath.Walk? Lemah. Kita pakai fungsi manual buatan tangan yang penuh kasih sayang.

    🖼️ Preview Gambar: Terintegrasi dengan chafa untuk melihat waifu (atau diagram) langsung di terminal.

    ⚙️ Advanced Config: Edit konfigurasi langsung di dalam aplikasi dengan kursor blok yang satisfying.

🛠️ Cara Install (Jangan Sampai Salah!)

Pastikan kamu punya Go (Golang) terinstall. Kalau belum, sana install dulu, jangan malas!

    Clone atau Download folder ini.

    Siapkan Gambar Keramat: Simpan gambar goldship.png (atau gambar apa saja) di folder yang sama, atau atur path-nya nanti di Settings.

    Build:
    Bash

go build -o goldfile

Jalankan:
Bash

    ./goldfile

    Catatan Penting: Kalau kamu pakai Linux, pastikan install chafa dulu biar fitur preview gambarnya jalan. sudo apt install chafa (atau semacamnya).

🎮 Cara Main (Controls)

Navigasinya gampang banget, bahkan Mejiro McQueen pun bisa pakai ini:
Tombol	Fungsi	Komentar Gold Ship
j / ↓	Turun	Gas pol ke bawah!
k / ↑	Naik	Balik lagi ke atas.
/	SEARCH MODE	Fitur paling OP. Ketik main, .go dan lihat keajaibannya.
Enter	Buka / Edit	Masuk ke folder atau edit file (pake nvim/nano).
Tab	Preview	Intip isinya tanpa buka.
Backspace	Back / Parent	Mundur satu langkah (seperti start lari yang telat).
q	Keluar	"Bye-bye, Trainer!"
🧠 The "Handmade" Flex

Di bagian logic.go, kamu akan menemukan keajaiban dunia ke-8. Kami tidak menggunakan import "strings". Kami membuat fungsi sendiri:

    manualToLower(): Mengubah huruf besar ke kecil dengan memanipulasi byte ASCII (+32). Hardcore.

    manualSplit(): Memecah string berdasarkan koma dengan loop manual. Artisanal.

    manualContains(): Mencocokkan substring dengan nested loop. Traditional.

    recursiveWalk(): Menjelajah folder secara rekursif tanpa filepath.WalkDir. Brave.

Ini bukan Spaghetti Code, ini adalah Yakisoba Code! Enak dan mengenyangkan.
⚙️ Konfigurasi

File config akan otomatis dibuat di: ~/.config/goldfile/config.json

Kamu bisa ganti:

    Editor: Mau pakai vim, nano, atau helix? Bebas.

    Dialogues: Ganti kata-kata mutiara Gold Ship sesuka hatimu.

    Theme: Ubah warna kalau bosan dengan tema Tokyo Night bawaan.

📜 License

Dilindungi oleh Gold Ship Protection Squad. Boleh dicopy, dimodifikasi, asal jangan lupa beli Yakisoba pas ngoding.

"Sekarang, kembali bekerja Trainer! Jangan lupa kasih bintang (Star) di repo ini atau aku tendang!" 🦶✨
