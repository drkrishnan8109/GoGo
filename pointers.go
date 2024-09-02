package main

import "fmt"

func main() {
	p := 3
	intpointer(&p)
	fmt.Println(p) // Output: 6

	p1 := Person{"drk", 5}
	makePersonOld(&p1)
	fmt.Println(p1.age)

	var p2 Person
	p2 = Person{"fakenamde", 35}
	makePersonOld(&p2)
	fmt.Println(p2.age)

}

func intpointer(p *int) {
	*p = *p + *p //Here *p is required bcs p is a pointer and we need to pick value of that pointer
}

type Person struct {
	name string
	age  uint32
}

func makePersonOld(p *Person) {
	p.age = p.age + 100 //Here * is not needed bcs p is a pointer and p.age is int
}
