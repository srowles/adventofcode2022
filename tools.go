package aoc2022

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

type Coord struct {
	X, Y int
}

func (c *Coord) Distance(b Coord) int {
	return Abs(c.X-b.X) + Abs(c.Y-b.Y)
}

func (c *Coord) Move(direction Coord) Coord {
	return Coord{
		X: c.X + direction.X,
		Y: c.Y + direction.Y,
	}
}

func (c *Coord) Diff(other Coord) Coord {
	return Coord{
		X: c.X - other.X,
		Y: c.Y - other.Y,
	}
}

// Abs returns the absolute (non negative) value of the input
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum value in the input slice
func Min[T constraints.Ordered](in ...T) T {
	if len(in) == 1 {
		return in[0]
	}

	min := in[0]
	for _, v := range in[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// Max returns the maximum value in the input slice
func Max[T constraints.Ordered](in ...T) T {
	if len(in) == 1 {
		return in[0]
	}

	min := in[0]
	for _, v := range in[1:] {
		if v > min {
			min = v
		}
	}

	return min
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD[T constraints.Signed | constraints.Unsigned](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM[T constraints.Signed | constraints.Unsigned](integers ...T) T {
	if len(integers) < 2 {
		return integers[1]
	}
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Count returns an int of the number if values in the input
// slice that return true when evaluated by countFn
func Count[T any](in []T, countFn func(T) bool) int {
	out := 0
	for _, v := range in {
		if countFn(v) {
			out++
		}
	}

	return out
}

// Slice creates a generic slice from the split input data
func Slice[T any](data string, separator string, convert func(string) T) []T {
	lines := strings.Split(strings.TrimSpace(data), separator)
	var ret []T
	for _, line := range lines {
		line = strings.TrimSpace(line)
		ret = append(ret, convert(line))
	}

	return ret
}

func Int(value string) int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func UInt(value string) uint64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint64(i)
}

// InputFromWebsite reads the AOC_SESSION env variable
// to form the session cookie and then reads the input for
// the appropriate day. NB says are not zero prefixed so day
// one is just "1"
//
// TODO oauth login via github etc. so I don't have to steal
// the session cookie from my browser
func InputFromWebsite(day string) string {
	session := strings.TrimSpace(os.Getenv("AOC_SESSION"))
	if session == "" {
		log.Fatal("AOC_SESSION env var must be ser")
	}
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 5,
		Transport: netTransport,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2022/day/%s/input", day), nil)
	if err != nil {
		log.Fatalf("failed to create new request for day %s: %v", day, err)
	}
	req.Header.Set("cookie", fmt.Sprintf("session=%s", session))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to get input for day %s: %v", day, err)
	}
	defer resp.Body.Close()
	data := readStringData(resp.Body)
	return data
}

func readStringData(reader io.Reader) string {
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read all from reader: %v", err)
	}
	return string(data)
}

type RuneStack []rune

func NewRuneStack() RuneStack {
	return make(RuneStack, 0)
}

func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *RuneStack) Push(runes ...rune) {
	for _, r := range runes {
		*s = append(*s, r)
	}
}

func (s *RuneStack) Pop() rune {
	if s.IsEmpty() {
		return '_'
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *RuneStack) String() string {
	var result string
	for _, r := range *s {
		result = result + string(r)
	}
	return result
}
