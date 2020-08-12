package garminizer

import "log"

// Run kicks off the process of pulling workouts from Peloton and uploading them to Garmin
func Run(source string) {
	peloton := &peloton{}

	// 1) first login to Peloton
	err := peloton.login()
	if err != nil {
		log.Printf("error logging into %s: %s", source, err)
	}

	// 2) then pull down the latest workouts

	// 3) convert them to Garmin's format

	// 4) login to Garmin

	// 5) upload them to Garmin
}
