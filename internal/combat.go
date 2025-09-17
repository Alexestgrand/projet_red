package internal

import (
	"fmt"
	"time"
)

// 🎮 characterTurn : gère le tour du joueur
func (c *Character) characterTurn(monster *Monster) {
	TypeWriter("\n🌟 --- Tour du joueur --- 🌟", 20*time.Millisecond)
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Print("👉 Choix : ")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	// ==== ATTAQUE / SORTS ====
	case 1:
		fmt.Printf("\n💥 Points de Mana : %d/%d\n", c.Mana, c.MaxMana)
		TypeWriter("=== 📖 Attaques & Sorts disponibles ===", 15*time.Millisecond)

		// Liste des sorts connus par le joueur
		for i, spell := range c.Skills {
			switch spell {
			case "Coup de poing":
				fmt.Printf("%d. %s (Dégâts: 8 | Mana: 0)\n", i+1, spell)
			case "Boule de feu":
				fmt.Printf("%d. %s (Dégâts: 18 | Mana: 10)\n", i+1, spell)
			case "Soin léger":
				fmt.Printf("%d. %s (+15 PV | Mana: 5)\n", i+1, spell)
			default:
				fmt.Printf("%d. %s\n", i+1, spell)
			}
		}

		// Choix du sort
		var spellChoice int
		fmt.Print("👉 Choix du sort : ")
		fmt.Scanln(&spellChoice)

		if spellChoice > 0 && spellChoice <= len(c.Skills) {
			c.CastSpell(c.Skills[spellChoice-1], monster)
		} else {
			TypeWriter("⚠️ Choix invalide, votre tour est perdu !", 20*time.Millisecond)
		}

	// ==== INVENTAIRE ====
	case 2:
		c.UseItem(monster)

	default:
		TypeWriter("⚠️ Choix invalide !", 20*time.Millisecond)
	}
}

// ==== Lancement des sorts ====
func (c *Character) CastSpell(spell string, target *Monster) {
	switch spell {
	case "Coup de poing":
		damage := 8
		TypeWriter(fmt.Sprintf("👊 %s utilise %s et inflige %d dégâts à %s !",
			c.Name, spell, damage, target.Name), 20*time.Millisecond)
		target.CurrentHP -= damage

	case "Boule de feu":
		manaCost := 10
		damage := 18
		if c.Mana < manaCost {
			TypeWriter("⚠️ Pas assez de mana pour lancer Boule de feu !", 20*time.Millisecond)
			return
		}
		c.Mana -= manaCost
		TypeWriter(fmt.Sprintf("🔥 %s lance %s et inflige %d dégâts à %s !",
			c.Name, spell, damage, target.Name), 20*time.Millisecond)
		target.CurrentHP -= damage

	case "Soin léger":
		manaCost := 5
		heal := 15
		if c.Mana < manaCost {
			TypeWriter("⚠️ Pas assez de mana pour utiliser Soin léger !", 20*time.Millisecond)
			return
		}
		c.Mana -= manaCost
		c.CurrentHP += heal
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		TypeWriter(fmt.Sprintf("✨ %s utilise %s et regagne %d PV ! ❤️ (%d/%d PV)",
			c.Name, spell, heal, c.CurrentHP, c.MaxHP), 20*time.Millisecond)

	default:
		TypeWriter("⚠️ Sort inconnu.", 20*time.Millisecond)
	}

	// Vérifie les PV de la cible
	if target.CurrentHP < 0 {
		target.CurrentHP = 0
	}

	// Affiche l’état du monstre après l’attaque
	TypeWriter(fmt.Sprintf("🛡️ %s → PV : %d/%d", target.Name, target.CurrentHP, target.MaxHP), 15*time.Millisecond)
}

// 🎮 Combat d'entraînement contre un gobelin
func trainingFight(c *Character) {
	goblin := InitGoblin()

	TypeWriter("\n⚔️ Début du combat d’entraînement !", 20*time.Millisecond)
	TypeWriter(fmt.Sprintf("Un %s apparaît avec %d PV.", goblin.Name, goblin.MaxHP), 20*time.Millisecond)

	tour := 1

	// Déterminer qui commence
	playerFirst := c.Initiative >= goblin.Initiative
	if playerFirst {
		TypeWriter(fmt.Sprintf("🎲 %s a plus d’initiative et commence le combat !", c.Name), 20*time.Millisecond)
	} else {
		TypeWriter(fmt.Sprintf("🎲 %s a plus d’initiative et attaque en premier !", goblin.Name), 20*time.Millisecond)
	}

	// Boucle de combat tour par tour
	for c.CurrentHP > 0 && goblin.CurrentHP > 0 {
		TypeWriter(fmt.Sprintf("\n===== TOUR %d =====", tour), 15*time.Millisecond)

		if playerFirst {
			// Tour du joueur
			c.characterTurn(&goblin)
			if goblin.CurrentHP <= 0 {
				TypeWriter(fmt.Sprintf("\n✅ %s a vaincu le %s !", c.Name, goblin.Name), 20*time.Millisecond)
				c.GainExp(30) // XP gagnée
				break
			}

			// Tour du gobelin
			goblinPattern(c, &goblin, tour)
			if c.CurrentHP <= 0 {
				isDead(c)
				TypeWriter("🏁 Combat terminé, retour au menu principal.", 20*time.Millisecond)
				return
			}
		} else {
			// Tour du gobelin
			goblinPattern(c, &goblin, tour)
			if c.CurrentHP <= 0 {
				isDead(c)
				TypeWriter("🏁 Combat terminé, retour au menu principal.", 20*time.Millisecond)
				return
			}

			// Tour du joueur
			c.characterTurn(&goblin)
			if goblin.CurrentHP <= 0 {
				TypeWriter(fmt.Sprintf("\n✅ %s a vaincu le %s !", c.Name, goblin.Name), 20*time.Millisecond)
				c.GainExp(30)
				break
			}
		}

		tour++
	}

	TypeWriter("\n🏁 Fin du combat d’entraînement, retour au menu principal.", 20*time.Millisecond)
}
