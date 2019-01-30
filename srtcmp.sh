GSYMBOL="1"
    if [ ${varsymbol}z != ${GSYMBOL}z ];then
        echored "----" "tokeninfo.info" "symbol ERROR,GSYMBOL:" ${GSYMBOL} ${varsymbol} 
        exit 1
    fi

