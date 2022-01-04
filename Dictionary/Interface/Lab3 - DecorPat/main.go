package main

import "fmt"

// BEVERAGE
type Beverage interface {
	description() string
	cost() float32
}
type beverage struct {
	description string
	milk        bool
	milkCost    float32
	soy         bool
	soyCost     float32
	mocha       bool
	mochaCost   float32
	whip        bool
	whipCost    float32
}

func (b *beverage) getCost() float32 {
	var condimentCost float32
	if b.milk {
		condimentCost += b.milkCost
	}

	if b.soy {
		condimentCost += b.soyCost
	}

	if b.mocha {
		condimentCost += b.mochaCost
	}

	if b.whip {
		condimentCost += b.whipCost
	}
	return condimentCost
}

// OTHER ADDON
// DARK ROAST
type darkRoast struct {
	beverage
	darkRoastCost float32
}

func (dr *darkRoast) getCost() float32 {
	return dr.darkRoastCost + dr.beverage.getCost()
}
func (dr *darkRoast) description() string {
	return "Dark Roast"
}
func (dr *darkRoast) cost() float32 {
	return 0.99
}

// HOUSEBLEND
type houseBlend struct {
	beverage
	houseBlendCost float32
}

func (hb *houseBlend) getCost() float32 {
	return hb.houseBlendCost + hb.beverage.getCost()
}
func (hb *houseBlend) description() string {
	return "House Blend"
}
func (hb *houseBlend) cost() float32 {
	return 0.89
}

// DECAF
type decaf struct {
	beverage
	decafCost float32
}

func (d *decaf) getCost() float32 {
	return d.decafCost + d.beverage.getCost()
}
func (d *decaf) description() string {
	return "Decaf"
}
func (d *decaf) cost() float32 {
	return 1.05
}

// ESPRESSO
type espresso struct {
	beverage
	espressoCost float32
}

func (e *espresso) getCost() float32 {
	return e.espressoCost + e.beverage.getCost()
}

func (e *espresso) description() string {
	return "Espresso"
}

func (e *espresso) cost() float32 {
	return 1.99
}

// MOCHA
type mocha struct {
	beverage Beverage
}

func (m *mocha) description() string {
	return m.beverage.description() + ", Mocha"
}
func (m *mocha) cost() float32 {
	return m.beverage.cost() + .2
}

// MILK
type milk struct {
	beverage Beverage
}

func (m *milk) description() string {
	return m.beverage.description() + ", Milk"
}
func (m *milk) cost() float32 {
	return m.beverage.cost() + .1
}

// SOY
type soy struct {
	beverage Beverage
}

func (s *soy) description() string {
	return s.beverage.description() + ", Soy"
}
func (s *soy) cost() float32 {
	return s.beverage.cost() + .15
}

// WHIP
type whip struct {
	beverage Beverage
}

func (w *whip) description() string {
	return w.beverage.description() + ", Whip"
}
func (w *whip) cost() float32 {
	return w.beverage.cost() + .1
}

func main() {
	beverage := &espresso{}
	fmt.Printf("%s %.2f\n", beverage.description(), beverage.cost())

	darkRoast := &darkRoast{}

	singleMocha := &mocha{
		beverage: darkRoast,
	}

	doubleMocha := &mocha{
		beverage: singleMocha,
	}
	doubleMochaWhip := &whip{
		beverage: doubleMocha,
	}
	fmt.Printf("%s $%.2f\n", doubleMocha.description(), doubleMochaWhip.cost())

	soyMochaWhipHouseBlend := &whip{
		beverage: &mocha{
			beverage: &soy{
				beverage: &houseBlend{},
			},
		},
	}
	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.description(), soyMochaWhipHouseBlend.cost())
}
