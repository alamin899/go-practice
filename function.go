package main
import "fmt"

func addCustom(a int, b int) int {
	return a + b
}

func multiplyCustom(a int, b int) int {
	return a * b
}

func getNumbers(a int,b int) (int, float32) {
	return (a+b), float32(a*b)*2.3
}


func main() {
	// Function to add two integers
	add := func(a int, b int) int {
		return a + b
	}

	// Function to multiply two integers
	multiply := func(a int, b int) int {
		return a * b
	}

	// Calling the functions and printing the results
	sum := add(5, 10)
	product := multiply(5, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)

	addCustomResult := addCustom(5, 10)
	multiplyCustomResult := multiplyCustom(5, 10)
	fmt.Println("Custom Sum:", addCustomResult)
	fmt.Println("Custom Product:", multiplyCustomResult)

	// Calling the function with multiple return values
	sumGetProd, productFromGetProduct := getNumbers(4, 5)
	fmt.Println("sumGetProd:", sumGetProd)
	fmt.Println("productFromGetProduct:", productFromGetProduct)

	// only one value from function will receive
	_, productFromGetProduct = getNumbers(4, 5)
	fmt.Println("productFromGetProductOnly:", productFromGetProduct)

	// Give input using scanf
	var a, b int
	fmt.Println("Enter two numbers for scanf:")
	fmt.Scanf("%d %d", &a, &b)
	sum = add(a, b)
	fmt.Println("Sum of", a, "and", b, "is:", sum)

	//Give input using scanln
	var c, d int
	fmt.Println("Enter two numbers for scanln:")
	fmt.Scanln(&c, &d)
	sum = add(c, d)
	fmt.Println("Sum of", c, "and", d, "is:", sum)

	// Give input using scan
	var e, f int
	fmt.Println("Enter two numbers for scan:")
	fmt.Scan(&e, &f)
	sum = add(e, f)
	fmt.Println("Sum of", e, "and", f, "is:", sum)
}