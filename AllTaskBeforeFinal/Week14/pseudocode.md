program iceCream
kamus
kTidur, tMain, tKebun : Integer
kondisi1, kondisi2, kondisi3 : boolean
hasil : string
algoritma
input(kTidur, tMain, tKebun)
kondisi1 = kTidur >= 100
kondisi2 = kTidur >= 80 and tMain >= 80
kondisi3 = kTidur >= 60 and tMain >= 75 and tKebun >= 60
if kondisi1 or kondisi2 or kondisi3 then
output("Ice Cream")
else
output("Tidak")
endif
endprogram
