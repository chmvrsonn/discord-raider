package main

import (
	"fmt"
	"github.com/chmvrsonn/discord-raider/discord"
	"github.com/chmvrsonn/discord-raider/utils"
	"github.com/common-nighthawk/go-figure"
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

	config, _ := utils.LoadConfig()

	err := utils.ValidateConfig(config)

	if err != nil {
		fmt.Println(fmt.Sprintf("Config error: %s", err.Error()))
		return
	}

	delayDuration := time.Duration(config.Delay) * time.Millisecond

	tokens := utils.ReadFileToArray("tokens.txt")

	if tokens == nil {
		return
	}

	for {
		for index, token := range tokens {
			response := discord.SendMessage(config.Message, config.TTS, token, config.ChannelID)

			if response >= 200 && response <= 299 {
				fmt.Println(fmt.Sprintf("-> Message sent! (%s)", strconv.Itoa(index)))
			} else {
				fmt.Println(fmt.Sprintf("-> Unknown error (%s) occurred! (%s)", strconv.Itoa(response), strconv.Itoa(index)))
			}

			time.Sleep(delayDuration)
		}
	}
}
