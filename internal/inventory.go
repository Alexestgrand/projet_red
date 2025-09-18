package internal

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 🎒 GESTION DE L'INVENTAIRE
////////////////////////////////////////////////////////////////////////////////

// ➕ Ajout d’un objet à l’inventaire
func (c *Character) addInventory(item string) {
	if len(c.Inventory) >= c.MaxInventory && c.Inventory[item] == 0 {
		TypeWriter("⚠️ Inventaire plein, objet non ajouté !", 20*time.Millisecond)
		return
	}

	c.Inventory[item]++
	TypeWriter(fmt.Sprintf("📦 Vous avez ajouté : %s", item), 15*time.Millisecond)
}

// ➖ Retrait d’un objet de l’inventaire
func (c *Character) removeInventory(item string, qty int) {
	if c.Inventory[item] > qty {
		c.Inventory[item] -= qty
	} else {
		delete(c.Inventory, item)
	}
}

// 📜 Gestion complète de l’inventaire (hors combat)
func (c *Character) AccessInventory() {
	for {
		fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("🎒 INVENTAIRE")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("1. Voir les objets")
		fmt.Println("2. Utiliser un objet")
		fmt.Println("3. Retour")
		fmt.Print("👉 Choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		// === VOIR OBJETS ===
		case 1:
			if len(c.Inventory) == 0 {
				TypeWriter("Votre inventaire est vide.", 20*time.Millisecond)
				continue
			}
			TypeWriter("📦 Contenu de l’inventaire :", 15*time.Millisecond)
			for item, qty := range c.Inventory {
				fmt.Printf("   ➝ %s (x%d)\n", item, qty)
			}

		// === UTILISER OBJET ===
		case 2:
			if len(c.Inventory) == 0 {
				TypeWriter("⚠️ Votre inventaire est vide.", 20*time.Millisecond)
				continue
			}

			// Liste des objets
			keys := make([]string, 0, len(c.Inventory))
			for item := range c.Inventory {
				keys = append(keys, item)
			}
			for i, item := range keys {
				fmt.Printf("%d. %s (x%d)\n", i+1, item, c.Inventory[item])
			}

			// Choix
			var choice int
			fmt.Print("👉 Choix (0 pour annuler) : ")
			fmt.Scanln(&choice)

			if choice == 0 {
				continue
			}
			if choice <= 0 || choice > len(keys) {
				TypeWriter("⚠️ Choix invalide !", 20*time.Millisecond)
				continue
			}

			selected := keys[choice-1]

			// Appliquer effets
			switch selected {
			case "Potion de vie":
				c.TakePot()
				c.removeInventory(selected, 1)

			case "Potion de mana":
				c.TakeManaPot()
				c.removeInventory(selected, 1)

			case "Livre de Sort : Boule de Feu":
				c.LearnSpell("Boule de feu")
				c.removeInventory(selected, 1)
				TypeWriter("📘 Vous avez appris un nouveau sort : Boule de feu 🔥", 15*time.Millisecond)

			case "Livre de Sort : Soin léger":
				c.LearnSpell("Soin léger")
				c.removeInventory(selected, 1)
				TypeWriter("📘 Vous avez appris un nouveau sort : Soin léger ✨", 15*time.Millisecond)

			default:
				TypeWriter("⚠️ Cet objet ne peut pas être utilisé hors combat.", 20*time.Millisecond)
			}

		// === RETOUR ===
		case 3:
			return

		default:
			TypeWriter("⚠️ Choix invalide !", 20*time.Millisecond)
		}
	}
}

// //////////////////////////////////////////////////////////////////////////////
// ⚔️ UTILISATION D'UN OBJET EN COMBAT
// //////////////////////////////////////////////////////////////////////////////
func (c *Character) UseItem(monster *Monster) {
	if len(c.Inventory) == 0 {
		TypeWriter("⚠️ Votre inventaire est vide.", 20*time.Millisecond)
		return
	}

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎒 INVENTAIRE (Combat)")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Génère une liste d’objets utilisables
	keys := make([]string, 0, len(c.Inventory))
	for item := range c.Inventory {
		keys = append(keys, item)
	}
	for i, item := range keys {
		fmt.Printf("%d. %s (x%d)\n", i+1, item, c.Inventory[item])
	}

	// Choix du joueur
	var choice int
	fmt.Print("👉 Choix : ")
	fmt.Scanln(&choice)

	if choice <= 0 || choice > len(keys) {
		TypeWriter("⚠️ Choix invalide !", 20*time.Millisecond)
		return
	}

	selected := keys[choice-1]

	// Effets des objets
	switch selected {
	case "Potion de vie":
		c.TakePot()
		c.removeInventory(selected, 1)

	case "Potion de poison":
		c.PoisonPot()
		c.removeInventory(selected, 1)

	case "Livre de Sort : Boule de Feu":
		c.LearnSpell("Boule de feu")
		c.removeInventory(selected, 1)
		TypeWriter("📘 Vous avez appris un nouveau sort : Boule de feu 🔥", 15*time.Millisecond)

	case "Livre de Sort : Soin léger":
		c.LearnSpell("Soin léger")
		c.removeInventory(selected, 1)
		TypeWriter("📘 Vous avez appris un nouveau sort : Soin léger ✨", 15*time.Millisecond)

	case "Potion de mana":
		c.TakeManaPot()
		c.removeInventory(selected, 1)

	default:
		TypeWriter("⚠️ Cet objet ne peut pas être utilisé en combat.", 20*time.Millisecond)
	}
}
