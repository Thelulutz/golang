package variables

import "fmt"

// Print self-assigned values of variables
func PrintVariables() {
	var varInt int
	var varFloat float32
	var varStr string
	var varBool bool

	fmt.Println(varInt)
	fmt.Println(varFloat)
	fmt.Println(varStr)
	fmt.Println(varBool)
}
