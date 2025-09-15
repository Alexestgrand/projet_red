package internal

import (
	"fmt"
	"time"
)

// ===== MARCHAND =====
func (c *Character) MarchantInterface() {
	var choix int
	fmt.Println("\n=== Marchand ===")
	fmt.Println("1. Potion de vie (3 Pokedollars)")
	fmt.Println("2. Potion de poison (6 Pokedollars)")
	fmt.Println("3. Livre de Sort : Boule de Feu (25 Pokedollars)")
	fmt.Println("4. Fourrure de Loup (4 Pokedollars)")
	fmt.Println("5. Peau de Troll (7 Pokedollars)")
	fmt.Println("6. Cuir de Sanglier (3 Pokedollars)")
	fmt.Println("7. Plume de Corbeau (1 Pokedollar)")
	fmt.Println("8. Augmentation d’inventaire (30 Pokedollars)")
	fmt.Println("9. Retour")
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if c.Pokedollar >= 3 {
			c.Pokedollar -= 3
			c.addInventory("Potion de vie")
			fmt.Println("Vous avez obtenu : Potion de vie")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 2:
		if c.Pokedollar >= 6 {
			c.Pokedollar -= 6
			c.addInventory("Potion de poison")
			fmt.Println("Vous avez obtenu : Potion de poison")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 3:
		if c.Pokedollar >= 25 {
			c.Pokedollar -= 25
			c.addInventory("Livre de Sort : Boule de Feu")
			fmt.Println("Vous avez obtenu : Livre de Sort : Boule de Feu 🔥")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 4:
		if c.Pokedollar >= 4 {
			c.Pokedollar -= 4
			c.addInventory("Fourrure de Loup")
			fmt.Println("Vous avez obtenu : Fourrure de Loup")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 5:
		if c.Pokedollar >= 7 {
			c.Pokedollar -= 7
			c.addInventory("Peau de Troll")
			fmt.Println("Vous avez obtenu : Peau de Troll")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 6:
		if c.Pokedollar >= 3 {
			c.Pokedollar -= 3
			c.addInventory("Cuir de Sanglier")
			fmt.Println("Vous avez obtenu : Cuir de Sanglier")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 7:
		if c.Pokedollar >= 1 {
			c.Pokedollar -= 1
			c.addInventory("Plume de Corbeau")
			fmt.Println("Vous avez obtenu : Plume de Corbeau")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 8:
		if c.Pokedollar >= 30 {
			c.Pokedollar -= 30
			upgradeInventorySlot(c) // appel correct
			fmt.Println("✅ Capacité d’inventaire augmentée !")
		} else {
			fmt.Println("T'as pas la moula chef dégage 💸")
		}
	case 9:
		return
	default:
		fmt.Println("⚠️ Choix invalide !")
	}
}

// ===== POTIONS =====
func (c *Character) PoisonPot() {
	damage := 10
	duration := 3

	fmt.Println("⚠️ Vous utilisez une Potion de poison !")

	for i := 1; i <= duration; i++ {
		c.CurrentHP -= damage
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf("⏳ Tour %d : -%d HP → %d/%d PV restants\n",
			i, damage, c.CurrentHP, c.MaxHP)

		time.Sleep(1 * time.Second)

		if c.CurrentHP == 0 {
			fmt.Println("💀 Vous êtes mort à cause du poison !")
			break
		}
	}
}
