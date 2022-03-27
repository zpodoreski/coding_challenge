package multiply

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coding_challenge/pkg/model"
)

func TestMultiply(t *testing.T) {

	var tests = []struct {
		x, y, result float64
	}{
		{16, 2, 32},
		{2, 16, 32},
	}
	t.Run("Return two params multiply", func(t *testing.T) {

		for _, test := range tests {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/multiply?x=%f&y=%f", test.x, test.y), nil)
			response := httptest.NewRecorder()

			Multiply(response, request)

			resp := &model.Response{}

			fmt.Println(response)
			err := json.NewDecoder(response.Body).Decode(resp)
			if err != nil {
				t.Fatalf("Error while unmarshaling response")
			}

			if resp.Answer != float64(test.result) {
				t.Errorf("got %f, want %f", resp.Answer, test.result)
			}
		}

	})
}
