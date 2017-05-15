package main
//
// import (
//     "fmt"
//     "runtime"
// )
import "time"
import "fmt"
//
// func main() {
//      c := make(chan int)
//      go func() {
//         c <- 42
//     }()
//          // write to a channel
//     //  time.Sleep(time.Second * 2)
//      val := <-c // read from a channel
//      println(val)
// }


func main() {
    //  message := make(chan string, 5) // no buffer
     count := 1
     var temp [2]int
    //  message := make(chan string, 5) // no buffer
     start := time.Now()
       func() {
          for i := 1; i <= count; i++ {
            // fmt.Println(len(temp))
               temp[i]  = i
               fmt.Println("send message " , temp)

              //  message <- fmt.Sprintf("message %d", i*1000)

          }
          // close (message)
     }()

    //  for msg := range message {
    //       fmt.Println(msg)
    //  }
     elapsed := time.Since(start)
     fmt.Println(elapsed)

}
