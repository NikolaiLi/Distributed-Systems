Attach to your submission, a *README* file answering the following questions:

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

    We used the following packages
        "fmt"          (to print logs to terminal)
        "math/rand"    (to generate random sequence for handshake)
        "sync"         (to use sync.waitgroup so main runs until handshake is done)
        "time"         (to generate timestamp for packets)

    We created a packet structure to transmit metadata that was useful to print to terminal during the handshake
        type Packet struct {
        	Type      string
        	Seq       int
        	Ack       int
        	ID        string
        	Timestamp time.Time

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

    We used threads to simulate the interaction. Threads are ideal for dealing with sequences/code acting in steps, whereas a 3-way handshake is a reactionary interaction.
    It is state-driven behaviour, not step-driven.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
    We already include a type on packets that could be checked when receiving packets.

    For mixups on packet ordering, we could introduce some kind of metadata that helps keep order on individual packets.
    When receiving packets, store packets in a structure and run through to check if the current ordering is matching the sequence numbers, and reorganize when new packet is not the newest in the sequence.
    When final transmission.

d) In case messages can be delayed or lost, how does your implementation handle message loss?
    We did not include handling of delay and loss.
    The goal would have been to add a timer after sending a packet, and if nothing arrives within the timelimit, the packet is resent, and the timer is halved.
    Do this for three rounds, to give a moment for the delay to get sorted.

e) Why is the 3-way handshake important?
    It ensures synchronization and security during transmission of data. Without it, we end up with delays and/or data loss.