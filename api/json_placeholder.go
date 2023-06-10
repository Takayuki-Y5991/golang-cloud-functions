package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 必要なパラメータのみ抽出
type Model struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Fetch(customerId int) (Model, error) {
	var generatedURL strings.Builder
	generatedURL.WriteString(os.Getenv("DUMMY_API_URI"))
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
