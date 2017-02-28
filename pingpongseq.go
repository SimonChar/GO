/*Ping Pong by Simon Charalambous */

package main

import (
    "fmt"
    "os"
    "flag"
    "strconv"
)
type PingPong struct {
    num uint64 
}


func (pingpong *PingPong) Set(inum uint64) {

    pingpong.num = inum

}


func (partner PingPong) Ping(num uint64) {

//time.Sleep(100 * time.Millisecond)

    
    if partner.num > 0 {
        //ch := make(chan int)
        fmt.Printf("Ping..\n")
        partner.num = partner.num -1
        	partner.Pong(partner.num)
    }


}

func (partner PingPong) Pong(num uint64) {

    if partner.num > 0 {
    partner.num = partner.num -1
    fmt.Printf("Pong!\n")
    partner.Ping(partner.num)
    }


}


func main(){

    var ping PingPong
   // var pong PingPong

    
    flag.Parse()
    s := flag.Arg(0)    	
   
    //string to int
    numPingPongs, err := strconv.ParseUint(s,0,32)

    if err != nil {
        fmt.Printf("Please enter a number\n")	
	os.Exit(2)
    }

   if numPingPongs == 0{
       numPingPongs = 1
    }

    //Creating ping and pong actors
     ping.Set(numPingPongs)
    ping.Ping(ping.num)        
    
    //start ping pong
        
    fmt.Println("Two actors have PingPonged each other ")


//fmt.Println(numPingPongs)

}
