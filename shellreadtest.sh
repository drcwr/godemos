
READTIMEOUT=5
for i in `seq 0 ${READTIMEOUT}`;do
    read -t 1 -p "you have ${READTIMEOUT} seconds,Please enter Y/N: (Y)" inputstr 
    echo "Input:$inputstr"
    if [ ${inputstr}z == "n"z ] || [ ${inputstr}z == "N"z ];then
        echo "exit 1"
        exit 1
    else
        echo $inputstr
        echo "go on todo"
    fi
done