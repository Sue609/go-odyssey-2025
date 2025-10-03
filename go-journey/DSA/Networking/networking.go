package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Message struct {
	Name string
	Content string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Step one: Get user input
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) //For removing mewline

	fmt.Print("Enter your message: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	//Step two: Create struct for input
	original := Message{Name: name, Content: content}
	fmt.Println("\nOriginal Struct", original)

	//Step three: Encode to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}
	fmt.Println("Encoded (JSON):", string(jsonData))

	//step Four: Decode back into struct
	var decoded Message
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil{
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded Struct:", decoded)
}