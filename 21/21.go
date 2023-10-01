package main

import (
	"fmt"
)

type flying interface {
	Fly()
}

func (this *Parrot) Fly() {
	fmt.Println("Parrot is flying")
}

func (this *Human) Fly() {
	this.BuildPlane()
	fmt.Println("Man is flying on Plane")
}

func (this *Human) BuildPlane() Plane {
	fmt.Println("Plane builded")
	var new_plane Plane
	return new_plane
}

func SendInFly(who_will_fly flying) {
	who_will_fly.Fly()
}

// plane == adapter for Human to fly
type Plane struct{}

type Human struct{}

type Parrot struct{}

func main() {
	some_parrot := Parrot{}
	some_human := Human{}
	SendInFly(&some_parrot)
	SendInFly(&some_human)
}
