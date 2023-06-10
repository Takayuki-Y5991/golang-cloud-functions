package function

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"konkon-t.com/sendmail-function/api"
)

func init() {
	functions.HTTP("SendGrindFunction", sendGrindFunction)
}

func sendGrindFunction(w http.ResponseWriter, r *http.Request) {
	batch()
}

func processFetch() ([]api.Model, error) {
	var wg sync.WaitGroup
	dataCh := make(chan api.Model, 10)
	errCh := make(chan error, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			data, err := api.Fetch(index)
			if err != nil {
				errCh <- err
				return
			}
			dataCh <- data
		}(i)
	}
	go func() {
		wg.Wait()
		close(dataCh)
		close(errCh)
	}()
	store := make([]api.Model, 0)

	for data := range dataCh {
		store = append(store, data)
	}

	if err := <-errCh; err != nil {
		return nil, err
	}
	return store, nil
}

func transformEmail(email string) string {
	index := strings.LastIndex(email, "@")
	if index == -1 {
		return email
	}
	return email[:index] + "@example.com"
}

func processTransform(res []api.Model) []api.EmailModel {
	data := make([]api.EmailModel, 0, 10)
	for _, v := range res {
		data = append(data, api.EmailModel{
			Username: v.Username,
			Email:    transformEmail(v.Email),
		})
	}
	return data
}

func processSendMail(data []api.EmailModel) error {
	err := api.SendEmail(data)
	if err != nil {
		return err
	}
	return nil
}

func batch() {
	res, err := processFetch()
	if err != nil {
		log.Println(err)
		return
	}
	data := processTransform(res)

	err = processSendMail(data)
	if err != nil {
		log.Println(err)
	}
}
