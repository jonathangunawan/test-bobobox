# Principle of this structure

Struktur dalam proyek ini bisa dibilang mengikuti CLEAN CODE PRINCIPLES tapi tidak terlalu mirip. Di proyek ini dibagi menjadi 7 bagian yaitu:

1. cmd (tempat seluruh package main, nantinya bisa dibagi per folder lagi tergantung service contohnya http)
2. entity (tempat seluruh entitas(struct, var, etc) yang akan digunakan secara bersama" dan tipenya umum)
3. files/config (tempat seluruh macam config)
4. handler (tempat seluruh macam handler yang akan menghandle segala data yang masuk ke sistem)
5. infra (tempat seluruh konfigurasi untuk infrastruktur misalnya baca config dan init database)
6. repository (tempat seluruh code yang akan mengakses data ke sumber data, misalnya database Postgre)
7. usecase (tempat seluruh code yang bersangkutan dengan businnes logic)
