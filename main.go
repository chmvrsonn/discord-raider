package main

import (
	"bufio"
	"fmt"
	"github.com/chmvrsonn/discord-raider/discord"
	"github.com/chmvrsonn/discord-raider/utils"
	"github.com/common-nighthawk/go-figure"
	"os"
	"strconv"
	"time"
)

func main() {
	myFigure := figure.NewFigure("Hakerka Jarocin", "", true)
	myFigure.Print()

	if !utils.FileExists("tokens.txt") {
		utils.CreateFile("tokens.txt")
		fmt.Println("Add tokens to \"tokens.txt\" file and start the program again.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter message: ")
	scanner.Scan()
	messageContent := scanner.Text()

	fmt.Print("Enter channel ID: ")
	scanner.Scan()
	channelId := scanner.Text()

	fmt.Print("Enter delay (in milliseconds): ")
	scanner.Scan()
	delayStr := scanner.Text()
	delay, err := strconv.Atoi(delayStr)

	if err != nil {
		fmt.Println("Error converting delay to integer:", err)
		return
	}

	delayDuration := time.Duration(delay) * time.Millisecond

	tokens := utils.ReadFile("tokens.txt")

	if tokens == nil {
		return
	}

	for {
		for index, token := range tokens {
			discord.SendMessage(messageContent, false, token, channelId)
			fmt.Println(fmt.Sprintf("-> Message sent! (%s)", strconv.Itoa(index)))
			time.Sleep(delayDuration)
		}
	}
}
