package internal

import "fmt"

func Menu(c1 Character) {
	for {
		var choix int
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Infos du personnage")
		fmt.Println("2. Inventaire")
		fmt.Println("3. Quitter")
		fmt.Print("Choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			c1.DisplayInfo()

		case 2:
			fmt.Println("\n--- Inventaire ---")
			c1.AccessInventory()
			fmt.Println("1. Utiliser la potion")
			fmt.Println("2. Retour")
			fmt.Print("Choix : ")
			var invChoice int
			fmt.Scanln(&invChoice)

			if invChoice == 1 {
				c1.TakePot()
			} else if invChoice == 2 {
				break
			} else {
				fmt.Println("Choix valide.")
			}
		case 3:
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide !")
		}
	}
}
