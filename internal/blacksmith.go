package internal

import (
	"fmt"
	"time"
)

// ForgeronInterface : permet au joueur de fabriquer des équipements
// en échange de pokedollars et de ressources précises.
func (c *Character) ForgeronInterface() {
	var choix int

	// Menu principal du forgeron
	TypeWriter("\n=================== ⚒️  FORGERON ⚒️ ===================", 15*time.Millisecond)
	TypeWriter("Bienvenue, aventurier ! Que souhaites-tu forger aujourd'hui ?", 15*time.Millisecond)
	TypeWriter("-----------------------------------------------------------", 15*time.Millisecond)
	TypeWriter("1. 🎩 Chapeau de l’aventurier (5💰) ; nécessite 1x Plume de Corbeau + 1x Cuir de Sanglier", 15*time.Millisecond)
	TypeWriter("2. 👕 Tunique de l’aventurier (10💰) ; nécessite 2x Fourrures de Loup + 1x Peau de Troll", 15*time.Millisecond)
	TypeWriter("3. 👢 Bottes de l’aventurier (7💰) ; nécessite 1x Fourrure de Loup + 1x Cuir de Sanglier", 15*time.Millisecond)
	TypeWriter("4. 🚪 Quitter", 15*time.Millisecond)
	TypeWriter("===========================================================", 15*time.Millisecond)
	fmt.Print("👉 Fais ton choix : ")
	fmt.Scanln(&choix)

	switch choix {
	// === FABRICATION DU CHAPEAU ===
	case 1:
		if c.Pokedollar < 5 {
			TypeWriter("💸 Tu n’as pas assez de pokedollars pour forger ce chapeau !", 20*time.Millisecond)
			return
		}
		if c.Inventory["Plume de Corbeau"] < 1 {
			TypeWriter("⚠️ Il te manque une Plume de Corbeau.", 20*time.Millisecond)
			return
		}
		if c.Inventory["Cuir de Sanglier"] < 1 {
			TypeWriter("⚠️ Il te manque un Cuir de Sanglier.", 20*time.Millisecond)
			return
		}

		// Déductions
		c.Pokedollar -= 5
		c.removeInventory("Plume de Corbeau", 1)
		c.removeInventory("Cuir de Sanglier", 1)
		c.addInventory("Chapeau de l’aventurier")

		// Confirmation
		TypeWriter("✅ Tu as forgé un magnifique 🎩 Chapeau de l’aventurier !", 20*time.Millisecond)

	// === FABRICATION DE LA TUNIQUE ===
	case 2:
		if c.Pokedollar < 10 {
			TypeWriter("💸 Tu n’as pas assez de pokedollars pour cette tunique !", 20*time.Millisecond)
			return
		}
		if c.Inventory["Fourrure de Loup"] < 2 {
			TypeWriter("⚠️ Il te faut 2 Fourrures de Loup.", 20*time.Millisecond)
			return
		}
		if c.Inventory["Peau de Troll"] < 1 {
			TypeWriter("⚠️ Il te manque une Peau de Troll.", 20*time.Millisecond)
			return
		}

		// Déductions
		c.Pokedollar -= 10
		c.removeInventory("Fourrure de Loup", 2)
		c.removeInventory("Peau de Troll", 1)
		c.addInventory("Tunique de l’aventurier")

		// Confirmation
		TypeWriter("✅ Tu as forgé une solide 👕 Tunique de l’aventurier !", 20*time.Millisecond)

	// === FABRICATION DES BOTTES ===
	case 3:
		if c.Pokedollar < 7 {
			TypeWriter("💸 Tu n’as pas assez de pokedollars pour ces bottes !", 20*time.Millisecond)
			return
		}
		if c.Inventory["Fourrure de Loup"] < 1 {
			TypeWriter("⚠️ Il te manque une Fourrure de Loup.", 20*time.Millisecond)
			return
		}
		if c.Inventory["Cuir de Sanglier"] < 1 {
			TypeWriter("⚠️ Il te manque un Cuir de Sanglier.", 20*time.Millisecond)
			return
		}

		// Déductions
		c.Pokedollar -= 7
		c.removeInventory("Fourrure de Loup", 1)
		c.removeInventory("Cuir de Sanglier", 1)
		c.addInventory("Bottes de l’aventurier")

		// Confirmation
		TypeWriter("✅ Tu as forgé des robustes 👢 Bottes de l’aventurier !", 20*time.Millisecond)

	// === QUITTER LE FORGERON ===
	case 4:
		TypeWriter("👋 À bientôt, aventurier !", 20*time.Millisecond)
		return

	// === CAS PAR DÉFAUT ===
	default:
		TypeWriter("⚠️ Choix invalide, essaie encore !", 20*time.Millisecond)
	}
}
