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

// 📜 Affichage du contenu de l’inventaire
func (c Character) AccessInventory() {
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎒 INVENTAIRE")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if len(c.Inventory) == 0 {
		TypeWriter("Votre inventaire est vide.", 20*time.Millisecond)
		return
	}

	counter := 1
	for item, qty := range c.Inventory {
		TypeWriter(fmt.Sprintf("%d. %s (x%d)", counter, item, qty), 15*time.Millisecond)
		counter++
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
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

	default:
		TypeWriter("⚠️ Cet objet ne peut pas être utilisé en combat.", 20*time.Millisecond)
	}
}
