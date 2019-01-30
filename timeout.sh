#!/bin/bash

#按空格开始60秒的倒计时
#-n表示接受字符的数量，1表示只接受一个字符 

a()
{
  for i in `seq -w 60 -1 0`
  do

  echo -ne "\b\b$i"
  sleep 0.1
  done
  echo ""

}

read -n 1 -p "[ Press Space ] " space
[[ "$space" = "" ]] && a
