package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type guidKey int

const key1 guidKey = 1

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key1, guid)
}
func guidFromContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(key1).(string)
	return g, ok
}

func Middleware1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if guid := r.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

type Logging struct{}

func (Logging) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	// do logging
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFromContext(ctx); ok {
		req.Header.Add("X-GUID", guid)
	}
	return req
}

type Logger interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (bl BusinessLogic) businessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, "starting businessLogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	req = bl.RequestDecorator(req)
	// resp, err := http.DefaultClient.Do(req)
	// // processing continues
	return "", nil
}
