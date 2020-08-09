package garminizer

type pelotonLoginRequestBody struct {
	Username string `json:"username_or_email"`
	Password string `json:"password"`
}

type pelotonAuthResponse struct {
	PubsubSession struct{} `json:"pubsub_session"`
	SessionID     string   `json:"session_id"`
	UserData      struct {
		Birthday         int64    `json:"birthday"`
		BlockExplicit    bool     `json:"block_explicit"`
		CreatedAt        int64    `json:"created_at"`
		CreatedCountry   string   `json:"created_country"`
		Email            string   `json:"email"`
		Username         string   `json:"username"`
		FirstName        string   `json:"first_name"`
		MiddleInitial    string   `json:"middle_initial"`
		LastName         string   `json:"last_name"`
		Location         string   `json:"location"`
		Name             string   `json:"name"`
		Gender           string   `json:"gender"`
		Height           float64  `json:"height"`
		Weight           float64  `json:"weight"`
		MemberGroups     []string `json:"member_groups"`
		ID               string   `json:"id"`
		IsProfilePrivate bool     `json:"is_profile_private"`
		ObfuscatedEmail  string   `json:"obfuscated_email"`
		PairedDevices    []struct {
			Name             string `json:"name"`
			PairedDeviceType string `json:"paired_device_type"`
			SerialNumber     string `json:"serial_number"`
		} `json:"paired_devices"`
		PhoneNumber   interface{} `json:"phone_number"`
		TotalWorkouts int64       `json:"total_workouts"`
		WorkoutCounts []struct {
			Count   int64  `json:"count"`
			IconURL string `json:"icon_url"`
			Name    string `json:"name"`
			Slug    string `json:"slug"`
		} `json:"workout_counts"`
	} `json:"user_data"`
	UserID string `json:"user_id"`
}

type pelotonAuthFullResponse struct {
	PubsubSession struct{} `json:"pubsub_session"`
	SessionID     string   `json:"session_id"`
	UserData      struct {
		Birthday           int64 `json:"birthday"`
		BlockExplicit      bool  `json:"block_explicit"`
		CanCharge          bool  `json:"can_charge"`
		ContractAgreements []struct {
			AgreedAt            int64  `json:"agreed_at"`
			BikeContractURL     string `json:"bike_contract_url"`
			ContractCreatedAt   int64  `json:"contract_created_at"`
			ContractDisplayName string `json:"contract_display_name"`
			ContractID          string `json:"contract_id"`
			ContractType        string `json:"contract_type"`
			TreadContractURL    string `json:"tread_contract_url"`
		} `json:"contract_agreements"`
		CreatedAt                    int64         `json:"created_at"`
		CreatedCountry               string        `json:"created_country"`
		CustomizedHeartRateZones     []interface{} `json:"customized_heart_rate_zones"`
		CustomizedMaxHeartRate       int64         `json:"customized_max_heart_rate"`
		CyclingFtp                   int64         `json:"cycling_ftp"`
		CyclingFtpSource             interface{}   `json:"cycling_ftp_source"`
		CyclingFtpWorkoutID          interface{}   `json:"cycling_ftp_workout_id"`
		CyclingWorkoutFtp            int64         `json:"cycling_workout_ftp"`
		DefaultHeartRateZones        []int64       `json:"default_heart_rate_zones"`
		DefaultMaxHeartRate          int64         `json:"default_max_heart_rate"`
		Email                        string        `json:"email"`
		EstimatedCyclingFtp          int64         `json:"estimated_cycling_ftp"`
		FacebookAccessToken          interface{}   `json:"facebook_access_token"`
		FacebookID                   interface{}   `json:"facebook_id"`
		FirstName                    string        `json:"first_name"`
		Gender                       string        `json:"gender"`
		HasActiveDeviceSubscription  bool          `json:"has_active_device_subscription"`
		HasActiveDigitalSubscription bool          `json:"has_active_digital_subscription"`
		HasSignedWaiver              bool          `json:"has_signed_waiver"`
		Height                       int64         `json:"height"`
		ID                           string        `json:"id"`
		ImageURL                     string        `json:"image_url"`
		InstructorID                 interface{}   `json:"instructor_id"`
		IsCompleteProfile            bool          `json:"is_complete_profile"`
		IsDemo                       bool          `json:"is_demo"`
		IsExternalBetaTester         bool          `json:"is_external_beta_tester"`
		IsFitbitAuthenticated        bool          `json:"is_fitbit_authenticated"`
		IsInternalBetaTester         bool          `json:"is_internal_beta_tester"`
		IsProfilePrivate             bool          `json:"is_profile_private"`
		IsProvisional                bool          `json:"is_provisional"`
		IsStravaAuthenticated        bool          `json:"is_strava_authenticated"`
		LastName                     string        `json:"last_name"`
		LastWorkoutAt                int64         `json:"last_workout_at"`
		Location                     string        `json:"location"`
		MemberGroups                 []string      `json:"member_groups"`
		MiddleInitial                string        `json:"middle_initial"`
		Name                         string        `json:"name"`
		ObfuscatedEmail              string        `json:"obfuscated_email"`
		PairedDevices                []struct {
			Name             string `json:"name"`
			PairedDeviceType string `json:"paired_device_type"`
			SerialNumber     string `json:"serial_number"`
		} `json:"paired_devices"`
		PhoneNumber interface{} `json:"phone_number"`
		QuickHits   struct {
			InclineShortcuts interface{} `json:"incline_shortcuts"`
			QuickHitsEnabled bool        `json:"quick_hits_enabled"`
			SpeedShortcuts   interface{} `json:"speed_shortcuts"`
		} `json:"quick_hits"`
		ReferralsMade                  int64  `json:"referrals_made"`
		TotalFollowers                 int64  `json:"total_followers"`
		TotalFollowing                 int64  `json:"total_following"`
		TotalNonPedalingMetricWorkouts int64  `json:"total_non_pedaling_metric_workouts"`
		TotalPedalingMetricWorkouts    int64  `json:"total_pedaling_metric_workouts"`
		TotalPendingFollowers          int64  `json:"total_pending_followers"`
		TotalWorkouts                  int64  `json:"total_workouts"`
		Username                       string `json:"username"`
		V1ReferralsMade                int64  `json:"v1_referrals_made"`
		Weight                         int64  `json:"weight"`
		WorkoutCounts                  []struct {
			Count   int64  `json:"count"`
			IconURL string `json:"icon_url"`
			Name    string `json:"name"`
			Slug    string `json:"slug"`
		} `json:"workout_counts"`
	} `json:"user_data"`
	UserID string `json:"user_id"`
}
