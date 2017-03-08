package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var (
	Port            = 8765
	SendIntervalMin = 2  // 2 seconds
	SendIntervalMax = 10 // 20 seconds
)

func main() {
	rand.Seed(time.Now().Unix())
	remoteNames := []string{
		"Remote000",
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
				repeat := rand.Intn(100)
				button := buttons[rand.Intn(len(buttons))]
				remoteName := remoteNames[rand.Intn(len(remoteNames))]
				fmt.Println("Sending event", remoteName, button, repeat)
				_, err := conn.Write([]byte(fmt.Sprintf("0000000000f40bf0 %x %s %s\n", repeat, button, remoteName)))
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
