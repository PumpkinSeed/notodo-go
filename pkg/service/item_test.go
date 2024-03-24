package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PumpkinSeed/notodo-go/internal/storage"
	"github.com/PumpkinSeed/notodo-go/internal/types"
)

func TestCreateItem(t *testing.T) {
	srv := httptest.NewServer(routes())

	data, err := json.Marshal(types.Item{
		Name:        "test",
		Description: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", srv.URL+"/items", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status should be %d, instead of %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestCreateItemError(t *testing.T) {
	srv := httptest.NewServer(routes())

	req, err := http.NewRequest("POST", srv.URL+"/items", bytes.NewReader([]byte("test")))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Status should be %d, instead of %d", http.StatusCreated, resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var expected = "The request body is invalid"
	if string(respBody) != expected {
		t.Errorf("Response body should be '%s', instead of '%s'", expected, respBody)
	}
}

func TestGetItems(t *testing.T) {
	ctx := context.Background()

	storage.SetItem(ctx, "test2", types.Item{
		Name:        "test2",
		Description: "Test 2",
	})
	storage.SetItem(ctx, "test3", types.Item{
		Name:        "test3",
		Description: "Test 3",
	})

	srv := httptest.NewServer(routes())

	req, err := http.NewRequest("GET", srv.URL+"/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, instead of %d", http.StatusCreated, resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(respBody))
}

func TestGetItem(t *testing.T) {
	ctx := context.Background()

	storage.SetItem(ctx, "test2", types.Item{
		Name:        "test2",
		Description: "Test 2",
	})

	srv := httptest.NewServer(routes())

	req, err := http.NewRequest("GET", srv.URL+"/items/test2", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be %d, instead of %d", http.StatusCreated, resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(respBody))
}
