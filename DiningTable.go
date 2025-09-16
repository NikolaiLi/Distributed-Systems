package main

func main() {
	forkChannels := make([]chan Request, 5)

	counter := make([]int, 5)
	counterFlag := make([]bool, 5)

	for i := 0; i < len(counter); i++ {
		println("Philosopher", i, "is thinking")
		if counter[i] >= 3 {
			counterFlag[i] = true
		} else {
			counterFlag[i] = false
		}
	}

	for i := 0; i < len(counterFlag); i++ {
		if counterFlag[i] == false {
			for i := 0; i < 5; i++ {
				forkChannels[i] = make(chan Request)
				go fork(i, forkChannels[i])
			}

			for i := 0; i < 5; i++ {
				left := i
				right := (i + 1) % 5
				go philosopher(i, forkChannels[left], forkChannels[right], counter)
			}
		}
	}
}

func philosopher(id int, left chan Request, right chan Request, counter []int) {
	replyChan := make(chan bool)

	for counter[id] < 3 {
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

		println("Philosopher", id, "is eating")
		counter[id]++

		left <- Request{philosopherID: id, action: "release", reply: replyChan}
		right <- Request{philosopherID: id, action: "release", reply: replyChan}

		println("Philosopher", id, "is thinking")
	}
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
