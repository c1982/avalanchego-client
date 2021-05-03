package api

import "time"

type GetLivenessResponse struct {
	Checks struct {
		C struct {
			Message interface{} `json:"message"`
			Error   struct {
				Message string `json:"message"`
			} `json:"error"`
			Timestamp          time.Time `json:"timestamp"`
			Duration           int       `json:"duration"`
			Contiguousfailures int       `json:"contiguousFailures"`
			Timeoffirstfailure time.Time `json:"timeOfFirstFailure"`
		} `json:"C"`
		P struct {
			Message struct {
				Percentconnected float64 `json:"percentConnected"`
			} `json:"message"`
			Timestamp          time.Time   `json:"timestamp"`
			Duration           int         `json:"duration"`
			Contiguousfailures int         `json:"contiguousFailures"`
			Timeoffirstfailure interface{} `json:"timeOfFirstFailure"`
		} `json:"P"`
		X struct {
			Timestamp          time.Time   `json:"timestamp"`
			Duration           int         `json:"duration"`
			Contiguousfailures int         `json:"contiguousFailures"`
			Timeoffirstfailure interface{} `json:"timeOfFirstFailure"`
		} `json:"X"`
		ChainsDefaultBootstrapped struct {
			Timestamp          time.Time   `json:"timestamp"`
			Duration           int         `json:"duration"`
			Contiguousfailures int         `json:"contiguousFailures"`
			Timeoffirstfailure interface{} `json:"timeOfFirstFailure"`
		} `json:"chains.default.bootstrapped"`
		NetworkValidatorsHeartbeat struct {
			Message struct {
				Heartbeat int `json:"heartbeat"`
			} `json:"message"`
			Timestamp          time.Time   `json:"timestamp"`
			Duration           int         `json:"duration"`
			Contiguousfailures int         `json:"contiguousFailures"`
			Timeoffirstfailure interface{} `json:"timeOfFirstFailure"`
		} `json:"network.validators.heartbeat"`
	} `json:"checks"`
	Healthy bool `json:"healthy"`
}
