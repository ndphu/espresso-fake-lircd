package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var (
	Port            = 8765
	SendIntervalMin = 2 // 2 seconds
	SendIntervalMax = 3 // 20 seconds
)

func main() {
	rand.Seed(time.Now().Unix())
	answers := []string{
		"It_is_certain",
		"It_is_decidedly_so",
		"Without_a_doubt",
		"Yes_definitely",
		"You_may_rely_on_it",
		"As_I_see_it_yes",
		"Most_likely",
		"Outlook_good",
		"Yes",
		"Signs_point_to_yes",
		"Reply_hazy_try_again",
		"Ask_again_later",
		"Better_not_tell_you_now",
		"Cannot_predict_now",
		"Concentrate_and_ask_again",
		"Don't_count_on_it",
		"My_reply_is_no",
		"My_sources_say_no",
		"Outlook_not_so_good",
		"Very_doubtful",
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
