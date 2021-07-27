package main

type AddressDetails struct{
	Street string
	House string
	Phone PhoneNumber
}

type PhoneNumber struct{
	Number1 string
	Number2 map[string]interface{}
}

type Employee struct {
	Name string
	Age  int
	Job  string
	Address AddressDetails
}




func main()  {

	e := Employee{
		Name: "Kayesb",
		Age: 13,
		Job: "CEO",
		Address: AddressDetails{
			Street: "h4",
			House: "43",
			Phone: PhoneNumber{
				Number1: "123",
				Number2: map[string]interface{}{"tel": "159", "tel2": "78478"},
			},
		},
	}

	matchingString := "Name.Address.Phone.Number2.tel2"
	StructParser(&e, matchingString)

}


