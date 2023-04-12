package typeconversion

import "fmt"

func ConvertSpeed() {

	// creating variables
	speed := 100 // speed is int
	force := 2.5 // force is float64

	speed_int := speed * int(force)
	speed = int(float64(speed) * force)

	fmt.Println(speed_int)
	fmt.Println(speed)
}
