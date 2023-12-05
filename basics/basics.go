package main

import "fmt"

var a int // Single Variable

var (
	b bool
	c float32
	d string
	e uint32
)

func main() {
	// Manual value assingments
	a = 69
	b, c = true, 32.9
	d = "Hi Mom"
	fmt.Println(a, b, c, d, e)

	// Initialize and assigns
	f := 32 // Will automtically be an integer
	fmt.Println(f)

	/* User specified types */
	const g int32 = 12        // 32-bit integer
	const h float32 = 20.5    // 32-bit float
	var i complex128 = 1 + 4i // 128-bit complex number
	var j uint16 = 14         // 16-bit unsigned integer

	/* Default types */
	k := 42              // int
	pi := 3.14           // float64
	x, y := true, false  // bool
	z := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", g, h, i, j)
	fmt.Printf("default types:\n %T %T %T %T %T\n", k, pi, x, y, z)

	/* Define an array of size 4 that stores string values */
	var DeploymentPlatforms = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the array */
	for i := 0; i < len(DeploymentPlatforms); i++ {
		option := DeploymentPlatforms[i]
		fmt.Println(i, option)
	}

	Lessons := [...]string{"Math", "English"}

	for index, option := range Lessons {
		fmt.Println(index, option)
	}
}
