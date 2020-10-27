package main

func counter(ch chan<- int, limit int) {
	i := 1
	for i != limit {
		i += 1
		ch <- i
	}
	close(ch)
}

func filter(prime int, recv <-chan int, send chan<- int) {
	var i int
	if i = <-recv; i%prime > 0 {
		send <- 1
	}
}

func sieve(prime chan int) {
	ch := make(chan int)
	go counter(ch, 100)
	var p int
	for {
		p = <-ch
		prime <- p
		send := make(chan int)
		go filter(p, ch, send)
		ch = send
	}
}

func main() {
	prime := make(chan int)
	go sieve(prime)
	println(<-prime, <-prime, <-prime, <-prime)
}
