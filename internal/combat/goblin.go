package internal

func InitGoblin(name string, maxHP, currentHP int) Monster {
	return Monster{
		Name:      "Gobelin d’entrainement",
		MaxHP:     40,
		CurrentHP: 40,
		AtkPoints: 5,
	}
}

func goblinPattern() {

}
