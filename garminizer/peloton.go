package garminizer

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// Peloton is the main struct use to interact with the Peloton API
type Peloton struct {
	auth *pelotonAuthResponse
}

// Login is used to login to the Peloton API
func (peloton *Peloton) Login() error {
	const api = "https://api.pelotoncycle.com/auth/login"
	var pelotonAuthResponse pelotonAuthResponse

	pelotonLoginRequestBody := &pelotonLoginRequestBody{
		Username: viper.GetString("peloton.username"),
		Password: viper.GetString("peloton.password"),
	}

	pelotonLoginJSON, err := json.Marshal(pelotonLoginRequestBody)
	if err != nil {
		return err
	}

	// call Peloton's login API with the right username & password
	pelotonLoginBody := bytes.NewReader([]byte(pelotonLoginJSON))
	bytesResponse, err := makeAPICall(http.MethodPost, api, nil, pelotonLoginBody)
	if err != nil {
		return err
	}

	// process the response from Peloton's login API
	err = json.Unmarshal(bytesResponse, &pelotonAuthResponse)
	if err != nil {
		return err
	}

	peloton.auth = &pelotonAuthResponse
	log.Printf("Peloton's Login response:\n=====\n%v\n", peloton.auth)

	return nil
}

func (peloton *Peloton) downloadWorkouts(numOfWorkoutsToGet int) {

}

func (peloton *Peloton) downloadWorkout() {

}
