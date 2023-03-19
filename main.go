package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	endPoint := "https://api.openai.com/v1/chat/completions"
	apiKey := "输入自己的OpenAI-Key"

	client := resty.New()
	//请求报文体
	messages := []map[string]string{{"role": "user", "content": "这是我第一次使用API"}}

	reqBody := make(map[string]interface{})
	reqBody["model"] = "gpt-3.5-turbo"
	reqBody["messages"] = messages
	reqBody["temperature"] = 1

	reqJSON, _ := json.Marshal(reqBody)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey)).
		SetBody(reqJSON).
		Post(endPoint)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp.String())
}
