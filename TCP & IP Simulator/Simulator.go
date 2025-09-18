package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //gives rand.int new seed each run
	ServerToClientChannel := make(chan Packet, 1)
	ClientToServerChannel := make(chan Packet, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go client(ClientToServerChannel, ServerToClientChannel, &wg)
	go server(ServerToClientChannel, ClientToServerChannel, &wg)

	wg.Wait()
}

func client(send chan<- Packet, receive <-chan Packet, wg *sync.WaitGroup) {
	defer wg.Done()
	id := "Client"
	//client sends first packet SYN
	InitSeq := rand.Int()
	initPacket := makeSYN(InitSeq, id)
	send <- initPacket           //sends initial SYN packet to server
	previousPacket := initPacket //stores client's most recent packet as previous packet
	fmt.Printf("[Client] SEND SYN Seq=%d at %s\n",
		initPacket.Seq,
		initPacket.Timestamp.Format(time.RFC3339Nano))

	//client receives servers packet SYNACK and sends ACK
	newPacket := <-receive //receives SYN-ACK packet from server
	fmt.Printf("[Client] RECV SYN-ACK Seq=%d Ack=%d after %v (sent %s)\n",
		newPacket.Seq,
		newPacket.Ack,
		time.Since(newPacket.Timestamp),
		newPacket.Timestamp.Format(time.RFC3339Nano))
	if newPacket.Ack == previousPacket.Seq+1 {
		newSeq := newPacket.Ack
		newAck := newPacket.Seq + 1
		newPacket := makeACK(newSeq, newAck, id)
		send <- newPacket          //sends ACK to server
		previousPacket = newPacket //stores client's most recent sent packet as previous packet
		fmt.Printf("[Client] SEND ACK Seq=%d Ack=%d at %s\n",
			newPacket.Seq,
			newPacket.Ack,
			newPacket.Timestamp.Format(time.RFC3339Nano))
	}

}

func server(send chan<- Packet, receive <-chan Packet, wg *sync.WaitGroup) {
	defer wg.Done()
	id := "Server"

	//server receives client's packet SYN and sends SYNACK
	initPacket := <-receive
	fmt.Printf("[Server] RECV SYN Seq=%d after %v (sent %s)\n",
		initPacket.Seq,
		time.Since(initPacket.Timestamp),
		initPacket.Timestamp.Format(time.RFC3339Nano))
	initialAck := initPacket.Seq + 1                         //stores client packet seq as ack
	serverResponse := makeSYNACK(rand.Int(), initialAck, id) //server packet is created with its own seq
	send <- serverResponse
	previousResponse := serverResponse //stores servers most recent response packet as previous packet
	fmt.Printf("[Server] SEND SYN-ACK Seq=%d Ack=%d at %s\n",
		serverResponse.Seq,
		serverResponse.Ack,
		serverResponse.Timestamp.Format(time.RFC3339Nano))

	newPacket := <-receive
	if newPacket.Ack == previousResponse.Seq+1 {
		fmt.Printf("[Server] RECV ACK Seq=%d Ack=%d after %v (sent %s)\n",
			newPacket.Seq,
			newPacket.Ack,
			time.Since(newPacket.Timestamp),
			newPacket.Timestamp.Format(time.RFC3339Nano))
		//Ack received and verified, handshake is completed
	}
}

func makeSYN(seq int, id string) Packet {
	return Packet{
		Type:      "syn",
		Seq:       seq,
		ID:        id,
		Timestamp: time.Now(),
	}
}

func makeSYNACK(seq int, ack int, id string) Packet {
	return Packet{
		Type:      "synack",
		Seq:       seq,
		Ack:       ack,
		ID:        id,
		Timestamp: time.Now(),
	}
}

func makeACK(seq int, ack int, id string) Packet {
	return Packet{
		Type:      "ack",
		Seq:       seq,
		Ack:       ack,
		ID:        id,
		Timestamp: time.Now(),
	}
}

type Packet struct {
	Type      string
	Seq       int
	Ack       int
	ID        string
	Timestamp time.Time
}
