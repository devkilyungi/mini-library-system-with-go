package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserChoice() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your choice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

func GetBookTitle() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	return ToTitle(title)
}

func ToTitle(input string) string {
	words := strings.Fields(strings.ToLower(input))
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}
