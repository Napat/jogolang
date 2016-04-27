/*
https://play.golang.org/p/Mt2p703Plx
(1). Using of empty interface
(2). Assertion: if / switch
(3). Iterate over any type value of slice of emplty interface
*/
package main

import "fmt"
import "reflect"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func voice(a interface{}) { // (1)
	//(2) Assertion if
	if d, ok := a.(dog); ok {
		fmt.Printf("Type: %T\n", d)
	} else if c, ok := a.(cat); ok {
		fmt.Printf("Type: %T\n", c)
	} else {
		fmt.Printf("Type: unknown\n")
	}

	//(2) Assertion: switch
	switch a.(type) {
	case dog:
		fmt.Println(a.(dog).animal.sound)
	case cat:
		fmt.Println(a.(cat).animal.sound)
	default:
		fmt.Println("...")
	}
}

func iterOverAnyTypeSliceInterface(t interface{}) {
	//(3)
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		//iterate over value of slice t
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}

func main() {
	fufu := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}

	voice(fufu)
	//Type: main.dog
	//woof
	voice(fifi)
	//Type: main.cat
	//meow
	critters := []interface{}{fufu, fifi} //(1)
	fmt.Println(critters)
	//[{{woof} true} {{meow} true}]
	for _, cr := range critters {
		fmt.Println(cr)
		//{{woof} true}
		//{{meow} true}
	}
	println("----------")

	data := []string{"one", "two", "three"}
	moredata := []int{1, 2, 3}
	iterOverAnyTypeSliceInterface(data)
	//one
	//two
	//three
	iterOverAnyTypeSliceInterface(moredata)
	//1
	//2
	//3
	iterOverAnyTypeSliceInterface(critters)
	//{{woof} true}
	//{{meow} true}
}
