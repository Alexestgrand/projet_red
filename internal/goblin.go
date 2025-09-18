package internal

import (
	"fmt"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// 👹 INITIALISATION DU GOBELIN D'ENTRAÎNEMENT
// //////////////////////////////////////////////////////////////////////////////
func InitGoblin() Monster {
	return Monster{
		Name:       "Gobelin d’entraînement",
		MaxHP:      40,
		CurrentHP:  40,
		AtkPoints:  5,
		Initiative: 5,
		Reward:     10, // 💰 Le gobelin rapporte 10 pokedollars
	}
}

func InitRaykaza() Monster {
	return Monster{
		Name:       "Raykaza le Dragon Ancien",
		MaxHP:      300,
		CurrentHP:  300,
		AtkPoints:  15,
		Initiative: 12,
		Reward:     100, // 💰 Raykaza rapporte 100 pokedollars
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

////////////////////////////////////////////////////////////////////////////////
// ⚔️ PATTERN D'ATTAQUE DE RAYKAZA
////////////////////////////////////////////////////////////////////////////////

// raykazaPattern : schéma de combat de Raykaza contre le joueur
// - Attaque normale = 100% de ses points d’attaque
// - Tous les 2 tours = Souffle draconique (150% dégâts)
// - Si PV < 50% = Rage du Dragon (inflige 200% dégâts)
// - Tous les 5 tours = Cataclysme (300% dégâts)
func raykazaPattern(c *Character, r *Monster, tour int) {
	damage := r.AtkPoints
	attaque := "Coup de griffe"

	// --- Attaque spéciale tous les 2 tours ---
	if tour%2 == 0 {
		damage = int(float64(r.AtkPoints) * 1.5)
		attaque = "🔥 Souffle draconique"
	}

	// --- Attaque ultime tous les 5 tours ---
	if tour%5 == 0 {
		damage = r.AtkPoints * 3
		attaque = "🌋 CATACLYSME"
	}

	// --- Mode Rage si PV < 50% ---
	if r.CurrentHP <= r.MaxHP/2 {
		damage = int(float64(r.AtkPoints) * 2)
		attaque = "💢 Rage du Dragon"
	}

	// --- Inflige les dégâts ---
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}

	// --- Affichage immersif ---
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	TypeWriter(fmt.Sprintf("🐉 %s utilise %s contre 🧝 %s et inflige ➝ %d dégâts !", r.Name, attaque, c.Name, damage), 15*time.Millisecond)
	TypeWriter(fmt.Sprintf("❤️ PV de %s : %d/%d", c.Name, c.CurrentHP, c.MaxHP), 15*time.Millisecond)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// --- Vérifie si le joueur est KO ---
	if c.CurrentHP == 0 {
		TypeWriter(fmt.Sprintf("💀 T'as cru t'allais battre Raykaza au niveau %d ?, retourne t'entrainé fréro", c.Level), 25*time.Millisecond)
	}
}
