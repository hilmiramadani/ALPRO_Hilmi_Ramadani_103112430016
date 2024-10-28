SOAL 4

Program   = Menentukan suatu bilangan memiliki digit terurut atau tidak

kamus     = var bilangan, angka1, angka2, angka3 : int

Algoritma = input(bilangan)
            angka1 = bilangan / 100
	        angka2 = (bilangan % 100) / 10
	        angka3 = bilangan % 10
            IF angka1 < angka2 AND angka2 < angka3 THEN
                PRINT "true"
                PRINT "terurut secara membesar", bilangan
            ELSE IF angka1 > angka2 AND angka2 > angka3 THEN
                PRINT "true"
                PRINT "terurut mengecil", bilangan
            ELSE
                PRINT "false"
                PRINT "tidak terurut", bilangan
            END IF
        
endprogram