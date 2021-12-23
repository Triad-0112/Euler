package main

import "fmt"

type Creature struct {
	Name  string
	Race  string
	BP    uint64
	SSJ   uint64
	SSJ2  uint64
	SSJ3  uint64
	SSJG  uint64
	SSJGS uint64
}

var Character = []string{"Son Goku"}
var Status = map[string]uint64{
	"HP":               0,
	"KI":               0,
	"Speed":            0,
	"Strength":         0,
	"Potential Threat": 0,
	"Aura":             0,
	"God Status":       0,
}
var Race = map[string]uint64{
	"Human":    75,
	"Saiyan":   120,
	"Namekian": 90,
	"Buu":      140,
	"Android":  125,
	"Cell":     150,
}
var BPower uint64
var character map[string]Creature
var ComposedPower string = "???"
var Corr string = "?"

func toSSJ(a uint64, b string) uint64 {
	switch b {
	case "1":
		a *= 20
	case "2":
		a *= 50
	case "3":
		a *= 100
	case "G":
		a = toSSJ(a, "1") * 10
	case "GS":
		a = toSSJ(a, "G") * 20
	case "UI":
		a = toSSJ(a, "GS") * 20
		a += 100000000
	default:
		fmt.Println("Nice Try!")
	}
	return a
}
func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}
func Supression(a uint64, b uint64) uint64 {
	a = toSSJ(a, "UI") / 1000000
	a /= b * 10000
	return a
}

func Search(x string, BP uint64) {
	character = make(map[string]Creature)
	character[x] = Creature{x, "Saiyan", BPower, 0, 0, 0, 0, 0}
	v := character[x]
	p := &v
	BP = p.BP
	fmt.Printf("Name\t: %s\nRace\t: %s\nBP\t: %d\n=============================================\nSuper Saiyan Power\t: %d\nSuper Saiyan 2 Power\t: %d\nSuper Saiyan 3 Power\t: %d\n=============================================\n", p.Name, p.Race, p.BP, toSSJ(p.BP, "1"), toSSJ(p.BP, "2"), toSSJ(p.BP, "3"))
	fmt.Println(center("GOD", 45), "\n=============================================")
	fmt.Printf("Super Saiyan God\t: %d\nSuper Saiyan Blue\t: %d\nUltra Instinct\t\t: %d\n=============================================\n", toSSJ(p.BP, "G"), toSSJ(p.BP, "GS"), toSSJ(p.BP, "UI"))
	fmt.Println(center("SUPRESS", 45), "\n=============================================")
	fmt.Printf("Hidden UI\t\t: %d%s", Supression(toSSJ(p.BP, "UI"), 5), ComposedPower)
}
func BPCalc(a, b, c, d, e, f, g, h, BPower uint64) {
	var godPower uint64
	a = 210 * a * h
	b = 300 * b * h
	c = 170 * c * h
	d = 400 * d * h
	switch e {
	case 1:
		e = 1 * 400 * e * h
	case 2:
		e = 3 * 500 * e * h
	case 3:
		e = 4 * 700 * e * h
	case 4:
		e = 7 * 1000 * e * h
	case 5:
		e = 10 * 1200 * e * h
	default:
		e = 0
	}
	f = 350 * f * h
	if g == 1 {
		godPower = h * 7000
	}
	BPower = a + b + c + d + e + f + godPower
}
func main() {
	var Name string
	fmt.Println("Use data in your glasses!")
	var pp string
	racial := Race[pp]
	fmt.Print("Race\t: \n")
	fmt.Print("Battle Power\t: \n")
	fmt.Print("HP\t: \n")
	fmt.Print("KI\t: \n")
	fmt.Print("Speed\t: \n")
	fmt.Print("Strength\t: \n")
	fmt.Print("Potential Threat\t: \n")
	fmt.Print("Aura\t: \n")
	fmt.Print("God Status\t: \n")
	fmt.Scanf("%s %d %d %d %d %d %d %d %d", racial, BPower, Status["HP"], Status["KI"], Status["Speed"], Status["Strength"], Status["Potential Threat"], Status["Aura"], Status["God Status"])
	BPCalc(Status["HP"], Status["KI"], Status["Speed"], Status["Strength"], Status["Potential Threat"], Status["Aura"], Status["God Status"], racial, BPower)
	Search(Name, BPower)
}
