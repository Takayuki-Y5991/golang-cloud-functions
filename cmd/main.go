package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	_ "konkon-t.com/sendmail-function"
)

func main() {
	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v", err)
	}
}
