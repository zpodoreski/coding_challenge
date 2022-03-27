package helper

import (
	"net/url"

	"github.com/coding_challenge/pkg/model"
	"github.com/gorilla/schema"
)

func ExtractQueryParams(query url.Values) (*model.Response, error) {
	resp := &model.Response{}

	decoder := schema.NewDecoder()

	if err := decoder.Decode(resp, query); err != nil {
		return nil, err
	}

	return resp, nil
}
