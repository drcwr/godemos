package main

import (
    "fmt"
)

func main() {
    var sa1 = make([]string,0,100)
    var sa2 = make([]string,0,100)
    
    sa1 = append(sa1,"abc")
    sa1 = append(sa1,"efg")
    sa1 = append(sa1,"hij")
    sa1 = append(sa1,"kl")

    fmt.Printf("1 sa1 p=%p,%#v,%d,%d\n",sa1,sa1,len(sa1),cap(sa1))
    fmt.Printf("1 sa2 p=%p,%#v,%d,%d\n",sa2,sa2,len(sa2),cap(sa2))
    for _,v := range sa1 {
        sa2 = append(sa2,v)
    }

    fmt.Printf("2 sa1 p=%p,%#v,%d,%d\n",sa1,sa1,len(sa1),cap(sa1))
    fmt.Printf("2 sa2 p=%p,%#v,%d,%d\n",sa2,sa2,len(sa2),cap(sa2))      

    sa1,sa2 = sa2,sa1

    sa2 = sa2[0:0]
    fmt.Printf("3 sa1 p=%p,%#v,%d,%d\n",sa1,sa1,len(sa1),cap(sa1))
    fmt.Printf("3 sa2 p=%p,%#v,%d,%d\n",sa2,sa2,len(sa2),cap(sa2))      
 
}

