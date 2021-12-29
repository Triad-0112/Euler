package main

/*type Message struct {
	str  string
	wait chan bool
}
*/

/*
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
*/

//Generator: functio that returns a channel
/*
func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I am leaving!")
}
*/

//Channels as handle on a service
/*
func main() {
	joe := boring("Joe")
	andy := boring("Andy")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-andy)
	}
	fmt.Println("Both of your are boring")
}
*/

//Multiplexing
/*
func fanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c
}
func main() {
	c := fanIn(boring("Joe"), boring("Andy"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Both of you are boring!")
}
*/

//Restoring sequence
/*
func main() {
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	waitForIt := make(chan bool)
	c <- Message(fmt.Sprintf("%s: %d", msg, i), waitForIt)
	time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	<-waitForIt

}
*/

//Fan-in using select
/*
func fanIN(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
*/

//Timeout using slect
/*
func main() {
	c := boring("joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Slow!!!")
			return
		}
	}
}
*/

//Timeout for whole conversation using select
/*
func main() {
	c := boring("Joe")
	timeout := time.After((5 * time.Second))
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Calm!!!")
			return
		}
	}
}
*/

//Quit channel
/*
func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	select {
	case c <- fmt.Sprintf("%s: %d", msg, i):
	case <-quit:
		return

	}
}
*/

//Daisy Chain
/*
func f(left, right chan int) {
	left <- 1 + <-right
}
func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}
*/

//
