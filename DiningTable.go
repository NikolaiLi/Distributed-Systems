package main

import (
	"fmt"
	"sync"
)

type Request struct {
	philosopherID int
	action        string
	reply         chan bool
}

func main() {
	forkChannels := make([]chan Request, 5)

	for i := 0; i < 5; i++ {
		forkChannels[i] = make(chan Request)
		go fork(i, forkChannels[i])
	}

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		left := forkChannels[i]
		right := forkChannels[(i+1)%5]
		go philosopher(i, left, right, &wg)
	}

	wg.Wait()
	fmt.Println("The philosophers had a nice dinner")
}

func philosopher(id int, left chan Request, right chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	meals := 0

	for meals < 3 {

		replyChan := make(chan bool)

		if id != 4 {
			left <- Request{philosopherID: id, action: "take", reply: replyChan}
			<-replyChan

			right <- Request{philosopherID: id, action: "take", reply: replyChan}
			<-replyChan
		} else {
			right <- Request{philosopherID: id, action: "take", reply: replyChan}
			<-replyChan

			left <- Request{philosopherID: id, action: "take", reply: replyChan}
			<-replyChan
		}

		fmt.Println("Philosopher", id, "is eating")
		meals++

		left <- Request{philosopherID: id, action: "release", reply: nil}
		right <- Request{philosopherID: id, action: "release", reply: nil}

		fmt.Println("Philosopher", id, "is thinking")
	}
}

func fork(id int, requests chan Request) {
	taken := false
	var queue []Request

	for {
		req := <-requests
		switch req.action {
		case "take":
			if !taken {
				taken = true
				if req.reply != nil {
					req.reply <- true
				}
			} else {
				queue = append(queue, req)
			}
		case "release":
			if len(queue) > 0 {
				next := queue[0]
				queue = queue[1:]
				if next.reply != nil {
					next.reply <- true
				}
			} else {
				taken = false
			}
		}
	}
}
