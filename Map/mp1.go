package main

import "fmt"

func main() {
	zipcodes := map[string]int{
		"SF": 94103,
		"NY": 10001,
		"CH": 10000,
	}

	zipcodes["MN"] = 50008
	fmt.Println("zipcodes", zipcodes)
	fmt.Println("len(zipcodes)=", len(zipcodes))

	// iteration

	for k, v := range zipcodes {
		fmt.Printf("zipcodes[%v]= %d\n", k, v)
	}

	//zipcode check
	zipcodeTocheck := "CA"
	if val, exists := zipcodes[zipcodeTocheck]; exists {
		fmt.Printf("Zipcode of %v is %v\n", zipcodeTocheck, val)
	} else {
		fmt.Printf("Zipcode of %v not found\n", zipcodeTocheck)
	}

	//DELETE A key from
	zipcodeToremove := "NY"
	delete(zipcodes, zipcodeToremove)
	fmt.Println("After removing Zipcode ", zipcodes)

}
