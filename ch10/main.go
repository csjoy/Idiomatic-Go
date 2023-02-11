package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	// "fmt"
	// "fmt"
	// "net/http"
	"time"
)

func main() {
	// ch := make(chan int) // build-in type created with make function
	// a := <-ch            // reads a value from ch and assigns it to a
	// ch <- a              // write the value in a to ch

	// ch1 := make(chan int)
	// var vv <-chan int
	// var vs <-chan string

	// ch := make(chan int, 10) // buffered channel
	// fmt.Println(len(ch), cap(ch))
	// uch := make(chan int) // unbuffered channel
	// fmt.Println(len(uch), len(uch))

	// ch := runs()
	// // for v := range ch {
	// // 	fmt.Println(v)
	// // }
	// v, ok := <-ch
	// if !ok {
	// 	fmt.Println("Channel is empty")
	// } else {
	// 	fmt.Println(v)
	// }

	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go func() {
	// 	v1 := 1
	// 	ch1 <- v1
	// 	v2 := <-ch2
	// 	fmt.Println("Not here", v1, v2)
	// }()
	// v1 := 2
	// var v2 int
	// select {
	// case ch2 <- v1:
	// case v2 = <-ch1:
	// }
	// fmt.Println("Here", v1, v2)
	// select {
	// case v := <-ch1:
	// 	fmt.Println("read from ch:", v)
	// default:
	// 	fmt.Println("no value written to ch")
	// }

	// a := []int{2, 4, 6, 8, 10}
	// ch := make(chan int, len(a))
	// for _, v := range a {
	// 	go func(val int) {
	// 		ch <- val * 2
	// 	}(v)
	// }
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(<-ch)
	// }

	// for i := range countTo(10) {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// ch, cancel := countToCancel(10)
	// for i := range ch {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	fmt.Println()
	// }
	// cancel()

	// pg := New(10)
	// http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
	// 	err := pg.Process(func() {
	// 		w.Write([]byte(doThingThatSouldBeLimited()))
	// 	})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusTooManyRequests)
	// 		w.Write([]byte("Too many requests"))
	// 	}
	// })
	// http.ListenAndServe(":8080", nil)

	// mp := map[int]int{1: 1, 2: 5}
	// x, ok := mp[2]
	// fmt.Println(x, ok)

	// in1 := make(chan struct{})
	// in2 := make(chan struct{})
	// done := make(chan struct{})

	// for {
	// 	select {
	// 	case _, ok := <-in1:
	// 		if !ok {
	// 			in1 = nil // the case will never succeed again!
	// 			continue
	// 		}
	// 		// process the v that was read from in
	// 	case _, ok := <-in2:
	// 		if !ok {
	// 			in2 = nil // the case will never succeed again!
	// 			continue
	// 		}
	// 	case <-done:
	// 		return
	// 	default:
	// 		time.Sleep(2 * time.Second)
	// 		close(done)
	// 	}
	// }

	// var wg sync.WaitGroup
	// wg.Add(3)
	// go func() {
	// 	defer wg.Done()
	// 	doThing1()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	doThing2()
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	doThing3()
	// }()
	// wg.Wait()

	// test := initParser()
	// test.Parse("ehlo")

	// test, close := NewChannelScoreboardManager()
	// defer close()
}

type MutexScoreboardManager struct {
	lock       sync.RWMutex
	scoreboard map[string]int
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.lock.Lock()
	defer msm.lock.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.lock.RLock()
	defer msm.lock.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type AOut struct{}
type BOut struct{}
type COut struct{}

type Input struct {
	A AOut
	B BOut
}

type CIn Input

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2),
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

func getResultA(ctx context.Context, data AOut) (AOut, error) {
	// business logic goes here
	return AOut{}, nil
}
func getResultB(ctx context.Context, data BOut) (BOut, error) {
	// business logic goes here
	return BOut{}, nil
}
func getResultC(ctx context.Context, data CIn) (COut, error) {
	// business logic goes here
	return COut{}, nil
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()
	go func() {
		select {
		case <-ctx.Done():
			return
		case inputC := <-p.inC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0
	for count < 2 {
		select {
		case a := <-p.outA:
			inputC.A = a
			count++
		case b := <-p.outB:
			inputC.B = b
			count++
		case err := <-p.errs:
			return CIn{}, err
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func initParser() SlowComplicatedParser {
	// do all sorts of setup and loading here
	return Demo{}
}

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

type Demo struct{}

func (d Demo) Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func doThing1() {
	// do something
	fmt.Println("doing thing 1...done")
}
func doThing2() {
	// do something
	fmt.Println("doing thing 2...done")
}
func doThing3() {
	// do something
	fmt.Println("doing thing 3...done")
}

func doSomeWork() (int, error) {
	// do your stuff here
	return 0, nil
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}

func doThingThatSouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func processChannel(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func countToCancel(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
			}
		}(searcher)
	}
	r := <-result
	close(done)
	return r
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func runs() <-chan int {
	ch := make(chan int)
	go func() {
		// ch <- 0
		defer close(ch)
	}()
	return ch
}

func process(val int) int {
	// do something
	return val
}

func runThingConcurrently(in <-chan int, out chan<- int) {
	go func() {
		for val := range in {
			result := process(val)
			out <- result
		}
	}()
}
