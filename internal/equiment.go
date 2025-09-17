package internal

import (
	"fmt"
	"time"
)

// upgradeInventorySlot : augmente la capacité de l’inventaire du joueur.
// ➡️ Chaque utilisation ajoute +10 emplacements et réduit le nombre d’améliorations restantes.
// ⚠️ Limité à 3 améliorations maximum (voir UpgradeUsed dans Character).
func upgradeInventorySlot(c *Character) {
	// Vérifie si le joueur peut encore améliorer son inventaire
	if c.UpgradeUsed <= 0 {
		TypeWriter("❌ Impossible d’augmenter l’inventaire, vous avez atteint la limite maximale.", 20*time.Millisecond)
		return
	}

	// Applique l’augmentation
	c.MaxInventory += 10
	c.UpgradeUsed--

	// Message immersif pour le joueur
	TypeWriter("\n====================== ⚒️ INVENTAIRE AMÉLIORÉ ======================", 15*time.Millisecond)
	TypeWriter(fmt.Sprintf("🎒 Nouvelle capacité maximale : %d objets", c.MaxInventory), 20*time.Millisecond)
	TypeWriter(fmt.Sprintf("✨ Utilisations restantes de l’augmentation : %d", c.UpgradeUsed), 20*time.Millisecond)
	TypeWriter("=================================================================", 15*time.Millisecond)
}
