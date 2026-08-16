package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

//sending
// func processNum(numChan chan int) {
// 	for num := range numChan{
// 		fmt.Println("Pro", num)
// 		time.Sleep(time.Millisecond*10)
// 	}

// }

// func sum (result chan int , num1 int , num2 int){
// 	res := num1 + num2
// 	result <- res
// }

// func task (done chan bool){
// 	defer func(){done<-true}()
// 	fmt.Println("Pro...")
// }

// func emailSend(emailChan <-chan string, done chan<- bool){
// 	defer func(){done<-true}()
// 	for email := range emailChan{
// 		fmt.Println("Sen :", email)
// 		time.Sleep(time.Millisecond*10)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(){
		chan1<-10
	}()

	go func(){
		chan2<-"pong"
	}()

	for i:=0 ;i<2;i++{
		select{
		case chan1val := <- chan1:
			fmt.Println("rec1",chan1val)
		case chan2val := <- chan2:
			fmt.Println("rec2",chan2val)
		}
	}

	// mes := make(chan string)

	// mes <- "ping" //blocking

	// msg := <- mes

	// fmt.Println(msg)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// result := make(chan int)
	// go sum(result,4,5)
	// res := <- result
	// fmt.Println(res)

	// done := make(chan bool)
	// go task(done)
	// <- done

	// done := make(chan bool)

	// emailChan := make(chan string, 100)

	// go emailSend(emailChan, done)

	// for i:=1 ;i<=10;i++{
	// 	emailChan <- fmt.Sprintf("%d@gamil.com",i)
	// }



	// fmt.Println("done sending")
	// close(emailChan)
	// emailChan <- "1@exp"
	// emailChan <- "2@exp"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	// <- done

}
