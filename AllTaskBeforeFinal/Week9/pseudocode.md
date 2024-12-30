kamus
siswa, mobil, remaining : integer
kapasitas : constant integer
algoritma
kapasitas = 7
input(siswa)
mobil = siswa div kapasitas
remaining = siswa mod kapasitas
if remaining > 0 then
mobil = mobil + 1
output(mobil - 1, "mobil penuh dan 1 mobil berisi", remaining, "orang")
else
output(mobil, "mobil penuh")
endif
endprogram
