package main

import (
	// "bytes"
	"compress/gzip"
	"log"
	// "context"
	// "encoding/json"

	// "encoding/json"
	"net/http"

	// "encoding/json"
	// "fmt"
	"io"
	"os"

	// "strings"
	"time"
	// "strings"
)

func main() {
	// s := "the quick brown fox jumps over the lazy dog"
	// sr := strings.NewReader(s)
	// count, err := countLetters(sr)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(count)

	// r, closer, err := buildGZipReader("my_data.txt.gz")
	// if err != nil {
	// 	panic(err)
	// }
	// defer closer()
	// counts, err := countLetters(r)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(counts)

	// d := 0*time.Hour + 0*time.Minute + 30*time.Second // d is of type time.Duration
	// df, _ := time.ParseDuration("30s")
	// time.Sleep(d)
	// fmt.Println("d stop... df starting", d.Seconds())
	// time.Sleep(df)
	// fmt.Println("df stopped", df.Milliseconds())

	// t, err := time.Parse("2006-01-02 15:04:05 -0700", "2016-03-13 00:00:00 +0000")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(t.Format("January 2, 2006 at 3:05:04PM MST"))
	// fmt.Println(t.Year())

	// var o Order
	// contents, err := ioutil.ReadFile("file.json")
	// if err != nil {
	// 	panic(err)
	// }
	// err = json.Unmarshal(contents, &o)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(o)
	// out, err := json.Marshal(o)
	// fmt.Println(out)

	// toFile := Person{
	// 	Name: "Fred",
	// 	Age:  40,
	// }
	// tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	// if err != nil {
	// 	panic(err)
	// }
	// defer os.Remove(tmpFile.Name())
	// err = json.NewEncoder(tmpFile).Encode(toFile)
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpFile.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// tmpFile2, err := os.Open(tmpFile.Name())
	// if err != nil {
	// 	panic(err)
	// }
	// var fromFile Person
	// err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpFile2.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", fromFile)
	// data := `{"name": "Fred", "age": 40}{"name": "Mary", "age": 21}{"name": "Pat", "age": 30}`
	// dec := json.NewDecoder(strings.NewReader(data))
	// t := Person{}
	// c := 0
	// for dec.More() {
	// 	err := dec.Decode(&t)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(t, c)
	// 	c++
	// }

	// client := &http.Client{
	// 	Timeout: 30 * time.Second,
	// }
	// req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// req.Header.Add("X-My-Client", "Learning Go")
	// res, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != http.StatusOK {
	// 	panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	// }
	// fmt.Println(res.Request.Proto)
	// fmt.Println(res.Header.Get("Content-Type"))
	// var data struct {
	// 	UserID    int    `json:"userId"`
	// 	ID        int    `json:"id"`
	// 	Title     string `json:"title"`
	// 	Completed bool   `json:"completed"`
	// }
	// err = json.NewDecoder(res.Body).Decode(&data)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", data)

	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good puppy!\n"))
	})
	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

	// terribleSecurity := TerribleSecurityProvider("GOPHER")
	// mux.Handle("/hello", terribleSecurity(RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello!\n"))
	// }))))

	// wrappedMux := terribleSecurity(RequestTimer(mux))
	// s := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: wrappedMux,
	// }
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		panic(err)
	}
	*rt = RFC822ZTime{t}
	return nil
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID           string    `json:"id"`
	Date_ordered time.Time `json:"date_ordered"`
	CustomerID   string    `json:"customer_id"`
	Items        []Item    `json:"items"`
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
