package internal

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 📘 APPRENTISSAGE DE SORTS
////////////////////////////////////////////////////////////////////////////////

// LearnSpell : permet d’apprendre un nouveau sort si le joueur ne le connaît pas déjà
func (c *Character) LearnSpell(spell string) {
	for _, skill := range c.Skills {
		if skill == spell {
			TypeWriter(fmt.Sprintf("⚠️ Vous connaissez déjà le sort %s !", spell), 20*time.Millisecond)
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	TypeWriter(fmt.Sprintf("📘 Nouveau sort appris : %s", spell), 20*time.Millisecond)
}

////////////////////////////////////////////////////////////////////////////////
// 🧪 UTILISATION D’UNE POTION DE VIE
////////////////////////////////////////////////////////////////////////////////

// TakePot : consomme une potion de vie et soigne le joueur
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

		fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		TypeWriter(fmt.Sprintf("🧪 Potion utilisée ! Vous récupérez %d PV", heal), 20*time.Millisecond)
		fmt.Printf("❤️ PV actuels : %d/%d\n", c.CurrentHP, c.MaxHP)
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	} else {
		TypeWriter("⚠️ Vous n’avez plus de Potion de vie dans votre inventaire.", 20*time.Millisecond)
	}
}
