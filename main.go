package main

import (
	"fmt"
	"reflect"
)

type AddressDetails struct {
	Street string
	House  string
	Phone  PhoneNumber
}

type PhoneNumber struct {
	Number1 string
	Number2 map[string]interface{}
}

type Employee struct {
	Name    string
	Age     int
	Job     string
	Address AddressDetails
}

func StructParser(obj interface{}, matchedString string) interface{}{
	e := reflect.ValueOf(obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName,varType,varValue)
	}
	return nil
}

func main() {

	e := Employee{
		Name: "Kayes Shawon",
		Age:  13,
		Job:  "CEO",
		Address: AddressDetails{
			Street: "h4",
			House:  "43",
			Phone: PhoneNumber{
				Number1: "123",
				Number2: map[string]interface{}{"tel": "159", "tel2": "78478", "tel3": 2322423},
			},
		},
	}

	//matchingString := "Address.Phone.Number2.tel2"
	matchingString := "Name"
	StructParser(&e, matchingString)

}
