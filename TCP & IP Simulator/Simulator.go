package main

func main() {
	serverToClientChannel := make(chan Request, 1)
	ClientToServerChannel := make(chan Request, 1)

	go client(ClientToServerChannel)
	go server(serverToClientChannel)
}

type Request struct {
}

func client(requests chan Request) {

}

func server(requests chan Request) {

}
