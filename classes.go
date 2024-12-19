package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Foo struct {
	name    string
	Surname string
}

func (f Foo) Say() string {
	return f.name
}

type Baz struct {
	Foo
}

func (Foo) Speak() string {
	return "foo"
}

type Bar struct {
}

func (Bar) Speak() string {
	return "bar"
}

func main() {
	var foo Speaker = new(Foo)
	var bar Speaker = &Bar{}
	fmt.Println(foo.Speak())
	fmt.Println(bar.Speak())

	baz := Baz{Foo{name: "one", Surname: "baz"}}
	fmt.Println(baz.name)
	fmt.Println(baz.Surname)

}
