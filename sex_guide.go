package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Position string

// Define the "enum" values for Role
const (
	Missionary = "Missionary"
	Doggy      = "Doggy"
	Cowgirl    = "Cowgirl"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var position Position = Missionary
	var difficulty int

	fmt.Println("This is a guide on how to finish any woman.")
	fmt.Printf("Starting in %s.\n", position)
	fmt.Println("Please enter the difficulty between 1 and 10: \t")
	for {
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			reader.ReadString('\n')
			fmt.Println("Please enter a valid number")
			continue
		}
		validateDifficulty(&difficulty)
		break
	}
	finishWoman(difficulty, &position)
}

func validateDifficulty(difficulty *int) {
	for *difficulty < 0 || *difficulty > 10 {
		fmt.Println("Please enter a difficulty between 1 and 10")
		fmt.Scanln(difficulty)
		continue
	}
}

func finishWoman(difficulty int, position *Position) {
	fmt.Println("Calculating how to finish her...")
	time.Sleep(2 * time.Second)
	if difficulty > 5 {
		fmt.Println("You can't finish her with just sex.")
		fmt.Println("Calculating foreplay necessary...")
		time.Sleep(2 * time.Second)
		basicForeplay()
		thrusting(position)
		if difficulty > 8 {
			time.Sleep(time.Second)
			fmt.Println("She's gonna need more work...")
			time.Sleep(time.Second)
			eatingOut()
			thrusting(position)
			if difficulty == 10 {
				fmt.Println("Busting out the pillow under the back...")
				thrusting(position)
				fmt.Println("Final thrust!")
				time.Sleep(2 * time.Second)
				fmt.Println("AHHHHHHHHH AHHHHHHHHH AHHHHHHHHH AHHHHHHHHH AHHHHHHHHHHHHHHHH")
				time.Sleep(2 * time.Second)
				fmt.Println("You have finished her! Congratulations!")
				return
			}
			fmt.Println("You have finished her!")
		} else {
			fmt.Println("Congratulations, you have finished her!")
		}
	} else {
		if difficulty == 0 {
			fmt.Println("She nutted to your presence.")
			return
		}
		fmt.Println("You can finish her easily.")
		thrusting(position)
		fmt.Println("Easy finish, good stuff!")
	}
}
func basicForeplay() {
	randomNumber := rand.Intn(3)
	for num := 0; num < randomNumber+5; num++ {
		if num%2 == 0 {
			fmt.Println("Fingering...")
		} else {
			fmt.Println("DJing her...")
		}
		time.Sleep(time.Second)
	}
}
func eatingOut() {
	randomNumber := rand.Intn(5)
	for num := 0; num < randomNumber+3; num++ {
		fmt.Println("Eating her out...")
		time.Sleep(time.Second)
		if num == 6 {
			fmt.Println("Jaw is starting cramp...")
			time.Sleep(time.Second)
		}
	}
	time.Sleep(2 * time.Second)
}

func thrusting(position *Position) {
	randomNumber := rand.Intn(4)
	for num := 0; num < randomNumber+4; num++ {
		fmt.Printf("Thrusting in %s...\n", *position)
		time.Sleep(time.Second)
		if num == 6 {
			fmt.Println("You're getting tired, but you're almost there!")
			time.Sleep(time.Second)
			fmt.Printf("Thrusting in %s...\n", *position)
			time.Sleep(time.Second)
		}
		if num == 5 {
			time.Sleep(time.Second)
			fmt.Println("Your penis is too weak!")
			time.Sleep(time.Second)
			fmt.Println("Maybe you should try a different position...")
			time.Sleep(time.Second)
			switchPositions(position)
			thrusting(position)
		}
	}
	time.Sleep(time.Second)
}

func switchPositions(position *Position) {
	randomNumber := rand.Intn(3)
	fmt.Printf("Currently in %s\n", *position)
	time.Sleep(time.Second)
	fmt.Println("Switching positions...")
	time.Sleep(2 * time.Second)
	if randomNumber == 0 {
		if *position == Doggy {
			fmt.Println("She doesn't wanna switch, trying again...")
			time.Sleep(time.Second)
			switchPositions(position)
		}
		*position = Doggy
		fmt.Printf("Switched to %s\n", *position)
	} else if randomNumber == 1 {
		if *position == Cowgirl {
			fmt.Println("She doesn't wanna switch, trying again...")
			time.Sleep(time.Second)
			switchPositions(position)
		}
		*position = Cowgirl
		fmt.Printf("Switched to %s\n", *position)
	} else if randomNumber == 2 {
		if *position == Missionary {
			fmt.Println("She doesn't wanna switch, trying again...")
			time.Sleep(time.Second)
			switchPositions(position)
		}
		*position = Missionary
		fmt.Printf("Switched to %s\n", *position)
	}
}
