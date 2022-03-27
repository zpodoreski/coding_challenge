package add

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coding_challenge/pkg/model"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		x, y, result float64
	}{
		{1, 1, 2},
		{2, 3, 5},
		{3, 2, 5},
	}
	t.Run("Return two params added", func(t *testing.T) {

		for _, test := range tests {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/add?x=%v&y=%v", test.x, test.y), nil)
			response := httptest.NewRecorder()

			Add(response, request)

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
