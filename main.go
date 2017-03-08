package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var (
	Port            = 8765
	SendIntervalMin = 6  // 6 seconds
	SendIntervalMax = 10 // 10 seconds
)

func main() {
	rand.Seed(time.Now().Unix())
	answers := []string{
		"Remote001",
		"Remote002",
		"Remote003",
		"Remote004",
		"Remote005",
		"Remote006",
		"Remote007",
		"Remote008",
		"Remote009",
		"Remote010",
		"Remote011",
		"Remote012",
		"Remote013",
		"Remote014",
		"Remote015",
		"Remote016",
		"Remote017",
		"Remote018",
		"Remote019",
		"Remote020",
	}
	buttons := []string{
		"NUM_1",
		"NUM_2",
		"NUM_3",
		"NUM_4",
		"NUM_5",
		"NUM_6",
		"NUM_7",
		"NUM_8",
		"NUM_9",
		"NUM_0",
	}
	fmt.Println("A fake lircd which open a port and send fake lircd message to clients which connect to it")

	fmt.Println("Using port", Port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Waiting connection...")
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected from", conn.RemoteAddr().String())

		go func() {
			for {
				fmt.Println("Sending message...")
				_, err := conn.Write([]byte(fmt.Sprintf("0000000000000009 %d %s %s\n", rand.Intn(100), buttons[rand.Intn(len(buttons))], answers[rand.Intn(len(answers))])))
				if err != nil {
					fmt.Println("Failed to write with error", err)
					break
				}
				// make it more random
				randSleep := rand.Intn(SendIntervalMax-SendIntervalMin) + SendIntervalMin
				time.Sleep(time.Duration(randSleep) * time.Second)

			}
		}()
	}

}
