/*Ping Pong by Simon Charalambous */                                                                     
                                                                                                         
package main                                                                                             
                                                                                                         
import (                                                                                                 
    "flag"                                                                                               
    "strconv"                                                                                            
)                                                                                                        
                                                                                                         
 func Player1(table_ping, table_pong, table_out chan uint64) {                                           
                                                                                                         
    var num uint64 =0                                                                                    
                                                                                                         
    for {                                                                                                
        num = <- table_ping                                                                              
            if num > 0 {                                                                                 
                //fmt.Println("Ping", num)                                                               
                table_pong <- (num-1)                                                                    
            }else{                                                                                       
                break                                                                                    
                }                                                                                        
    }                                                                                                    
    table_out <- num                                                                                     
}                                                                                                        
                                                                                                         
func Player2(table_ping, table_pong, table_out chan uint64) {                                            
                                                                                                         
    var num uint64 = 0                                                                                   
                                                                                                         
    for {                                                                                                
        num = <- table_pong                                                                              
        if num > 0 {                                                                                     
            //fmt.Println("Pong", num)                                                                   
            table_ping <- num -1                                                                         
         }else{                                                                                          
         break                                                                                           
                }                                                                                        
    }                                                                                                    
    table_out <- num                                                                                     
} 

func main() {                                                                                            
                                                                                                         
    flag.Parse()                                                                                         
    s := flag.Arg(0)                                                                                     
                                                                                                         
    //string to int                                                                                      
    numPingPongs, err :=  strconv.ParseUint(s, 0, 32)                                                    
                                                                                                         
    if err != nil || numPingPongs ==0 {                                                                  
    numPingPongs = 2                                                                                     
        }                                                                                                
                                                                                                         
    numPingPongs = 2 * numPingPongs                                                                      
                                                                                                         
    table_ping := make(chan uint64)                                                                      
    table_pong := make(chan uint64)                                                                      
    table_out := make(chan uint64)                                                                       
                                                                                                         
    //start ping pong with go routines                                                                   
    go Player1(table_ping, table_pong, table_out)                                                        
    go Player2(table_ping, table_pong, table_out)                                                        
                                                                                                         
    table_ping <- numPingPongs                                                                           
                                                                                                         
    <- table_out                                                                                         
}                     
                        
