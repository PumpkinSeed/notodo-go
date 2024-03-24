package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/PumpkinSeed/notodo-go/internal/storage"
	"github.com/PumpkinSeed/notodo-go/internal/types"
)

func createItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var item types.Item
	if err := bind(ctx, r, &item); err != nil {
		slog.ErrorContext(ctx, "bind error in createItem", err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("The request body is invalid")); err != nil {
			slog.ErrorContext(ctx, "w.Write error in createItem controller", err)
		}
		return
	}

	key := base64.StdEncoding.EncodeToString([]byte(item.Name))
	storage.SetItem(ctx, key, item)
	slog.InfoContext(ctx, "Item successfully created", slog.String("key", key), slog.String("name", item.Name))
	w.WriteHeader(http.StatusCreated)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	items, err := storage.GetItems(ctx)
	if err != nil {
		internalError(ctx, w)
		return
	}

	resp, err := json.Marshal(items)
	if err != nil {
		internalError(ctx, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		slog.ErrorContext(ctx, "w.Write error in getItems controller, resp", err)
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := r.PathValue("key")

	item, err := storage.GetItem(ctx, key)
	if err != nil {
		internalError(ctx, w)
		return
	}

	resp, err := json.Marshal(item)
	if err != nil {
		internalError(ctx, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		slog.ErrorContext(ctx, "w.Write error in getItem controller, resp", err)
	}
}

func internalError(ctx context.Context, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte("Internal error")); err != nil {
		slog.ErrorContext(ctx, "w.Write error in controller, internal error", err)
	}
}
