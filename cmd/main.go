package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"

	_ "konkon-t.com/sendmail-function"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failure godotenv.Load: %v", err)
	}
}

func main() {
	if err := funcframework.Start(os.Getenv("PORT")); err != nil {
		log.Fatalf("funcframework.Start: %v", err)
	}
}
