package main

import "fmt"

type iPizza interface {
	prepare()
	bake()
	cut()
	box()
	getName() string
}
type iPizzaStore interface {
	orderPizza(pizzaType string) iPizza
	createPizza(pizzaType string) (iPizza, error)
}

type pizza struct {
	name      string
	pizzaType string
	dough     string
	sauce     string
	toppings  []string
}

type aPizzaStore struct {
	createPizza func(pizzaType string) (iPizza, error)
}

func (p *pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s \n", p.dough)
	fmt.Printf("Adding %s \n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for i, t := range p.toppings {
		fmt.Printf("\t[%d].%s\n", i+1, t)
	}
	fmt.Printf("\n\n")
}
func (p *pizza) bake() {
	fmt.Println("Baking Pizza for 25 minutes to 120 minutes	")
}
func (p *pizza) cut() {
	fmt.Println("Cutting pizza in the diagonal direction")
}
func (p *pizza) box() {
	fmt.Println("Place pizza in official PizzaStore Box")
}
func (p *pizza) getName() string {
	return p.name
}
func (a *aPizzaStore) orderPizza(pizzaType string) iPizza {
	if pizza, err := a.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
		return pizza
	}
}

type cheesePizza struct {
	*pizza
}

/*
func newCheesePizza() iPizza {
	p := &pizza{
		pizzaType: "Cheese",
	}
	return &cheesePizza{
		pizza: p,
	}
}
*/
type greekPizza struct {
	*pizza
}

/*
func newGreekPizza() iPizza {
	p := &pizza{
		pizzaType: "Greek",
	}
	return &greekPizza{
		pizza: p,
	}
}
/*
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

/*
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
*/

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
type nyPizzaStore struct {
	*aPizzaStore
}

func newNYPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}
	nyPizzaStore := &nyPizzaStore{basePizzaStore}
	nyPizzaStore.aPizzaStore.createPizza = nyPizzaStore.createPizza
	return nyPizzaStore
}
func (n *nyPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newNYStyleCheesePizza(), nil
	case "greek":
		return newNYStyleGreekPizza(), nil
	case "pepperoni":
		return newNYStylePepperoniPizza(), nil
	}
	return nil, fmt.Errorf("Invalid pizza type")
}

/*
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
*/

type pizzaStore struct {
	factory iPizzaFactory
}

func newPizzaStore(factory iPizzaFactory) *pizzaStore {
	return &pizzaStore{
		factory: factory,
	}
}

/*
func main() {
	nyPizzaFactory := &nyPizzaFactory{}

	nyStore := newPizzaStore(nyPizzaFactory)
	nyStore.orderPizza("greek")
}
*/

// CHICAGO
type chicagoPizzaStore struct {
	*aPizzaStore
}

func newChicagoPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}
	chicagoPizzaStore := &chicagoPizzaStore{basePizzaStore}
	chicagoPizzaStore.aPizzaStore.createPizza = chicagoPizzaStore.createPizza
	return chicagoPizzaStore
}
func (c *chicagoPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newChicagoStyleCheesePizza(), nil
	case "greek":
		return newChicagoStyleGreekPizza(), nil
	case "pepperoni":
		return newChicagoStylePepperoniPizza(), nil
	}
	return nil, fmt.Errorf("Invalid pizza type")
}

// PIZZA STYLE
type nyStyleCheesePizza struct {
	*pizza
}
type chicagoStyleCheesePizza struct {
	*pizza
}

func newNYStylePizza() iPizza {
	p := &pizza{
		name:     "New York Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}
	return &nyStyleCheesePizza{
		pizza: p,
	}
}
func newChicagoStyleCheesePizza() iPizza {
	p := &pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}
	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}
func newNYStyleGreekPizza() iPizza {
	p := &pizza{
		name:     "Greek Style ft. New York Exclusive Pizza",
		dough:    "Extra Thick and Juicy",
		sauce:    "King of Sauce",
		toppings: []string{"Sun Abbah", "Jabbar"},
	}
	return &nyStyleCheesePizza{
		pizza: p,
	}
}
func newNYStyleCheesePizza() iPizza {
	p := &pizza{
		name:     "New York Pizza with Exclusive Cheese",
		dough:    "Military Dough!",
		sauce:    "Special Sauce",
		toppings: []string{"Dutch's Marinate", "Dutch Milk"},
	}
	return &nyStyleCheesePizza{
		pizza: p,
	}
}
func newNYStylePepperoniPizza() iPizza {
	p := &pizza{
		name:     "New York Pizza with Special Pepperoni of Taiwan",
		dough:    "Juicy Dough with thin thick flavour",
		sauce:    "Taiwan Special Hawaian Ingredient",
		toppings: []string{"Taiwan Sauce", "Taiwan Flavour"},
	}
	return &nyStyleCheesePizza{
		pizza: p,
	}
}
func newChicagoStylePepperoniPizza() iPizza {
	p := &pizza{
		name:     "Chicago Pizza with Delicious Pepperoni",
		dough:    "Rounded Extra Juicy!",
		sauce:    "Special Chicago Flavour",
		toppings: []string{"Extra Chicago Golden Salmon", "Marinate Salmon"},
	}
	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}
func newChicagoStyleGreekPizza() iPizza {
	p := &pizza{
		name:     "Chicago Pizza ft. Greek Aura",
		dough:    "Special Dough from Greek",
		sauce:    "Chicago Special Sauce",
		toppings: []string{"Chicago Golden Salmon", "Marinate Sauce of Salmon"},
	}
	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}

func (c *chicagoStyleCheesePizza) cut() {
	fmt.Println("Cutting the pizza into the square slices")
}
func (b *nyStyleCheesePizza) bake() {
	fmt.Println("Special baking at 1 hours with extra Spicy Moderna")
}

func main() {
	nyPizzaStore := newNYPizzaStore()
	chicagoPizzaStore := newChicagoPizzaStore()
	pizza := nyPizzaStore.orderPizza("cheese")
	fmt.Printf("Ethan ordered %s pizza\n\n", pizza.getName())
	pizza = chicagoPizzaStore.orderPizza("cheese")
	fmt.Printf("Samuel ordered %s pizza\n\n", pizza.getName())
}
