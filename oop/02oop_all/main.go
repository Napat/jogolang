/*
https://play.golang.org/p/tu-lcJ-NkT
Struct / explicitly fields / anonymous(embedded) fields / Method / Interface
https://golang.org/ref/spec#Struct_types
A struct is a sequence of named elements, called _**fields**
(1).An embedded type must be specified as a type name T
(2).or as a pointer to a non-interface type name *T
(3).T itself may not be a pointer type
(4).An embedded can be a pointer to a non-interface
 (4.1).Embedded interface type is ok
 (4.2).But pointer to interface does not allow

(5).Tag
A field declaration may be followed by an optional string literal
The tags are made visible through a reflection interface
and take part in type identity for structs but are otherwise ignored.

Method
(6).In Go, a method is a function that is declared with a receiver.
 (6.1).A receiver is a value or
 (6.2).a pointer of a named or
 (6.3).struct type.
-----------------------------------------------
 Values          Attached methods by Receivers
 -----------------------------------------------
 T               (t T)
 *T              (t T) and (t *T)
 -----------------------------------------------
 Rules for receiver
 -If the receiver is a map, func or chan, don't use a pointer to it.
 -If the method needs to mutate the receiver, the receiver must be a pointer.
 -If the receiver is a struct that contains a sync.Mutex or similar synchronizing field,
 the receiver must be a pointer to avoid copying.
 -When in doubt, use a pointer receiver.
 (6.4).A method call x.m() is valid if the method set of (the type of) x contains m.
	If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m():

(7).Interfaces
Interfaces is a TYPE that specifies a method set in Go provide a way to specify the behavior of an object.
It is a convention in Go to name interfaces with an -er suffix when the interface contains only one method.

(8).Promoted the embed type
 All fields and methods of AnonymousField(inner-type) are promoted to the outer-type struct
 (8.1) Promoted fields act like ordinary fields of a struct except that
 they cannot be used as field names in composite literals of the struct.
 (8.2) Given a struct type S and a type named T, promoted methods are included in the method set of the struct as follows:
If S contains an anonymous field T, the method sets of S and *S both include promoted methods with receiver T.
The method set of *S also includes promoted methods with receiver *T.
If S contains an anonymous field *T, the method sets of S and *S both include promoted methods with receiver T or *T.
 -----------------------------------------------
inner-type in outer-type    Promoted methods by inner-type receiver
 -----------------------------------------------
  T     in S               (t T)             (8.3.1).different(weird part)
  T     in *S              (t T) and (t *T)  (8.3.2)
  *T    in S               (t T) and (t *T)  (8.3.3)
  *T    in *S              (t T) and (t *T)  (8.3.4)
Suggestion: When in doubt use STRUCT pointer *S to avoid (8.3.1)
 -----------------------------------------------

*/
package main

import "fmt"

//Junior employee
type Junior struct {
	ID   int    `json:"ID"`
	Name string `json:"MyName"`
}

//Senior employee have some underlying junior
type Senior struct {
	Junior        // (1)
	Underling int // Number of underlying junior
}

//SeniorOSC : Outsourcing senior
type SeniorOSC struct {
	*Junior   // (2)
	Underling int
}

// JuniorOSCp : For demonstate embedded pointer type error
type JuniorOSCp *Junior

/*// SeniorOSCp
type SeniorOSCp struct {
	JuniorOSCp // error (3)
	Underling  int
}*/

// EmployeeInfo :
type EmployeeInfo struct {
	Hellower // (4.1) embedded interface type is ok
}

/*// EmployeeInfoOSC :
type EmployeeInfoOSC struct {
	*Hellower // (4.2) error pointer to interface does not allow
}*/

// Hello :(6) , (6.3)
func (Junior) Hello() {
	fmt.Println("Hello")
}

// Notify : (6.2) Notify method has pointer receiver
// then only implement to *Junior not include Junior
func (j *Junior) Notify() error {
	fmt.Printf("My name is %s, ID: %d\n", j.Name, j.ID)
	return nil
}

// Hellower : (7)
type Hellower interface {
	Hello()
}

// Notifier :
type Notifier interface {
	Notify() error
}

// SendNotification :
func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	// Actually should use struct pointer(&Senior)
	// to shirk not implement interface problem(weird part)
	senior01 := Senior{
		Junior: Junior{
			ID:   1,
			Name: "Alphabet Numberone",
		},
		Underling: 4,
	}
	fmt.Println(senior01)
	senior01.Junior.Hello()
	senior01.Junior.Notify()
	senior01.Hello()
	//(8.3.1) Weird part, senior01 does not contains Notify() but (&senior01)
	senior01.Notify() // (6.4) x.m() is shorthand for (&x).m(): (&senior01).Notify()
	//SendNotification(senior01)	//Error, senior01 does not implement Notifier interface
	SendNotification(&senior01) //Pointer to senior01 contains Notify() then implement Notifier interface

	fmt.Println("----------------------")

	senior02 := SeniorOSC{
		Junior: &Junior{
			ID:   2,
			Name: "Bumblebee Guardian Scout",
		},
		Underling: 3,
	}
	fmt.Println(senior02)
	senior02.Junior.Hello()
	senior02.Junior.Notify()
	senior02.Hello()
	senior02.Notify()          // (8.3.3) senior02 contains methond Notify()
	SendNotification(senior02) // then senior02 imaplement Notifier interface
	SendNotification(&senior02)

	fmt.Println("----------------------")
}
