package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile() {
	fileName := "dvhcvn.csv"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mymap map[string]int = map[string]int{}

	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), ",")
		line := strings[4] + "," + strings[5] + "," + strings[3]
		mymap[line] = 1
		fmt.Println(line)
	}

	f, _ := os.OpenFile("ward.txt", os.O_WRONLY, 0777)
	for key := range mymap {
		fmt.Fprintln(f, key)
	}
}

func main() {
	ReadFile()
}
