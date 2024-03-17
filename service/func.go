package service

import (
	"io"
	"net/http"
)

func deferResponseClose(s *http.Response) {
	if s != nil {
		defer s.Body.Close()
	}
}

//handleHTTPResponse handle
func handleHTTPResponse(res *http.Response, err error) (io.ReadCloser, error) {
	if err != nil {
		deferResponseClose(res)
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, ErrUnexpectedStatusCode(res.StatusCode)
	}

	return res.Body, nil
}
