package main

import (
	"fmt"
)

func main() {
	// map with string key and float value
	stock := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70,
	}

	fmt.Println(len(stock))

	fmt.Println(stock["MSFT"])

	fmt.Println(stock["TSLA"])

	value, ok := stock["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	stock["TSLA"] = 822.12
	fmt.Println(stock)

	delete(stock, "AMZN")
	fmt.Println(stock)

	for key := range stock {
		fmt.Println(key)
	}

	for key, values := range stock {
		fmt.Printf("%s -> %.2f\n", key, values)
	}
}
