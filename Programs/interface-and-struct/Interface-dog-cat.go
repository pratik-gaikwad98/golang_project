package interfaceandstruct

import "fmt"

type Speak interface {
	SpeakLoudly()
}

func AnimalSpeak(s Speak) {
	s.SpeakLoudly() //this function will call method cat or dog at runtime
}

type Dog struct{}

func (d Dog) SpeakLoudly() {
	fmt.Println("BHOW BHOW")
}

type Cat struct{}

func (c Cat) SpeakLoudly() {
	fmt.Println("MEOW MEOW")
}

func CallSpeakAnimal() {
	d := Dog{}
	AnimalSpeak(d)

	c := Cat{}
	AnimalSpeak(c)
}
