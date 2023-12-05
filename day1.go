package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {

	// open file input.txt
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	// read line by line
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	var total int = 0
	for filescanner.Scan() {
		line := filescanner.Text()
		calibration_number := getCalibrationNumber(line)
		total = total + calibration_number
		fmt.Println(line, calibration_number)
	}
	// fmt.Println("word: 3eighteightllkbxkbs9zgznxtj8lfflcst")
	// fmt.Println(getCalibrationNumber("3eighteightllkbxkbs9zgznxtj8lfflcst"))
	fmt.Println("Total: ", total)
}

func getCalibrationNumber(line string) int {
	calibration_number := 0
outer:
	for i := 0; i < len(line); i++ {
		number := getDigitFromChar(line[i])
		if number >= 0 && number <= 9 {
			calibration_number = int(line[i]-'0') * 10
			break outer
		}
		for j := 0; j <= i; j++ {
			if n := getNumberFromWord(line[j : i+1]); n != -1 {
				calibration_number = n * 10
				break outer
			}
		}
	}

outer2:
	for i := len(line) - 1; i >= 0; i-- {
		number := getDigitFromChar(line[i])
		if number >= 0 && number <= 9 {
			calibration_number = calibration_number + int(line[i]-'0')
			break outer2
		}
		for j := len(line) - 1; j >= i; j-- {
			if n := getNumberFromWord(line[i : j+1]); n != -1 {
				calibration_number = calibration_number + n
				break outer2
			}
		}
	}

	return calibration_number
}

// the number can be also a word (e.g. one, two, three, etc.)
// the calibration value can be found by combining the first digit and the last digit (in that order)
func getNumberFromWord(word string) int {
	string_numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(string_numbers); i++ {
		if word == string_numbers[i] {
			return i
		}
	}
	return -1
}

func getDigitFromChar(char byte) int {
	return int(char - '0')
}
