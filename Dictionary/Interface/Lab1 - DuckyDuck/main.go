package main

import "fmt"

type flyBehaviour interface {
	fly()
}

type QuackBehaviour interface {
	quack()
}

type Duck struct {
	display           func()
	flyingBehaviour   flyBehaviour
	quackingBehaviour QuackBehaviour
}
type flyWithWings struct{}
type flyNoWay struct{}
type Quack struct{}
type MuteQuack struct{}
type Squeak struct{}
type flyRocketPowered struct{}

//Basic func
func (fnw *flyNoWay) fly() {
	fmt.Println("I can't fly")
}

func (fw *flyWithWings) fly() {
	fmt.Println("I'am flying")
}

func (q *Quack) quack() {
	fmt.Println("Quack!!")
}

func (mq *MuteQuack) quack() {
	fmt.Println("...")
}

func (sq *Squeak) quack() {
	fmt.Println("Squeak!!!")
}

//Duck Function
func (d *Duck) performFly() {
	d.flyingBehaviour.fly()
}

func (d *Duck) performQuack() {
	d.quackingBehaviour.quack()
}

func (d *Duck) performSwim() {
	fmt.Println("All ducks float, even decoys!")
}

//Set dynamically
func (d *Duck) setFlyingBehaviour(fb flyBehaviour) {
	d.flyingBehaviour = fb
}

func (d *Duck) setQuackingBehaviour(qb QuackBehaviour) {
	d.quackingBehaviour = qb
}

type MallardDuck struct {
	*Duck
}

func newMallardDuck() *MallardDuck {
	d := &Duck{
		display: func() {
			fmt.Println("I'm real Mallard Duck")
		},
		flyingBehaviour:   &flyWithWings{},
		quackingBehaviour: &Quack{},
	}
	return &MallardDuck{d}
}

type modelDuck struct {
	*Duck
}

func newModelDuck() *modelDuck {
	d := &Duck{
		display: func() {
			fmt.Println("I'm real Model Duck")
		},
		flyingBehaviour:   &flyNoWay{},
		quackingBehaviour: &MuteQuack{},
	}
	return &modelDuck{d}
}

func (frp *flyRocketPowered) fly() {
	fmt.Println("I'm flying with rocket!")
}

func main() {
	MallardDuck := newMallardDuck()
	MallardDuck.display()
	MallardDuck.performFly()
	MallardDuck.performQuack()
	MallardDuck.performSwim()

	ModelDuck := newModelDuck()
	ModelDuck.display()
	ModelDuck.performFly()
	ModelDuck.performQuack()
	ModelDuck.performSwim()

	ModelDuck.setFlyingBehaviour(&flyRocketPowered{})
	fmt.Print("New flying behaviour of Model duck : ")
	ModelDuck.performFly()
	fmt.Println()

	MallardDuck.setQuackingBehaviour(&Squeak{})
	fmt.Print("New quacking behaviour of Mallard duck : ")
	MallardDuck.performQuack()
	fmt.Println()
}
