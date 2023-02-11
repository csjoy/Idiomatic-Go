package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// c := context.Background()
	// ctx := context.TODO()
	// result, err := logic(c, "Hello")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result, c, ctx)

	// ctx := context.Background()
	// parent, cancel := context.WithTimeout(ctx, 4*time.Second)
	// defer cancel()
	// child, cancel2 := context.WithTimeout(parent, 3*time.Second)
	// defer cancel2()
	// start := time.Now()
	// <-child.Done()
	// end := time.Now()
	// fmt.Println(end.Sub(start))

	bl := BusinessLogic{
		RequestDecorator: Request,
		Logger:           Logging{},
		Remote:           "http://www.example.com/query",
	}

}

// // better approach of passing value to context
// func (c Controller) handleRequest(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	user, ok := indentify.UserFromContext(ctx)
// 	if !ok {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	data := r.URL.Query().Get("data")
// 	result, err := c.Logic.businessLogic(ctx, user, data)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.Write([]byte(result))
// }

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = ContextWithUser(ctx, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

type wrapper struct {
	Result string
	Error  error
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	ch := make(chan wrapper, 1)
	go func() {
		// do the long running thing

		result, err := longRunningThing(ctx, data)
		ch <- wrapper{Result: result, Error: err}
	}()
	select {
	case data := <-ch:
		return data.Result, data.Error
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	time.Sleep(1 * time.Minute)
	return data, nil
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	res, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d", res.StatusCode)
	}
	// do the rest of the stuff to process the response
	id, err := processResponse(res.Body)
	return id, err
}

func processResponse(r io.ReadCloser) (string, error) {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String(), nil
}

func logic(ctx context.Context, info string) (string, error) {
	// do some interesting stuff here
	return info, nil
}

func _Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// wrap the context with stuff -- we'll see how soon!
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data := r.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(result))
}
