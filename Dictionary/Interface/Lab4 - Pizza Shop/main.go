package main

import "fmt"

type iPizza interface {
	prepare()
	bake()
	cut()
	box()
}

type pizza struct {
	pizzaType string
}

func (p *pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.pizzaType)
}
func (p *pizza) bake() {
	fmt.Printf("Baking %s Pizza\n", p.pizzaType)
}
func (p *pizza) cut() {
	fmt.Printf("Cutting %s Pizza\n", p.pizzaType)
}
func (p *pizza) box() {
	fmt.Printf("Boxing %s Pizza\n", p.pizzaType)
}

type cheesePizza struct {
	*pizza
}

func newCheesePizza() iPizza {
	p := &pizza{
		pizzaType: "Cheese",
	}
	return &cheesePizza{
		pizza: p,
	}
}

type greekPizza struct {
	*pizza
}

func newGreekPizza() iPizza {
	p := &pizza{
		pizzaType: "Greek",
	}
	return &greekPizza{
		pizza: p,
	}
}

type pepperoniPizza struct {
	*pizza
}

func newPepperoniPizza() iPizza {
	p := &pizza{
		pizzaType: "Pepperoni",
	}
	return &pepperoniPizza{
		pizza: p,
	}
}

/*
func (ps *pizzaStore) orderPizza(pizzaType string) {
	var pizza iPizza
	switch pizzaType {
	case "cheese":
		pizza = newCheesePizza()
	case "greek":
		pizza = newGreekPizza()
	case "pepperoni":
		pizza = newPepperoniPizza()
	}
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
}
*/

type simplePizzaFactory struct{}

func (spf *simplePizzaFactory) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newCheesePizza(), nil
	case "greek":
		return newGreekPizza(), nil
	case "pepperoni":
		return newPepperoniPizza(), nil
	}
	return nil, fmt.Errorf("Invalid pizza type")
}

/*
type pizzaStore struct {
	factory *simplePizzaFactory
}
*/
/*
func newPizzaStore() *pizzaStore {
	return &pizzaStore{
		factory: &simplePizzaFactory{},
	}
}
*/
func (ps *pizzaStore) orderPizza(pizzaType string) {
	if pizza, err := ps.factory.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
	}
}

type iPizzaFactory interface {
	createPizza(pizzaType string) (iPizza, error)
}
type nyPizzaFactory struct{}

func (nypf *nyPizzaFactory) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newCheesePizza(), nil
	case "greek":
		return newGreekPizza(), nil
	case "pepperoni":
		return newPepperoniPizza(), nil
	}
	return nil, fmt.Errorf("Invalid pizza type")
}

type pizzaStore struct {
	factory iPizzaFactory
}

func newPizzaStore(factory iPizzaFactory) *pizzaStore {
	return &pizzaStore{
		factory: factory,
	}
}

func main() {
	nyPizzaFactory := &nyPizzaFactory{}

	nyStore := newPizzaStore(nyPizzaFactory)
	nyStore.orderPizza("greek")
}
