package fizzbuzz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Fizzbuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "fizzbuzz"
	}

	if input%5 == 0 {
		return "buzz"
	}

	if input%3 == 0 {
		return "fizz"
	}

	return strconv.Itoa(input)
}

func FizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	input, err := strconv.Atoi(req.URL.Query().Get("input"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", Fizzbuzz(input))
}

type FizzbuzzClient interface {
	Fizzbuzz(input int) (string, error)
}

type FizzbuzzClientImpl struct {
	url string
}

func NewClient(url string) *FizzbuzzClientImpl {
	return &FizzbuzzClientImpl{
		url: url,
	}
}

// resp, err := http.Get("http://example.com/")
func (f FizzbuzzClientImpl) Fizzbuzz(input int) (string, error) {
	requestURL := fmt.Sprintf("%s/fizzbuzz?input=%d", f.url, input)
	resp, err := http.Get(requestURL)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
