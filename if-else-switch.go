package main

import "fmt"

func main() {
	var age int = 20
	if age > 18 {
		fmt.Println("greater than sign implemented")
	} else if age < 18 {
		fmt.Println("less than condition implemented")
	}else if age <= 18 {
		fmt.Println("lessthan or equal sign implemented")
	}else if age == 19 {
		fmt.Println("equal sign impelemted")
	} else {
		fmt.Println("no sign")
	}

	var sex string="male"
	if age == 20 && sex == "male" {
		fmt.Println("and sign implemented")
	} else if age == 20 || sex == "male" {
		fmt.Println("or sign implemented")
	} else {
		fmt.Println("no sign implemented")
	}

	var swithData int = 50

    // Switch case implementation
    switch swithData {
    case 10:
        fmt.Println("Value is 10")
    case 20:
        fmt.Println("Value is 20")
    case 50:
        fmt.Println("Value is 50")
    default:
        fmt.Println("Value is unknown")
    }	
	// Switch case with fallthrough (it will not break after case match it will go next case)
	switch swithData {
	case 10:
		fmt.Println("Value is 10")
		fallthrough
	case 20:
		fmt.Println("Value is 20")
		fallthrough
	case 50:
		fmt.Println("Value is 50")
		fallthrough
	case 30:
		fmt.Println("Value is 50")
		fallthrough
	default:
		fmt.Println("Value is unknown")	
	}		
}