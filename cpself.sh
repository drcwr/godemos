cp cpself.sh /tmp/
cd /tmp
sed -i '1,5s/^/#/' cpself.sh
sh cpself.sh
exit
pwd
ls
echo "hello"

