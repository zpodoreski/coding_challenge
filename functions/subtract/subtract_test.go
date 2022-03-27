package subtract

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coding_challenge/pkg/model"
)

func TestSubtract(t *testing.T) {

	var tests = []struct {
		x, y, result float64
	}{
		{16, 2, 14},
		{2, 16, -14},
	}
	t.Run("Return two params subtracted", func(t *testing.T) {

		for _, test := range tests {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/subtract?x=%f&y=%f", test.x, test.y), nil)
			response := httptest.NewRecorder()

			Subtract(response, request)

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
