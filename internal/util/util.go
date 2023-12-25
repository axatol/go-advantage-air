package util

import (
	"net/http"
)

func SetRecursively(dest map[string]any, value any, keys ...string) map[string]any {
	if len(keys) < 1 {
		return dest
	}

	if len(keys) == 1 {
		dest[keys[0]] = value
		return dest
	}

	child, ok := dest[keys[0]]
	if !ok {
		child = make(map[string]any)
	}

	sub, ok := child.(map[string]any)
	if !ok {
		return dest
	}

	dest[keys[0]] = SetRecursively(sub, value, keys[1:]...)
	return dest
}

func DefaultRetryValidator(res *http.Response) bool {
	return res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices
}

type RetryRoundTripper struct {
	roundTripper http.RoundTripper
	retryCount   int
	validator    func(*http.Response) bool
}

func NewRetryRoundTripper(roundTripper http.RoundTripper, retryCount int, validator func(*http.Response) bool) http.RoundTripper {
	if retryCount < 1 {
		retryCount = 1
	}

	return &RetryRoundTripper{roundTripper, retryCount, validator}
}

func (r *RetryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	for i := 0; i < r.retryCount; i++ {
		res, err = r.roundTripper.RoundTrip(req)
		if err != nil {
			continue
		}

		if r.validator != nil && r.validator(res) {
			break
		}
	}

	return res, err
}
