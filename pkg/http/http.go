package http

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type Msg map[string]string

func JSON(w http.ResponseWriter, payload interface{}, code int) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error while marshalling the response"))
	}
	w.WriteHeader(code)
	w.Write(response)
}

func SetUpHandler(w http.ResponseWriter, ctx context.Context) (*zerolog.Logger, context.Context, context.CancelFunc) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	l := zerolog.Ctx(ctx)
	return l, ctx, cancel
}

func setCookie(w http.ResponseWriter, name string, token string, expiresAt time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiresAt,
		HttpOnly: true,
		Secure:   true,
	})
}
