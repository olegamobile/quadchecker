package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

//
func runCommand(command, a, b string) (string, error) {
	cmd := exec.Command(command, a, b)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func counter(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
		} else {
			if flag == true {
				countX++
			}
		}
	}
	return countX, countY
}

func getpipe() string {
	scanner := bufio.NewScanner(os.Stdin)
	result := ""
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	return result
}

func PrintNbr(n int) string {
	str := ""
	digit := ""
	sign := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n *= -1
		sign = "-"
	}
	for n > 0 {
		digit = string(rune(n%10) + '0')
		str = digit + str
		n /= 10
	}
	str = sign + str
	return str
}

func main() {
	input := getpipe()
	var quadout string
	count := 0
	x, y := counter([]rune(input))

	for i := 'A'; i <= 'E'; i++ {
		quadout, _ = runCommand("./quad"+string(i), PrintNbr(x), PrintNbr(y))
		if quadout == input {
			if count != 0 {
				fmt.Print(" || ")
			}
			fmt.Printf("[quad%v] [%v] [%v]", string(i), PrintNbr(x), PrintNbr(y))
			count++
		}
	}
	if count == 0 {
		fmt.Print("Not a quad function")
	}
	fmt.Println()
}
