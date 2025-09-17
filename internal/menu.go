package internal

import (
	"fmt"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// 🏰 MENU PRINCIPAL
// //////////////////////////////////////////////////////////////////////////////
func Menu(c1 *Character) {
	for {
		var choix int

		fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		TypeWriter("🏰 MENU PRINCIPAL", 20*time.Millisecond)
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println("1. 📋 Infos du personnage")
		fmt.Println("2. 🎒 Inventaire")
		fmt.Println("3. 🛒 Marchand")
		fmt.Println("4. ⚒️ Forgeron")
		fmt.Println("5. 🛡️ Equipement")
		fmt.Println("6. ⚔️ Entraînement")
		fmt.Println("7. 🚪 Quitter")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Print("👉 Choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			c1.DisplayInfo()
		case 2:
			c1.AccessInventory()
		case 3:
			c1.MarchantInterface()
		case 4:
			c1.ForgeronInterface()
		case 5:
			c1.EquipementInterface()
		case 6:
			trainingFight(c1) // 👈 lancement du combat
		case 7:
			TypeWriter("👋 Au revoir, aventurier !", 20*time.Millisecond)
			return
		default:
			TypeWriter("⚠️ Choix invalide, réessaie !", 20*time.Millisecond)
		}
	}
}
