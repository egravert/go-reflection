package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	values := make([]int, 10000000, 10000000)
	for i := 0; i < len(values); i++ {
		values[i] = i + 1
	}
	// double no reflection
	timeRun(func() { _ = Map(func(i int) int { return i * 2 }, values) })

	// double with reflection
	timeRun(func() { _ = Map2(func(i int) int { return i * 2 }, values).([]int) })

}

func Map(f func(int) int, in []int) []int {
	results := make([]int, len(in), len(in))
	for i := 0; i < len(in); i++ {
		results[i] = f(in[i])
	}
	return results
}

func Map2(f interface{}, in interface{}) interface{} {
	// get values of the func and incoming slice
	fVal := reflect.ValueOf(f)
	inVal := reflect.ValueOf(in)

	qty := inVal.Len()
	// create slice using reflection (necessary since we dont know type till runtime
	sliceType := reflect.SliceOf(fVal.Type().Out(0))
	results := reflect.MakeSlice(sliceType, qty, qty)

	// iterate the input slice calling f() on each element
	for i := 0; i < qty; i++ {
		v := []reflect.Value{inVal.Index(i)}
		r := fVal.Call(v)[0]
		results.Index(i).Set(r)
	}

	// to accomodate type converstion, Value must be returned as interface{}
	return results.Interface()
}

func timeRun(f func()) {
	t0 := time.Now()
	f()
	fmt.Println("Time spent: ", time.Duration(time.Since(t0)))
}
