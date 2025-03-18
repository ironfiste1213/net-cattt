package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var (
	mu sync.Mutex

)
var maxclient = 10
func main() {
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} 
	lisnner, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalln("not connect")
	}
	
    asciiArt := `
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    \.       | \' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     \-\       \--'
    `
	for {

		conn, err := lisnner.Accept()
		if err != nil {
			fmt.Println("err!!!")
			continue
		}
		_, err = conn.Write([]byte(asciiArt))
		if err != nil {
			fmt.Println("err?!!")
			continue
		}
	}

	 

}