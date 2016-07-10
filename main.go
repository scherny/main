package main

import "fmt"
import "pkg1"
import "pkg3"

func main() {
	fmt.Println("Main")

	pkg1.Pkg1()
	fmt.Println(pkg1.Pkg2())

	pkg3.Pkg3()

}
