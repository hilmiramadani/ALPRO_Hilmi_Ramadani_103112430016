kamus
T, V, VolumeSekarang : integer
algoritma
input(T)
VolumeSekarang = 0

    while VolumeSekarang < T do
        input(V)
        VolumeSekarang = VolumeSekarang + V

        if VolumeSekarang >= T then
            output("true")
            break
        else
            output("false")
        endif
    endwhile

endprogram
