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
	Additional     Skill
	Additional1    Weapon
}
type Skill struct {
	Reflect             bool
	AbsoluteDefense     bool
	DragonHeart         bool
	IronBody            bool
	AssassinLineage     bool
	RoyalLineage        bool
	WarriorLineage      bool
	MahaChampion        bool
	BlessingGodOfWar    bool
	BlessingGodOfWisdom bool
	DivineExcalibur     bool
}
type Weapon struct {
	Excalibur bool
}

func (calc Base) MyStatus(
	Name string,
	Level float64,
	BaseHP float64,
	BaseATK float64,
	BaseDEF float64,
	BaseCritical float64,
	Reflect bool,
	AD bool,
	DH bool,
	IB bool,
	AL bool,
	RL bool,
	WL bool,
	MC bool,
	BGO bool,
	BGW bool,
	Exca bool) {
	if Exca {
		BaseATK = BaseATK * 3
		BaseHP = BaseHP * 1.3
		BaseCritical = BaseCritical * 0.5
		BaseDEF = BaseDEF * 3
		calc.Additional.DivineExcalibur = true
	}
	if BGO {
		Level = Level * 2
	}
	if BGW {
		BaseHP *= 2
		BaseATK *= 2
		BaseDEF *= 2
		BaseCritical *= 1.2
	}
	HP := BaseHP + (float64(Level) * BaseHP * 0.6)
	ATK := BaseATK + (float64(Level) * BaseATK * 0.4)
	DEF := BaseDEF + (Level * BaseDEF * 0.3)
	Critical := BaseCritical + (Level * BaseCritical * 0.2)
	if Reflect {
		AddRef := DEF*0.7 + Level*10
		HP = HP + AddRef
		ATK = ATK + AddRef
	}
	if AD {
		AddAD := HP*0.01 + DEF*0.2
		DEF = DEF + AddAD
		HP = HP + AddAD*0.5
	}
	if DH {
		AddDH := HP*2 + DEF*5
		HP = HP + AddDH*1.5 + DEF
		DEF = (DEF + AddDH*0.05) - DEF*0.5
	}
	if IB {
		AddIB := DEF*3 + HP*0.001
		Critical = Critical - DEF*0.002
		DEF = DEF + AddIB
	}
	if AL {
		AddAL := Critical * 2.5
		DEF = DEF - Critical*50
		Critical = Critical + AddAL
	}
	if RL {
		HP = HP + HP*2.3
		ATK = ATK + ATK*1.8
		DEF = DEF + DEF*1.7
		Critical = Critical * 0.8
	}
	if WL {
		AddWL := HP*0.7 - Critical*10
		HP = HP + AddWL*1.1
		ATK = ATK + AddWL*1.3
		DEF = DEF + AddWL*0.4
	}
	if MC {
		MahaSpecial := HP + ATK + DEF + Critical*10
		HP = HP + MahaSpecial*2
		ATK = ATK + MahaSpecial*1.2
		DEF = DEF + MahaSpecial*0.8
		Critical = Critical + MahaSpecial*0.0005
	}
	if calc.Additional.DivineExcalibur {
		ATK = ATK * 1.5
		DEF = ATK * 1.5
	}
	fmt.Printf("\nHere is status of %s\nHP:\t\t\t %.f\nATK:\t\t\t %.f\nDEF:\t\t\t %.f\nCritical:\t\t %.2f\n\n", Name, HP, ATK, DEF, Critical)
}

//STATUS

//Main
func main() {
	values := Base{
		Name:           "Reinhard Novus",
		Level:          200,
		BaseHP:         90,
		BaseATK:        50,
		BaseDEF:        34,
		BaseCritChance: 3,
		Additional: Skill{
			Reflect:             true,
			AbsoluteDefense:     true,
			DragonHeart:         true,
			IronBody:            true,
			AssassinLineage:     false,
			RoyalLineage:        true,
			WarriorLineage:      true,
			MahaChampion:        true,
			BlessingGodOfWar:    true,
			BlessingGodOfWisdom: true,
		},
		Additional1: Weapon{},
	}
	Bojack := Base{
		Name:           "Bojack The Terminator",
		Level:          310,
		BaseHP:         400,
		BaseATK:        20,
		BaseDEF:        120,
		BaseCritChance: 0.5,
		Additional: Skill{
			Reflect:          true,
			AbsoluteDefense:  true,
			DragonHeart:      true,
			IronBody:         true,
			AssassinLineage:  true,
			RoyalLineage:     false,
			WarriorLineage:   true,
			MahaChampion:     true,
			BlessingGodOfWar: true,
		},
	}
	Gerald := Base{
		Name:           "Gerald The Champion",
		Level:          152,
		BaseHP:         80,
		BaseATK:        80,
		BaseDEF:        30,
		BaseCritChance: 7,
		Additional: Skill{
			Reflect:          true,
			AbsoluteDefense:  true,
			DragonHeart:      true,
			IronBody:         true,
			AssassinLineage:  false,
			RoyalLineage:     true,
			WarriorLineage:   true,
			MahaChampion:     true,
			BlessingGodOfWar: true,
		},
		Additional1: Weapon{
			Excalibur: true,
		},
	}
	values.MyStatus(
		values.Name,
		float64(values.Level),
		values.BaseHP,
		values.BaseATK,
		values.BaseDEF,
		values.BaseCritChance,
		values.Additional.Reflect,
		values.Additional.AbsoluteDefense,
		values.Additional.DragonHeart,
		values.Additional.IronBody,
		values.Additional.AssassinLineage,
		values.Additional.RoyalLineage,
		values.Additional.WarriorLineage,
		values.Additional.MahaChampion,
		values.Additional.BlessingGodOfWar,
		values.Additional.BlessingGodOfWisdom,
		values.Additional1.Excalibur)

	Bojack.MyStatus(
		Bojack.Name,
		float64(Bojack.Level),
		Bojack.BaseHP,
		Bojack.BaseATK,
		Bojack.BaseDEF,
		Bojack.BaseCritChance,
		Bojack.Additional.Reflect,
		Bojack.Additional.AbsoluteDefense,
		Bojack.Additional.DragonHeart,
		Bojack.Additional.IronBody,
		Bojack.Additional.AssassinLineage,
		Bojack.Additional.RoyalLineage,
		Bojack.Additional.WarriorLineage,
		Bojack.Additional.MahaChampion,
		Bojack.Additional.BlessingGodOfWar,
		Bojack.Additional.BlessingGodOfWisdom,
		Bojack.Additional1.Excalibur)
	Gerald.MyStatus(
		Gerald.Name,
		float64(Gerald.Level),
		Gerald.BaseHP,
		Gerald.BaseATK,
		Gerald.BaseDEF,
		Gerald.BaseCritChance,
		Gerald.Additional.Reflect,
		Gerald.Additional.AbsoluteDefense,
		Gerald.Additional.DragonHeart,
		Gerald.Additional.IronBody,
		Gerald.Additional.AssassinLineage,
		Gerald.Additional.RoyalLineage,
		Gerald.Additional.WarriorLineage,
		Gerald.Additional.MahaChampion,
		Gerald.Additional.BlessingGodOfWar,
		Gerald.Additional.BlessingGodOfWisdom,
		Gerald.Additional1.Excalibur)
}
