// package main

// import "fmt"

// type paymenter interface{
// 	pay(amount float32)
// }

// type payment struct{
// 	gateway paymenter
	
// }

// func (p payment) makePayment(amount float32){
// 	// razorpayGw := razorpay{}
// 	//stripeGw := stripe{}
// 	//razorpayGw.pay(amount)
// 	p.gateway.pay(amount)
// }

// type razorpay struct{}

// func (r razorpay) pay(amount float32){
// 	fmt.Println("making pay raz : ",amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	fmt.Println("making pay str : ",amount)
// }

// func main(){
// 	stripeGw := razorpay{}
// 	newpay := payment{
// 		gateway: stripeGw,
// 	}
// 	newpay.makePayment(100)
// }


package main

func prsl[T any](sl []T){
	for _, i := range sl{
		println(i)
	}
}

func prsls(sl []string){
	for _,i := range sl{
		println(i)
	}
}

type stack [T any] struct {
	ele []int
}

func main(){
	var sl = []int{1,2,3}
	prsl(sl)
	var sls = []string{"ram","raj"}
	prsls(sls)
	mys := stack [int]{
		ele : []int{1,2,3},
	}
	println(mys)
}