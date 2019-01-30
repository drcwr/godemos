# time.sh
# %Y-%m-%d
# %H:%M:%S

s=$(date +%s);
t=$(date +%F%T)
y=$(date +%Y)
m=$(date +%m)
d=$(date +%d)
dh=$(date +%H)
dm=$(date +%M)
ds=$(date +%S)
daysecond=$((24*60*60))

echo $s
echo $t
echo $y
echo $m
echo $d
echo $dh
echo $dm
echo $ds
echo $daysecond

ddh=$(echo $dh|sed s'/^0//')
echo $ddh
ddm=$(echo $dm|sed s'/^0//')
echo $ddh
dds=$(echo $ds|sed s'/^0//')
echo $ddh

modes=$((${s}%${daysecond} + 8*3600))
echo $modes
days=$((${ddh}*3600+${ddm}*60+${dds}))
echo $days

