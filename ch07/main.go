package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	// "reflect"
	"time"
)

// Person type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Counter type
type Counter struct {
	total      int
	lastUpdate time.Time
}

// Increment method
func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdate)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}
func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

// IntTree type
type IntTree struct {
	val         int
	left, right *IntTree
}

// Insert method
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// Contains method
func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// Adder type
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type Score int
type HightScore Score

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

type BitField int

const (
	Field1 BitField = 1 << iota
	Field2
	Field3
	Field4
)

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	return m.Reports
}

// type Inner struct {
// 	X int
// }

// type Outer struct {
// 	Inner
// 	X int
// }

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return data
}

type LogicB interface {
	Process(data string) string
}

type Client struct {
	L LogicB
}

func (c Client) Program() {
	// get data from somewhere
	c.L.Process("Hello")
}

func process(r io.Reader) error {

	data := make([]byte, 2048)
	for {
		count, err := r.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type ReaderCloser interface {
	Reader
	Closer
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

type MyInt int
type MyInt2 int

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is interface{}
	case int:
		// j is of type int
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// j is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}
		fmt.Println(j)
	}
}

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
type Logger interface {
	Log(message string)
}
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	// p := Person{
	// 	FirstName: "Prosenjit",
	// 	LastName:  "Majumder",
	// 	Age:       26,
	// }
	// fmt.Println(p.String())
	// var c Counter
	// fmt.Println((&c).String())
	// (&c).Increment()
	// fmt.Println(c.String())

	// var c Counter
	// doUpdateWrong(c)
	// fmt.Println("in main:", c.String())
	// doUpdateRight(&c)
	// fmt.Println("in main:", c.String())

	// var it *IntTree
	// it = it.Insert(5)
	// it = it.Insert(3)
	// it = it.Insert(10)
	// it = it.Insert(2)
	// fmt.Println(it.Contains(2))
	// fmt.Println(it.Contains(12))

	// myAdder := Adder{start: 10}
	// fmt.Println(myAdder.AddTo(5))
	// f1 := myAdder.AddTo
	// fmt.Println(f1(5))
	// f2 := Adder.AddTo
	// fmt.Println(f2(myAdder, 5))

	// var i int = 300
	// var s Score = 100
	// var hs HightScore = 200
	// hs = HightScore(s)
	// s = Score(i)
	// fmt.Println(hs, s)
	// fmt.Println(reflect.TypeOf(hs), reflect.TypeOf(s))

	// fmt.Println(Uncategorized)
	// fmt.Println(reflect.TypeOf(Uncategorized))

	// fmt.Println(Field1)
	// fmt.Println(reflect.TypeOf(Field1))

	// m := Manager{
	// 	Employee: Employee{
	// 		Name: "Joy",
	// 		ID:   "1234",
	// 	},
	// 	Reports: []Employee{
	// 		{Name: "aj", ID: "12"},
	// 		{Name: "cs", ID: "43"},
	// 		{Name: "kl", ID: "23"},
	// 	},
	// }
	// fmt.Println(m.Name, m.ID)
	// fmt.Println(m.Reports)

	// o := Outer{
	// 	Inner: Inner{
	// 		X: 10,
	// 	},
	// 	X: 20,
	// }
	// fmt.Println(o.X)
	// fmt.Println(o.Inner.X)

	// // var eFail Employee = m
	// var eOK Employee = m.Employee
	// fmt.Println(eOK)

	// o := Outer{
	// 	Inner: Inner{
	// 		A: 10,
	// 	},
	// 	S: "Hello",
	// }
	// fmt.Println(o.Double())

	// c := Client{
	// 	L: LogicProvider{},
	// }
	// c.Program()

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// err = process(file)
	// if err != nil {
	// 	panic(err)
	// }

	// var s *string
	// fmt.Println(s == nil)
	// var i interface{}
	// fmt.Println(i == nil)
	// i = s
	// fmt.Println(i == nil)

	// var i interface{}
	// i = 20
	// fmt.Println(i)
	// fmt.Println(reflect.TypeOf(i))
	// i = "hello"
	// fmt.Println(i)
	// fmt.Println(reflect.TypeOf(i))
	// i = struct {
	// 	FirstName string
	// 	LastName  string
	// }{"Fred", "Fredson"}
	// fmt.Println(i)
	// fmt.Println(reflect.TypeOf(i))

	// data := map[string]interface{}{}
	// contents, err := ioutil.ReadFile("generated.json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(contents)
	// json.Unmarshal(contents, &data)
	// for k := range data {
	// 	fmt.Println(k)
	// }

	// var i interface{}
	// var mine MyInt = 20
	// // var mine2 MyInt2 = 10
	// i = mine
	// i2, ok := i.(MyInt)
	// if !ok {
	// 	result := fmt.Errorf("unexpected type for %v", i)
	// 	fmt.Println(result)
	// }
	// fmt.Println(i2 + 1)

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
