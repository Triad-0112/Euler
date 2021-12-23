package main

import (
	"fmt"
)

//type damager interface {
//EnemyDamage()
//OwnDamage()
//}
type OwnStatus interface {
	MyStatus()
}
type Base struct {
	Name           string
	Level          int
	BaseHP         float64
	BaseATK        float64
	BaseDEF        float64
	BaseCritChance float64
	Perk           Skill
}
type Skill struct {
	Reflect         bool
	AbsoluteDefense bool
	DragonHeart     bool
	IronBody        bool
	AssassinLineage bool
	RoyalLineage    bool
	WarriorLineage  bool
	MahaChampion    bool
}

func (calc Base) MyStatus(
	Name string,
	Level float64,
	BaseHP float64,
	BaseATK float64,
	BaseDEF float64,
	BaseCritical float64,
	Reflect bool,
	AbsoluteDefense bool,
	DragonHeart bool,
	IronBody bool,
	AssassinLineage bool,
	RoyalLineage bool,
	WarriorLineage bool,
	MahaChampion bool,
) {
	HP := BaseHP + (float64(Level) * BaseHP * 0.6)
	ATK := BaseATK + (float64(Level) * BaseATK * 0.4)
	DEF := BaseDEF + (Level * BaseDEF * 0.3)
	Critical := BaseCritical + (Level * BaseCritical * 0.2)
	fmt.Printf("\nHere is status of %s\nHP:\t\t\t %.f\nATK:\t\t\t %.f\nDEF:\t\t\t %.f\nCritical:\t\t %.2f\n\n", Name, HP, ATK, DEF, Critical)
}

//Main
func main() {
	values := Base{
		Name:           "Reinhard Novus",
		Level:          100,
		BaseHP:         89,
		BaseATK:        15,
		BaseDEF:        31,
		BaseCritChance: 2,
		Perk: Skill{
			Reflect:         true,
			AbsoluteDefense: true,
			DragonHeart:     true,
			IronBody:        true,
			AssassinLineage: true,
			RoyalLineage:    true,
			WarriorLineage:  true,
			MahaChampion:    true,
		},
	}
	values.MyStatus(
		values.Name,
		float64(values.Level),
		values.BaseHP,
		values.BaseATK,
		values.BaseDEF,
		values.BaseCritChance,
		values.Perk.Reflect,
		values.Perk.AbsoluteDefense,
		values.Perk.DragonHeart,
		values.Perk.IronBody,
		values.Perk.AssassinLineage,
		values.Perk.RoyalLineage,
		values.Perk.WarriorLineage,
		values.Perk.MahaChampion,
	)
}
