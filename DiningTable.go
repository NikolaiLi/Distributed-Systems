package main

func main() {
	forkChannels := make([]chan Request, 5)
	for i := 0; i < 5; i++ {
		forkChannels[i] = make(chan Request)
		go fork(i, forkChannels[i])
	}

	for i := 0; i < 5; i++ {
		left := i
		right := (i + 1) % 5
		go philosopher(i, forkChannels[left], forkChannels[right])
	}
}

func philosopher(id int, left chan Request, right chan Request) {
	replyChan := make(chan bool)

	left <- Request{philosopherID: id, action: "take", reply: replyChan}
	<-replyChan

	right <- Request{philosopherID: id, action: "take", reply: replyChan}
	<-replyChan

	println("Philosopher", id, "is eating")

	left <- Request{philosopherID: id, action: "release", reply: replyChan}
	right <- Request{philosopherID: id, action: "release", reply: replyChan}

	println("Philosopher", id, "is thinking")
}

func fork(id int, requests chan Request) {
	taken := false
	for {
		req := <-requests
		if !taken && req.action == "take" {
			taken = true
			req.reply <- true
		}

		if req.action == "release" {
			taken = false
		}
	}
}

type Request struct {
	philosopherID int
	action        string
	reply         chan bool
}
