package main

import "fmt"

func main() {
	jobs := make(chan int, 7)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)
			} else {
				fmt.Println("recevied all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("sent jobs")
	}
	close(jobs)
	fmt.Println("close jobs")
	<-done

}
