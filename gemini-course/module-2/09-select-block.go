/*
In Go, when you want a goroutine to listen to multiple channel operations at the exact same time, you use the select block.

Think of a select block exactly like a Linux select() or poll() system call. It monitors multiple channels simultaneously. Whichever channel is ready to send or receive first wins, its block of code executes, and the rest are ignored.

## Unbuffered and Buffered Channel
// Unbuffered Channel (Capacity = 0) - Requires an immediate handshake
ch := make(chan string)

// Buffered Channel (Capacity = 1) - Has a 1-slot mailbox!
ch := make(chan string, 1)

Unbuffered Channel (make(chan string)): A Direct Handshake.
You are standing in front of a colleague trying to hand them a physical document. You cannot let go of the document until their fingers are physically touching it. If they walk away, you stand there holding the document out forever.
Buffered Channel (make(chan string, 1)): A Single-Slot Mailbox.
There is a physical inbox tray on the desk. You walk up, drop the document into the tray, and walk away immediately. You don't care whether your colleague is in the room or off taking a coffee break. Your job is done, and you exit the building.

UNBUFFERED (Cap = 0):
Sender ───► [ HANDOFF ] ───► Receiver   (Both MUST be present at the exact same microsecond)

BUFFERED (Cap = 1):
Sender ───► [ Slot 1: "Data" ]           (Sender drops data and EXITS immediately)

	│
	└───────► Receiver   (Receiver reads whenever it's ready)
*/
package main

import (
	"fmt"
	"time"
)

func slowWorker(ch chan<- string) {
	time.Sleep(5 * time.Second) // Simulating a hung connection
	ch <- "Successfully fetched payload!"
}

func main() {
	ch := make(chan string)

	go slowWorker(ch)

	fmt.Println("Monitoring worker progress...")

	// --- YOUR CODE HERE ---
	// Write a select statement that listens to:
	// 1. Receiving a message from 'ch'
	// 2. A timeout channel of 2 seconds using 'time.After(2 * time.Second)'

	select {
	case msg := <-ch:
		fmt.Printf("ReceivedMessage: %s", msg)
	case <-time.After(2 * time.Second):
		fmt.Printf("Timeout Breached ! Move on")
	}

}
