package solver

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	var io info
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expression := r.URL.Query().Get("expression")
			if expression != io.expression {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "expected expression `%s`, got `%s`", io.expression, expression)
				return
			}
			w.WriteHeader(io.code)
			w.Write([]byte(io.body))
		}))
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}

	data := []struct {
		name   string
		io     info
		result float64
		errMsg string
	}{
		{"case1", info{"2+2*10", http.StatusOK, "22"}, 22, ""},
		// remaining cases
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
