package main

import "fmt"

type Name interface {
	SpeakName() string
}

type SayMyName interface {
	SayDogName() string
}

type dog struct {
	Name string
}

func GetAdog(name string) Name {
	return dog{Name: name}
}

func (d dog) SpeakName() string {
	return d.Name
}

func (d dog) SayDogName() string {
	return fmt.Sprintf("Dog name is %s", d.Name)
}

func GetSpecialDog(name string) SayMyName {
	return dog{Name: name}
}
