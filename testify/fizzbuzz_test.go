package fizzbuzz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFizzbuzzFizz(t *testing.T) {
	assert.Equal(t, "fizz", Fizzbuzz(3))
}

func TestFizzbuzzBuzz(t *testing.T) {
	assert.Equal(t, "buzz", Fizzbuzz(5))
}

func TestFizzbuzz(t *testing.T) {
	type teststruct struct {
		input    int
		expected string
	}
	tests := []teststruct{
		{0, "fizzbuzz"},
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{5, "buzz"},
		{6, "fizz"},
		{10, "buzz"},
		{15, "fizzbuzz"},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%d should give %s", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, Fizzbuzz(test.input))
		})
	}
}

func TestFizzbuzzApi(t *testing.T) {
	type teststruct struct {
		input    int
		expected string
	}
	tests := []teststruct{
		{0, "fizzbuzz"},
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{5, "buzz"},
		{6, "fizz"},
		{10, "buzz"},
		{15, "fizzbuzz"},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%d should give %s", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest("GET", fmt.Sprintf("/fizzbuzz?input=%d", test.input), nil)
			w := httptest.NewRecorder()

			FizzbuzzHandler(w, req)

			assert.Equal(t, 201, w.Code)
			body, err := ioutil.ReadAll(w.Body)
			require.NoError(t, err)
			assert.Equal(t, test.expected, string(body))
		})
	}
}
func TestFizzbuzzApiBadInput(t *testing.T) {

	req := httptest.NewRequest("GET", "/fizzbuzz?input=foobar", nil)
	w := httptest.NewRecorder()

	FizzbuzzHandler(w, req)

	assert.Equal(t, 400, w.Code)

}

func TestFizzbuzzClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "3", r.URL.Query().Get("input"))
		input, _ := strconv.Atoi(r.URL.Query().Get("input"))
		fmt.Fprint(w, Fizzbuzz(input))
	}))
	defer ts.Close()

	client := NewClient(ts.URL)
	response, err := client.Fizzbuzz(3)
	require.NoError(t, err)
	assert.Equal(t, "fizz", response)
}
