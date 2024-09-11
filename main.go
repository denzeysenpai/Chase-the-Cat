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

const up = 'w'
const down = 's'
const left = 'a'
const right = 'd'
const size = 16

func Controller(c rune, pos []int, p bool) {
	var min int = 1
	var max int = size - 1
	if p {
		min = 0
		max = size
	}
	switch c {
	case left:
		if pos[1] > min {
			pos[1]--
		}
	case right:
		if pos[1] < max {
			pos[1]++
		}
	case up:
		if pos[0] > min {
			pos[0]--
		}
	case down:
		if pos[0] < max {
			pos[0]++
		}
	}
}

func __Game__() {
	const player = "▓▓"
	const g = "░░"
	const cat = "{}"
	var pos = []int{0, 0}
	var coin = []int{0, 0}
	var score int
	var moves int = (size * 2) + 6
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
		var floor [size][size]string
		for s := 0; s < size; s++ {
			for d := 0; d < size; d++ {
				floor[s][d] = g
			}
		}
		floor[pos[0]][pos[1]] = player
		floor[coin[0]][coin[1]] = cat
		var i int
		for i = 0; i < size; i++ {
			fmt.Println(floor[i])
		}
		if moves == 0 {
			fmt.Println("Game over!")
			break
		} else {
			moves--
		}
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		Controller(char, pos, true)

		rand := rand.Intn(4)
		move := coinMovement[rand]

		if coin[0] == pos[0] && coin[1] == pos[1] {
			score++
			moves = moves + bonus
		} else {
			Controller(move, coin, false)
		}

		Clear()
	}
}
