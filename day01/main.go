package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Rotation struct {
	RotationDirection string
	NumberOfClicks    int
}

func main() {
	inputFile := "input.txt"

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	firstPassword, secondPassword, err := solve(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First password:", firstPassword)
	fmt.Println("Second password:", secondPassword)

}

func solve(filename string) (int, int, error) {
	rotations, err := readRotationsInput(filename)
	if err != nil {
		return 0, 0, err
	}

	return findFirstPassword(rotations), findSecondPassword(rotations), nil
}

func readRotationsInput(filename string) ([]Rotation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rotations []Rotation

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		numberOfClicks, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		rotations = append(rotations, Rotation{
			RotationDirection: direction,
			NumberOfClicks:    numberOfClicks,
		})
	}
	return rotations, scanner.Err()
}

func findFirstPassword(rotations []Rotation) int {
	currentPosition := 50
	password := 0

	for _, rotation := range rotations {
		switch rotation.RotationDirection {
		case "L":
			currentPosition -= rotation.NumberOfClicks
			for currentPosition < 0 {
				currentPosition += 100
			}
		case "R":
			currentPosition += rotation.NumberOfClicks
			for currentPosition > 99 {
				currentPosition -= 100
			}
		}

		if currentPosition == 0 {
			password += 1
		}
	}

	return password
}

func findSecondPassword(rotations []Rotation) int {
	currentPosition := 50
	password := 0

	for _, rotation := range rotations {

		var step int
		switch rotation.RotationDirection {
		case "L":
			step = -1
		case "R":
			step = 1
		default:
			panic("unknown direction: " + rotation.RotationDirection)
		}

		for i := 0; i < rotation.NumberOfClicks; i++ {
			currentPosition += step

			if currentPosition < 0 {
				currentPosition += 100
			}

			if currentPosition > 99 {
				currentPosition -= 100
			}

			if currentPosition == 0 {
				password += 1
			}
		}
	}

	return password
}
