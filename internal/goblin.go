package internal

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 👹 INITIALISATION DU GOBELIN D'ENTRAÎNEMENT
////////////////////////////////////////////////////////////////////////////////

// InitGoblin : crée un gobelin d’entraînement avec ses stats de base
func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d’entraînement", // Nom du monstre
		MaxHP:      40,                       // PV max
		CurrentHP:  40,                       // PV actuels (au départ = max)
		AtkPoints:  5,                        // Points d’attaque
		Initiative: 5,                        // Initiative (détermine l’ordre d’action)
	}
}

////////////////////////////////////////////////////////////////////////////////
// ⚔️ PATTERN D'ATTAQUE DU GOBELIN
////////////////////////////////////////////////////////////////////////////////

// goblinPattern : schéma de combat du gobelin contre le joueur
// - Chaque tour → inflige 100% de son attaque
// - Tous les 3 tours → inflige 200% de son attaque
func goblinPattern(c *Character, g *Monster, tour int) {
	// --- Définir les dégâts de base ---
	damage := g.AtkPoints

	// --- Tous les 3 tours → attaque spéciale à 200% ---
	if tour%3 == 0 {
		damage *= 2
		TypeWriter(fmt.Sprintf("💥 ATTENTION ! %s prépare une attaque dévastatrice !", g.Name), 20*time.Millisecond)
	}

	// --- Infliger les dégâts au joueur ---
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	// --- Affichage immersif du résultat ---
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	TypeWriter(fmt.Sprintf("👹 %s attaque 🧝 %s et inflige ➝ %d dégâts !", g.Name, c.Name, damage), 15*time.Millisecond)
	TypeWriter(fmt.Sprintf("❤️ PV de %s : %d/%d", c.Name, c.CurrentHP, c.MaxHP), 15*time.Millisecond)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// --- Vérification si le joueur est KO ---
	if c.CurrentHP == 0 {
		TypeWriter(fmt.Sprintf("💀 %s est tombé au combat...", c.Name), 25*time.Millisecond)
	}
}
