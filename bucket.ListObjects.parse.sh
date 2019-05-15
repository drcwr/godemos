

#head -n 1 res600all2enternf
#201905/00/00/00/f4/000000f4869cb0ccca7fdce2ee7ad0ae.jpg

file res600all
#res600all: Little-endian UTF-16 Unicode text, with very long lines, with CRLF line terminators

iconv -f "UTF-16" -t ASCII res600all -o res600all2
sed 's/  201905/\n201905/g' res600all2 >res600all2enter
awk NF res600all2enter > res600all2enternf
cat res600all2enternf|awk -F "/" '{print $NF}'|awk -F "." '{print $1}' >res600all2enternfonlynamelast
cat res600all2enternfonlynamelast |sort -n -u > res600all2enternfonlynamelastsort


#head -n 2 2019-05-14-20-43-32_EXPORT_CSV_867_913_0.csv 
#goods_id,old_img_url
#8407142,https://img-g.taojiji.com/201905/4d/47/c1/14/4d47c1146866868fb7cc62d6188f48a8.jpg
#8429391,http://sjimg.taojiji.com/20190514/img/155781559944738131.jpg
#8429392,https://img-g.taojiji.com/201905/8c/6b/6e/a9/8c6b6ea943685b43683c7904ce7b8e24.jpg

cat  2019-05-14-20-43-32_EXPORT_CSV_867_913_0.csv |awk -F "/" '{print $NF}'|awk -F "." '{print $1}' > csvonlyname
cat csvonlyname|sort -n -u >csvonlynamesort



diff csvonlynamesort res600all2enternfonlynamelastsort > diffres
grep "<" diffres > csverr
cat csverr|awk '{print $NF}' > csverronlyname









