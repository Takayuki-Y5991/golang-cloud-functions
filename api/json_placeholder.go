package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 必要なパラメータのみ抽出
type Model struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

const URL = "https://jsonplaceholder.typicode.com/users/"

func Fetch(customerId int) (Model, error) {
	var generatedURL strings.Builder
	generatedURL.WriteString(URL)
	generatedURL.WriteString(strconv.Itoa(customerId))
	resp, err := http.Get(generatedURL.String())
	if err != nil {
		log.Printf("failure cause: %s", err.Error())
		return Model{}, err
	}
	defer resp.Body.Close()

	var response Model
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Model{}, err
	}
	return response, nil
}
