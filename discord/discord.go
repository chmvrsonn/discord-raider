package discord

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SendMessage(content string, tts bool, authorization string, channelId string) int {
	url := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", channelId)

	payload := []byte(fmt.Sprintf(`{"content": "%s", "nonce": "%s","tts": %s}`, content, strconv.FormatInt(Snowflake(), 10), strings.ToLower(strconv.FormatBool(tts))))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return 0
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.5")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Add("x-debug-options", "bugReporterEnabled")
	req.Header.Add("x-discord-locale", "en-US")
	req.Header.Add("x-discord-timezone", "Etc/GMT-1")
	req.Header.Add("authorization", authorization)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	defer res.Body.Close()

	return res.StatusCode
}

// https://github.com/V4NSH4J/discord-mass-DM-GO/blob/main/utilities/misc.go#L23
func Snowflake() int64 {
	snowflake := strconv.FormatInt((time.Now().UTC().UnixNano()/1000000)-1420070400000, 2) + "0000000000000000000000"
	nonce, _ := strconv.ParseInt(snowflake, 2, 64)
	return nonce
}
