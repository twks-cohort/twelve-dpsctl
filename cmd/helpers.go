package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpHandler(r *http.Request, s *StatusCodes) ([]byte, error) {
	r.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, transientError{err: err}
	}

	defer response.Body.Close()

	if s != nil {
		for _, customError := range s.statusErrors {
			if response.StatusCode == customError.code {
				return nil, fmt.Errorf("%s", customError.errMessage)
			}
		}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", response.Status)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, transientError{err: err}
	}

	return responseBytes, nil
}
