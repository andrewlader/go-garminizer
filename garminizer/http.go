package garminizer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func makeAPICall(httpVerb string, api string, headers http.Header, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(httpVerb, api, body)
	if err != nil {
		return nil, err
	}

	request.Header = headers

	return doAPIRequest(request)
}

func doAPIRequest(request *http.Request) ([]byte, error) {
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("web error: %v\n", err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if (response.StatusCode != http.StatusOK) && (response.StatusCode != http.StatusCreated) &&
		(response.StatusCode != http.StatusAccepted) {
		if err == nil {
			err = fmt.Errorf("HTTP %d: %s", response.StatusCode, string(data))
			log.Printf("web request error: %v", err)
		}
	}

	return data, err
}
