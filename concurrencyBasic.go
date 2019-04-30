package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")           //go concurrency call; occurs at the same time as the second call to count
	count("My $ going bye-bye") // without the call to "go" command in previous call to count,
} //this will never be executed. Now they are executed simultaionously. (Will run forever)

// func main() { //wait groups allow for go to recognize when a concurent func is running, so main doesn't kill the function while running
// 	var waiter sync.WaitGroup //wait group allows us to wait until the
// 	waiter.Add(1)
//
// 	go func() { //annonymous func
// 		count("sheep")
// 		waiter.Done()
// 	}()
// 	waiter.Wait()
// }
func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second)
	}
}

// func main() { // uses channels to optimize appropriate waits
// 	c := make(chan string) // make a channel of type string
// 	go count("sheep", c)   //pass thing along with the channel to count()
//
// 	for {
// 		msg, open := <-c // get message and wait for the channel to close
// 		if !open {
// 			break // now it know the channel is closed exit
// 		}
// 		fmt.Println(msg)
// 	}
// }
// func count(thing string, c chan string) {
// 	for i := 0; i < 5; i++ {
// 		c <- thing //send the value of thing through the channel
// 		time.Sleep(time.Second)
// 	}
// 	close(c) //we need to close the channel in order so our main knows we are finished or it will wait forever
// }
