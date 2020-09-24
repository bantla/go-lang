package errors

type APIError struct {
	Status int `json:"-"`

	Code string `json:"code,omitempty"`

	Message string `json:"message,omitempty"`

	Err error `json:"-"`

	TraceID string `json:"traceId,omitempty"`

	Errors interface{} `json:"errors,omitempty"`
}

func (err APIError) Error() string {
	return err.Message
}

func (err APIError) Unwrap() error {
	return err.Err
}
