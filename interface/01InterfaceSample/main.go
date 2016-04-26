/*
https://play.golang.org/p/RDlZ3oZ4L-

// Interface example by using internal sort package
// Implement default print format with String() method

https://golang.org/pkg/sort/#Sort
https://golang.org/pkg/sort/#Interface
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

https://golang.org/doc/effective_go.html#printing
If you want to control the default format for a custom type, all that's required
is to define a method with the signature String() string on the type.
For our simple type T
however: don't construct a String method by calling Sprintf in a way
that will recur into your String method indefinitely
*/
package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d Zoo!", p.Name, p.Age)
	//return fmt.Sprintf("try recure %s", p) // Error: will recur forever.
}

// ByAge implements sort.Interface for []person based on the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []person based on the Age field.
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []person{
		{"Judy", 20},
		{"Big", 41},
		{"Nick", 25},
		{"Flash", 26},
	}

	fmt.Println(people[0])
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
