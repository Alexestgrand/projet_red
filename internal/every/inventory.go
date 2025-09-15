package internal

import "fmt"

// ===== INVENTAIRE =====
func (c *Character) addInventory(item string) {
	if len(c.Inventory) >= c.MaxInventory && c.Inventory[item] == 0 {
		fmt.Println("⚠️ Inventaire plein, objet non ajouté !")
		return
	}

	c.Inventory[item]++
	fmt.Printf("Vous avez ajouté : %s\n", item)
}

func (c *Character) removeInventory(item string, qty int) {
	if c.Inventory[item] > qty {
		c.Inventory[item] -= qty
	} else {
		delete(c.Inventory, item)
	}
}
func (c Character) AccessInventory() {
	fmt.Println("=== Inventaire ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	counter := 1
	for item := range c.Inventory {
		fmt.Printf("%d. %s\n", counter, item)
		counter = counter + 1
	}

	fmt.Println("=================")
}
