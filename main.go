package main

import "fmt"
import "pkg1"

func main() {
	fmt.Println("Main")
	pkg1.Pkg1()
	fmt.Println(pkg1.Pkg2())
}
