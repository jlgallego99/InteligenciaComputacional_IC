set title "Evolución del fitness a lo largo de las generaciones en el algoritmo genético genérico"
set term pngcairo dashed size 1400,1050
set output "result/evolution_generic.png"

# columna en eje x : columna en eje y
set xlabel "Generación"
set ylabel "Fitness"
plot "result/generic_100_10000.txt" using 1:2 title "Algoritmo genérico" with l lw 4, \

set title "Evolución del fitness a lo largo de las generaciones en las variantes de los evolutivos"
set term pngcairo dashed size 1400,1050
set output "result/evolution_variants.png"

# columna en eje x : columna en eje y
set xlabel "Generación"
set ylabel "Fitness"
plot "result/baldwinian_10_10.txt" using 1:2 title "Algoritmo baldwiniano (10 ind)" with l lw 4, \
    "result/baldwinian_50_10.txt" using 1:2 title "Algoritmo baldwiniano (50 ind)" with l lw 4, \
    "result/lamarckian_50_10.txt" using 1:2 title "Algoritmo lamarckiano (50 ind)" with l lw 4, \
    "result/lamarckian_10_10.txt" using 1:2 title "Algoritmo lamarckiano (10 ind)" with l lw 4, \
    "result/lamarckian_10_100.txt" using 1:2 title "Algoritmo lamarckiano (10 ind)" with l lw 4
