package main // looks for a func called main, only applies to main keyword

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// static typing
// strongly typed
// compiled
// fast compile
// built in concurrency
// simple

func main() {
	__Game__()
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func __Game__() {
	const up = 'w'
	const down = 's'
	const left = 'a'
	const right = 'd'

	const player = "▓▓"
	const g = "░░"
	var pos = []int{0, 0}
	var coin = []int{0, 0}
	var score int
	var moves int = 30
	var coinMovement = [4]rune{'w', 'a', 's', 'd'}

	var bonus = 9
	for {
		if coin[0] == pos[0] && coin[1] == pos[1] {
			rand.Seed(time.Now().UnixNano())
			randomInt1 := rand.Intn(10)
			randomInt2 := rand.Intn(10)
			coin = []int{randomInt1, randomInt2}
		}
		if score == 10 || score == 20 || score == 30 {
			bonus--
		}
		fmt.Println(" ")
		fmt.Println("Score: ", score)
		fmt.Println("Moves: ", moves)
		fmt.Println(" ")
		var floor = [10][10]string{
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
			{g, g, g, g, g, g, g, g, g, g},
		}
		floor[pos[0]][pos[1]] = player
		floor[coin[0]][coin[1]] = "M"
		var i int
		for i = 0; i < 10; i++ {
			fmt.Println(floor[i])
		}
		reader := bufio.NewReader(os.Stdin)
		if moves == 0 {
			break
		} else {
			moves--
		}
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case left:
			if pos[1] > 0 {
				pos[1]--
			}
		case right:
			if pos[1] < 10 {
				pos[1]++
			}
		case up:
			if pos[0] > 0 {
				pos[0]--
			}
		case down:
			if pos[0] < 10 {
				pos[0]++
			}
		}

		rand := rand.Intn(4)
		move := coinMovement[rand]
		switch move {
		case left:
			if coin[1] > 1 {
				coin[1]--
			}
		case right:
			if coin[1] < 9 {
				coin[1]++
			}
		case up:
			if coin[0] > 1 {
				coin[0]--
			}
		case down:
			if coin[0] < 9 {
				coin[0]++
			}
		}

		if coin[0] == pos[0] && coin[1] == pos[1] {
			score++
			moves = moves + bonus
		}

		Clear()
	}
}
