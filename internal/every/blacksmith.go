package internal

import "fmt"

func (c *Character) ForgeronInterface() {
	var choix int
	fmt.Println("\n=== Forgeron ===")
	fmt.Println("1. Chapeau de l’aventurier (5 Pokedollars) ; nécessite 1 Plume de Corbeau + 1 Cuir de Sanglier")
	fmt.Println("2. Tunique de l’aventurier (10 Pokedollars) ; nécessite 2 Fourrures de Loup + 1 Peau de Troll")
	fmt.Println("3. Bottes de l’aventurier (7 Pokedollars) ; nécessite 1 Fourrure de Loup + 1 Cuir de Sanglier")
	fmt.Println("4. Retour")
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1: // Chapeau
		if c.Pokedollar < 5 {
			fmt.Println("💸 Pas assez de pokedollars.")
			return
		}
		if c.Inventory["Plume de Corbeau"] < 1 {
			fmt.Println("⚠️ Il vous manque 1 Plume de Corbeau.")
			return
		}
		if c.Inventory["Cuir de Sanglier"] < 1 {
			fmt.Println("⚠️ Il vous manque 1 Cuir de Sanglier.")
			return
		}

		c.Pokedollar -= 5
		c.removeInventory("Plume de Corbeau", 1)
		c.removeInventory("Cuir de Sanglier", 1)
		c.addInventory("Chapeau de l’aventurier")
		fmt.Println("✅ Vous avez fabriqué un Chapeau de l’aventurier 🎩")

	case 2: // Tunique
		if c.Pokedollar < 10 {
			fmt.Println("💸 Pas assez de pokedollars.")
			return
		}
		if c.Inventory["Fourrure de Loup"] < 2 {
			fmt.Println("⚠️ Il vous faut 2 Fourrures de Loup.")
			return
		}
		if c.Inventory["Peau de Troll"] < 1 {
			fmt.Println("⚠️ Il vous manque 1 Peau de Troll.")
			return
		}

		c.Pokedollar -= 10
		c.removeInventory("Fourrure de Loup", 2)
		c.removeInventory("Peau de Troll", 1)
		c.addInventory("Tunique de l’aventurier")
		fmt.Println("✅ Vous avez fabriqué une Tunique de l’aventurier 👕")

	case 3: // Bottes
		if c.Pokedollar < 7 {
			fmt.Println("💸 Pas assez de pokedollars.")
			return
		}
		if c.Inventory["Fourrure de Loup"] < 1 {
			fmt.Println("⚠️ Il vous manque 1 Fourrure de Loup.")
			return
		}
		if c.Inventory["Cuir de Sanglier"] < 1 {
			fmt.Println("⚠️ Il vous manque 1 Cuir de Sanglier.")
			return
		}

		c.Pokedollar -= 7
		c.removeInventory("Fourrure de Loup", 1)
		c.removeInventory("Cuir de Sanglier", 1)
		c.addInventory("Bottes de l’aventurier")
		fmt.Println("✅ Vous avez fabriqué des Bottes de l’aventurier 👢")

	case 4:
		return

	default:
		fmt.Println("⚠️ Choix invalide !")
	}
}
