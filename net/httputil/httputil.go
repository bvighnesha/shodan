package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Response(response *http.Response, e error) (string, error) {
	defer response.Body.Close()

	if e != nil {
		fmt.Println(e)
		return "", e
	}

	data, e := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	return string(data), e
}

func ErrorFromResponse(response *http.Response) {

}
