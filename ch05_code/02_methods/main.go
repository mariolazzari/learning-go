package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 9, "woof"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Speak()
	poodle.SpeakThreeTimes()

}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// method via receaver
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
