package internal

import (
	"fmt"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// 🛒 INTERFACE DU MARCHAND
// //////////////////////////////////////////////////////////////////////////////
func (c *Character) MarchantInterface() {
	var choix int

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	TypeWriter("🛒 Bienvenue chez le Marchand, aventurier !", 20*time.Millisecond)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("1. 🧪 Potion de vie (3💰)")
	fmt.Println("2. ☠️ Potion de poison (6💰)")
	fmt.Println("3. 📘 Livre de Sort : Boule de Feu (25💰)")
	fmt.Println("4. 🐺 Fourrure de Loup (4💰)")
	fmt.Println("5. 👹 Peau de Troll (7💰)")
	fmt.Println("6. 🐗 Cuir de Sanglier (3💰)")
	fmt.Println("7. 🪶 Plume de Corbeau (1💰)")
	fmt.Println("8. 🎒 Augmentation d’inventaire (+10 slots, 30💰)")
	fmt.Println("9. 🚪 Quitter la boutique")
	fmt.Println("10. ✨ Sort de soin (20💰)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("💰 Vous avez actuellement : %d pokedollars\n", c.Pokedollar)
	fmt.Print("👉 Choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if c.Pokedollar >= 3 {
			c.Pokedollar -= 3
			c.addInventory("Potion de vie")
			TypeWriter("✅ Vous avez obtenu : Potion de vie 🧪", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 2:
		if c.Pokedollar >= 6 {
			c.Pokedollar -= 6
			c.addInventory("Potion de poison")
			TypeWriter("✅ Vous avez obtenu : Potion de poison ☠️", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 3:
		if c.Pokedollar >= 25 {
			c.Pokedollar -= 25
			c.addInventory("Livre de Sort : Boule de Feu")
			TypeWriter("📘 Vous avez acheté : Livre de Sort - Boule de Feu 🔥", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 4:
		if c.Pokedollar >= 4 {
			c.Pokedollar -= 4
			c.addInventory("Fourrure de Loup")
			TypeWriter("✅ Vous avez obtenu : Fourrure de Loup 🐺", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 5:
		if c.Pokedollar >= 7 {
			c.Pokedollar -= 7
			c.addInventory("Peau de Troll")
			TypeWriter("✅ Vous avez obtenu : Peau de Troll 👹", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 6:
		if c.Pokedollar >= 3 {
			c.Pokedollar -= 3
			c.addInventory("Cuir de Sanglier")
			TypeWriter("✅ Vous avez obtenu : Cuir de Sanglier 🐗", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 7:
		if c.Pokedollar >= 1 {
			c.Pokedollar -= 1
			c.addInventory("Plume de Corbeau")
			TypeWriter("✅ Vous avez obtenu : Plume de Corbeau 🪶", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 8:
		if c.Pokedollar >= 30 {
			c.Pokedollar -= 30
			upgradeInventorySlot(c)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 10:
		if c.Pokedollar >= 20 {
			c.Pokedollar -= 20
			c.LearnSpell("Soin léger")
			TypeWriter("📘 Nouveau sort appris : Soin léger ✨", 15*time.Millisecond)
		} else {
			TypeWriter("💸 Pas assez de pokedollars !", 15*time.Millisecond)
		}

	case 9:
		TypeWriter("👋 Merci pour votre visite, aventurier !", 15*time.Millisecond)
		return

	default:
		TypeWriter("⚠️ Choix invalide !", 15*time.Millisecond)
	}
}

// //////////////////////////////////////////////////////////////////////////////
// ☠️ Potion de poison
// //////////////////////////////////////////////////////////////////////////////
func (c *Character) PoisonPot() {
	damage := 10
	duration := 3

	TypeWriter("⚠️ Vous utilisez une Potion de poison !", 20*time.Millisecond)

	for i := 1; i <= duration; i++ {
		c.CurrentHP -= damage
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		fmt.Printf("⏳ Tour %d : -%d HP → %d/%d PV restants\n",
			i, damage, c.CurrentHP, c.MaxHP)

		time.Sleep(1 * time.Second)

		if c.CurrentHP == 0 {
			TypeWriter("💀 Vous êtes mort à cause du poison !", 20*time.Millisecond)
			break
		}
	}
}
