package internal

import "fmt"

func upgradeInventorySlot(c *Character) {
	if c.UpgradeUsed <= 0 {
		fmt.Println("⚠️ Impossible d’augmenter l’inventaire, limite atteinte.")
		return
	}

	c.MaxInventory += 10
	c.UpgradeUsed--
	fmt.Printf("🎒 Inventaire augmenté ! Capacité max = %d (utilisations restantes : %d)\n",
		c.MaxInventory, c.UpgradeUsed)
}
