package internal

import "fmt"

func (c *Character) SpellBook() {
	for _, skill := range c.Skills {
		if skill == "Boule de feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de feu !")
			return
		}
	}
	c.Skills = append(c.Skills, "Boule de feu")
	fmt.Println("Nouveau sort appris : Boule de feu 🔥")
}

func (c *Character) TakePot() {
	if c.Inventory["Potion de vie"] > 0 {
		// Consommer la potion
		c.Inventory["Potion de vie"]--
		if c.Inventory["Potion de vie"] == 0 {
			delete(c.Inventory, "Potion de vie")
		}

		// Rendre des PV
		heal := 50
		c.CurrentHP += heal
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}

		fmt.Printf("🧪 Potion utilisée ! Vous récupérez %d PV → %d/%d\n", heal, c.CurrentHP, c.MaxHP)
	} else {
		fmt.Println("⚠️ Vous n’avez plus de Potion de vie dans votre inventaire.")
	}
}
