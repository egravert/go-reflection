package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type Context struct {
	SessionId int
}

func main() {
  // kick off the routed request
	var resp http.ResponseWriter = nil
	req := &http.Request{}
	route(resp, req)
}

func route(resp http.ResponseWriter, req *http.Request) {
	// pretend we are finding the handler to execute
	h := lookupHandler("/")

	// create a type to value mapping (hold this in a mapping somewhere else in a real example)
	typ := reflect.TypeOf(h)
	params := valueList(typ, resp, req, &Context{42})

	// call our handler passing in our value list
	val := reflect.ValueOf(h)
	val.Call(params)
}

func lookupHandler(path string) interface{} {
	return handler
}

func handler(context *Context, resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling request. SessionId is: ", context.SessionId)
}

func valueList(t reflect.Type, params ...interface{}) []reflect.Value {
	var ok bool
	vtMap := make(map[reflect.Type]reflect.Value)
	// map Value to Type
	for _, p := range params {
		t := reflect.TypeOf(p)
		v := reflect.ValueOf(p)
		vtMap[t] = v
	}

	inputs := t.NumIn()
	output := make([]reflect.Value, inputs, inputs)
	for i := 0; i < inputs; i++ {
		if output[i], ok = vtMap[t.In(i)]; !ok {
			output[i] = reflect.Zero(t.In(i))
		}
	}

	return output
}
