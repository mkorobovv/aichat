package controller

type responseError struct {
	Kind   string `json:"kind"`
	Detail string `json:"detail"`
	status int
}

func (err responseError) Error() string {
	return err.Detail
}
