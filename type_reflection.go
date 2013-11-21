package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"reflect"
	"time"
)

type (
	Name string

	Address struct {
		Street1 string
		Street2 *string
		City    string
		Region  string
		Zipcode string
	}

	Hashable interface {
		Hash() string
	}
)

func (addr *Address) Hash() string {
	str := fmt.Sprintf("%+v\n", addr)
	return hashString(&str)
}

func (name *Name) Hash() string {
	return hashString((*string)(name))
}

func hashString(str *string) string {
	hasher := sha512.New()
	hasher.Write([]byte(*str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	//name := Name("Eric Gravert")
	address := &Address{
		"Eric Gravert",
		nil,
		"Buffalo",
		"NY",
		"14086"
	}

	nt := reflect.TypeOf(address)
	nv := reflect.ValueOf(address)

	fmt.Println("Type: ", nt)
	fmt.Println("Value: ", nv)

	hashable := Hashable(address)
	fmt.Println("Address Hash: ", hashable.Hash())

	nt = reflect.TypeOf(hashable)
	nv = reflect.ValueOf(hashable)

	fmt.Println("Type: ", nt)
	fmt.Println("Value: ", nv)

	values := make([]int, 10000000, 10000000)
	for i := 0; i < len(values); i++ {
		values[i] = i + 1
	}

	t0 := time.Now()
	Map(func(i int) int { return i * 2 }, values)
	fmt.Println("Time spent: ", time.Duration(time.Since(t0)))

	t0 = time.Now()
	Map2(func(i int) int { return i * 2 }, values) //.([]int)
	fmt.Println("Time spent: ", time.Duration(time.Since(t0)))
}
