package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

type Renamable interface {
	Rename(newName string)
}

func (salutation *Salutation) Rename(newName string) {
	salutation.name = newName
}

func (salutation *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salutation.Rename(s)
	n = len(s)
	err = nil
	return
}

type Salutations []Salutation

type Printer func(string)

func (salutations Salutations) Greet(do Printer, isFormal bool, times int) {

	for _, s := range salutations {
		message, alternate := CreateMessage(s.name, s.greeting)

		if prefix := GetPrefix(s.name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutations {
		c <- s
	}
	close(c)
}

func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "
	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "Hey!! " + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
func main() {
	// var message = Salutation{greeting: "Hii", name: "Bob"}

	// s := []int{1, 10, 500, 25}

	salutations := Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "Whats up?"},
	}

	// salutations[0].Rename("John")

	// RenameToFrog(&salutations[0])

	fmt.Fprintf(&salutations[0], "Count is %d", 10)

	c := make(chan Salutation)
	c2 := make(chan Salutation)
	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}

	}

	// for s := range c {
	// 	fmt.Println(s.name)
	// }

	// Greet(salutations, CreatePrintFunction("!!!!"), true, 5)

	// done := make(chan bool)

	// go func() {
	// 	salutations.Greet(CreatePrintFunction("AAAAAAA"), true, 5)
	// 	done <- true
	// 	time.Sleep(100 * time.Millisecond)
	// 	done <- true
	// 	fmt.Println("Donee")
	// }()

	// salutations.Greet(CreatePrintFunction("!!!!"), true, 5)
	// <-done
	// for {
	// 	time.Sleep(100 * time.Millisecond)
	// }
}

func RenameToFrog(r Renamable) {
	r.Rename("Frog")
}
