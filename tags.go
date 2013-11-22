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
	//status  bool    `json:"fraud_address"` --uncomment to test unexported fields
}

func main() {
	eric := &Address{
		"1 Main St",
		nil,
		"Buffalo",
		"NY",
		"14086",
		// true,--uncomment to test unexported fields
	}
	displayStructTags(eric)
}

func displayStructTags(s interface{}) {
	// get the value and type of s
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// #1

	// lets make sure we are actually working with a struct
	if t.Kind() != reflect.Struct {
		log.Fatalf("Wow, you can't even get the right type in a demo.... fail!\nKind()  was: %s\n ", t.Kind())
	}
	log.Println(t.Kind())

	// get the number of fields and loop through them
	fieldCount := t.NumField()
	for i := 0; i < fieldCount; i++ {
		// tag is accessed through the type, the value through value.. derp
		// #2 see below
		fmt.Printf("{\"%s\":\"%s\"}\n", t.Field(i).Tag.Get("json"), v.Field(i))
	}
}

/*
  #1 to get the type of the right kin,d use elem which returns the value pointed to by a ptr
  in the real world, it would be necessary to check the underlying type to check for a ptr or nil
  // check for pointer and get underlying type
  if t.Kind() == reflect.Ptr {
    t = t.Elem()
  }
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}


  #2 Use CanSet() to see if a field is exported. You can also use the PkgPath to determine if a field is exported
	  if v.Field(i).CanSet() {
		fmt.Printf("{\"%s\":\"%s\"}\n", t.Field(i).Tag.Get("json"), v.Field(i))
	  }
*/
