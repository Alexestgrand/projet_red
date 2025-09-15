package internal

import (
	"fmt"
)

// ===== MENU PRINCIPAL =====
func Menu(c1 *Character) {
	for {
		var choix int
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Infos du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Equipement")
		fmt.Println("6. Quitter")
		fmt.Print("Choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			c1.DisplayInfo()

		case 2:
			fmt.Println("\n--- Inventaire ---")
			c1.AccessInventory()
			fmt.Println("1. Utiliser un objet")
			fmt.Println("2. Retour")
			fmt.Print("Choix : ")
			var invChoice int
			fmt.Scanln(&invChoice)

			if invChoice == 1 {
				if len(c1.Inventory) == 0 {
					fmt.Println("⚠️ Votre inventaire est vide.")
					continue
				}

				fmt.Println("\nQuel objet voulez-vous utiliser ?")
				keys := make([]string, 0, len(c1.Inventory))
				for item := range c1.Inventory {
					keys = append(keys, item)
				}
				for i, item := range keys {
					fmt.Printf("%d. %s (x%d)\n", i+1, item, c1.Inventory[item])
				}
				var itemChoice int
				fmt.Print("Choix : ")
				fmt.Scanln(&itemChoice)

				if itemChoice > 0 && itemChoice <= len(keys) {
					selected := keys[itemChoice-1]
					switch selected {
					case "Potion de vie":
						c1.TakePot()
						c1.removeInventory(selected, 1)

					case "Potion de poison":
						c1.PoisonPot()
						c1.removeInventory(selected, 1)

					case "Livre de Sort : Boule de Feu":
						c1.SpellBook()
						c1.removeInventory(selected, 1)

					default:
						fmt.Println("⚠️ Cet objet ne peut pas être utilisé.")
					}
				} else {
					fmt.Println("⚠️ Choix invalide.")
				}
			}

		case 3:
			c1.MarchantInterface()

		case 4:
			c1.ForgeronInterface()

		case 5:
			c1.EquipementInterface()

		case 6:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("⚠️ Choix invalide !")
		}
	}
}
