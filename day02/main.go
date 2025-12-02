package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	inputFile := "input.txt"

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	sumOfInvalidIds1, sumOfInvalidIds2, err := solve(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("#1: Sum of invalid IDs:", sumOfInvalidIds1)
	fmt.Println("#2: Sum of invalid IDs:", sumOfInvalidIds2)

}

func solve(filename string) (int, int, error) {
	intervals, err := readIdsinput(filename)
	if err != nil {
		return 0, 0, err
	}

	return InvalidIds1(intervals), InvalidIds2(intervals), nil
}

func readIdsinput(filename string) ([]Interval, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ids []Interval

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for _, part := range parts {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) == 2 {
				start, err := strconv.Atoi(rangeParts[0])
				if err != nil {
					return nil, err
				}
				end, err := strconv.Atoi(rangeParts[1])
				if err != nil {
					return nil, err
				}
				ids = append(ids, Interval{
					Start: start,
					End:   end,
				})
			}
		}
	}
	return ids, scanner.Err()
}

func InvalidIds1(intervals []Interval) int {
	sumOfInvalidIds := 0
	for _, interval := range intervals {
		for i := interval.Start; i <= interval.End; i++ {
			s := strconv.Itoa(i)

			if len(s)%2 != 0 {
				continue
			}

			mid := len(s) / 2
			left := s[:mid]
			right := s[mid:]

			if left == right {
				sumOfInvalidIds += i
			}
		}
	}

	return sumOfInvalidIds
}

func InvalidIds2(intervals []Interval) int {
	sumOfInvalidIds := 0
	for _, interval := range intervals {
		for i := interval.Start; i <= interval.End; i++ {

			if isInvalidID(i) {
				sumOfInvalidIds += i
			}
		}
	}

	return sumOfInvalidIds
}

func isInvalidID(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		repeats := length / patternLen
		if repeats < 2 {
			continue
		}

		pattern := s[:patternLen]
		candidate := strings.Repeat(pattern, repeats)

		if candidate == s {
			return true
		}
	}

	return false
}
