package main

import (
	"fmt"
	"log"
	"reflect"
)

type Address struct {
	Street1 string  `json:"street_1"`
	Street2 *string `json:"street_2"`
	City    string  `json:"city"`
	Region  string  `json:"state"`
	Zipcode string  `json:"zip_code"`
	//status  bool    `json:"fraud_address"`
}

func main() {
	eric := &Address{
		"1 Main St",
		nil,
		"Buffalo",
		"NY",
		"14086",
		// true,
	}
	displayStructTags(eric)
}

func displayStructTags(s interface{}) {
	// get the value and type of s
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// lets make sure we are actually working with a struct
	if t.Kind() != reflect.Struct {
		log.Fatalf("Wow, you can't even get the right type in a demo.... fail!\nKind()  was: %s\n ", t.Kind())
	}
	log.Println(t.Kind())

	// get the number of fields and loop through them
	fieldCount := t.NumField()
	for i := 0; i < fieldCount; i++ {
		// tag is accessed through the type, the value through value.. derp
		fmt.Printf("{\"%s\":\"%s\"}\n", t.Field(i).Tag.Get("json"), v.Field(i))
	}
}

/*
  // check for pointer and get underlying type
  if t.Kind() == reflect.Ptr {
    t = t.Elem()
  }
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

  // check to see if a feed is exported and settable
		if v.Field(i).CanSet() {
*/
