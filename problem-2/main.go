package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {

	var dump = make(map[string]string)

	//Open file and read line by line
	file, err := os.Open("passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	l := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		val, ok := dump[split[0]]
		if ok == true {
			fmt.Printf("***KEY EXISTS*** KEY: %v VALUE: %v", split[0], split[1])
		}
		dump[split[0]] = split[1]
		l++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("VALID PASSWORDS: %i\n", parsePolicy(dump))
	fmt.Printf("SCAN COUNT: %v\n", l)
	fmt.Printf("DUMP LENGTH: %v\n", len(dump))

}

func parsePolicy(raw map[string]string) int {
	c := 0

	for k, v := range raw {
		policy := strings.Split(k, " ")
		bound := strings.Split(policy[0], "-")
		count := strings.Count(v, policy[1])

		min, err := strconv.Atoi(bound[0])
		if err != nil {
			log.Fatal(err)
		}

		max, err := strconv.Atoi(bound[1])
		if err != nil {
			log.Fatal(err)
		}

		if (count >= min) && (count <= max) {
			c++
			fmt.Printf("***PASS*** KEY: %v  VALUE: %v\n", k, v)
		} else {
			fmt.Printf("***FAIL*** KEY: %v  VALUE: %v MIN: %v MAX: %v COUNT: %v\n", k, v, min, max, count)
		}
	}	

	return c
}

