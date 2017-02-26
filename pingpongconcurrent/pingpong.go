/*Ping Pong by Simon Charalambous */

package main

import (
       "flag"
       "fmt"
       "os"
       "strconv"
)

func PingPong(name string, num uint64, table chan uint64) {

     for num > 0 {
     	 num = num - 1
	     fmt.Println(name)
	     }
	     table <- num
}

func main() {

     // var pong PingPong

     flag.Parse()
     s := flag.Arg(0)

     //string to int
     numPingPongs, err := strconv.ParseUint(s, 0, 32)

     if err != nil {
     	fmt.Printf("Please enter a number\n")
	 os.Exit(2)
	 }

     if numPingPongs == 0 {
         numPingPongs = 1
	 }

    table := make(chan uint64)

    //start ping pong with go routines
    go PingPong("Ping", numPingPongs/2, table)
    go PingPong("Pong", numPingPongs/2, table)

    ping, pong := <-table, <-table

    fmt.Println(ping, pong)
    fmt.Println("Two actors have PingPonged each other ")

}