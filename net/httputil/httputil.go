package httputil

import (
	"io/ioutil"
	"net/http"
)

func Response(response *http.Response, e error) (string, error) {

	if e != nil {
		return "", e
	}

	defer response.Body.Close()

	data, e := ioutil.ReadAll(response.Body)

	return string(data), e
}
