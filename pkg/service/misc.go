package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func bind(_ context.Context, r *http.Request, v any) error {
	rawData, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawData, v)
}
