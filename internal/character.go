package internal

import (
	"fmt"
	"time"
	"unicode"
)

// ===============================
// 🧙 STRUCTURE DU PERSONNAGE
// ===============================
type Character struct {
	Name         string         // Nom du personnage
	Class        string         // Classe choisie (Humain, Elfe, Nain)
	Level        int            // Niveau actuel
	MaxHP        int            // Points de vie maximum
	CurrentHP    int            // Points de vie actuels
	Inventory    map[string]int // Inventaire (nom de l’objet → quantité)
	Skills       []string       // Liste des compétences connues
	Pokedollar   int
	Initiative   int // Initiative (détermine qui commence en combat)
	MaxInventory int // Capacité maximale de l’inventaire
	UpgradeUsed  int // Nombre d’augmentations d’inventaire restantes

	// Expérience
	Exp      int // Exp actuelle
	ExpToLvl int // Exp nécessaire pour monter de niveau

	// Mana
	Mana    int
	MaxMana int

	// Slots d’équipement
	Head *Headset // Casque / Chapeau
	Body *Torso   // Armure / Tunique
	Feet *Foot    // Bottes
}

// ===============================
// 🖋️ Machine à écrire
// ===============================
func TypeWriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

// ===============================
// ⚙️ INITIALISATION6
// ===============================
func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int) Character {
	return Character{
		Name:         name,
		Class:        class,
		Level:        level,
		MaxHP:        maxHP,
		CurrentHP:    currentHP,
		Inventory:    inventory,
		Skills:       []string{"Coup de poing"}, // le joueur commence avec une seule attaque
		Pokedollar:   100,
		MaxInventory: 10,
		UpgradeUsed:  3,
		Mana:         30,
		MaxMana:      30,
		Exp:          0,
		ExpToLvl:     50, // niveau 1 → il faut 50 XP pour monter
	}
}

// ===============================
// ✨ CREATION DU PERSONNAGE
// ===============================
func CharacterCreation() Character {
	name := ""
	class := ""
	maxHP := 100
	level := 1

	// 🔹 Choix du nom
	for {
		TypeWriter("Bienvenue dans notre monde fantastique Jeune Padawan !", 20*time.Millisecond)
		TypeWriter("Comment te prénommes-tu donc ?", 20*time.Millisecond)
		fmt.Scanln(&name)

		valid := true
		for _, letter := range name {
			if !unicode.IsLetter(letter) {
				valid = false
				break
			}
		}

		if valid {
			break
		} else {
			TypeWriter("⚠️ T'as cru que c'était chez ta daronne ? Que des lettres seulement !", 20*time.Millisecond)
		}
	}

	// 🔹 Choix de la classe
	for {
		TypeWriter("Quelle classe souhaites-tu avoir ?", 20*time.Millisecond)
		TypeWriter("👉 Humain (100 PV), Elfe (80 PV), Nain (120 PV)", 20*time.Millisecond)
		fmt.Scanln(&class)

		switch class {
		case "Humain":
			maxHP = 100
		case "Elfe":
			maxHP = 80
		case "Nain":
			maxHP = 120
		default:
			TypeWriter("⚠️ Classe invalide. Choisis entre Humain, Elfe ou Nain.", 20*time.Millisecond)
			continue
		}
		break
	}

	// PV actuels = 50% des PV max
	currentHP := maxHP / 2

	// Inventaire vide
	inventory := map[string]int{}

	// Création du personnage
	character := InitCharacter(name, class, level, maxHP, currentHP, inventory)

	TypeWriter("\n🎉 Félicitations ! Votre personnage a bien été créé :", 15*time.Millisecond)
	character.DisplayInfo()

	return character
}

// ===============================
// ⭐ EXPERIENCE
// ===============================
func (c *Character) GainExp(amount int) {
	TypeWriter(fmt.Sprintf("⭐ %s gagne %d points d’expérience !", c.Name, amount), 15*time.Millisecond)
	c.Exp += amount

	for c.Exp >= c.ExpToLvl {
		c.Level++
		c.Exp -= c.ExpToLvl
		c.ExpToLvl += 25 // Chaque niveau demande +25 XP en plus
		c.MaxHP += 10
		c.CurrentHP = c.MaxHP
		TypeWriter(fmt.Sprintf("🎉 %s passe au niveau %d ! ❤️ PV max = %d", c.Name, c.Level, c.MaxHP), 20*time.Millisecond)
	}

	c.displayExpBar()
}

func (c *Character) displayExpBar() {
	totalBars := 20
	filledBars := (c.Exp * totalBars) / c.ExpToLvl

	bar := ""
	for i := 0; i < totalBars; i++ {
		if i < filledBars {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	fmt.Printf("🔸 XP [%s] %d/%d\n", bar, c.Exp, c.ExpToLvl)
}

// ===============================
// 📋 INFOS DU PERSONNAGE
// ===============================
func (c Character) DisplayInfo() {
	fmt.Println("\n========== 📋 INFOS DU PERSONNAGE ==========")
	fmt.Printf("🧑 Nom       : %s\n", c.Name)
	fmt.Printf("🏹 Classe    : %s\n", c.Class)
	fmt.Printf("⭐ Niveau    : %d\n", c.Level)
	fmt.Printf("❤️ PV        : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("💧 Mana      : %d/%d\n", c.Mana, c.MaxMana)
	fmt.Printf("💰 Argent    : %d pokedollars\n", c.Pokedollar)

	// Inventaire
	fmt.Printf("🎒 Inventaire: %v\n", c.Inventory)

	// Compétences
	fmt.Print("🔥 Compétences : ")
	if len(c.Skills) == 0 {
		fmt.Println("Aucune compétence apprise.")
	} else {
		for _, skill := range c.Skills {
			fmt.Printf("%s ", skill)
		}
		fmt.Println()
	}
	fmt.Println("============================================")
}

// ===============================
// 💀 MORT ET RESURRECTION
// ===============================
func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		TypeWriter("\n💀 Le joueur est mort...", 30*time.Millisecond)
		c.CurrentHP = c.MaxHP / 2
		TypeWriter(fmt.Sprintf("✨ Le joueur est ressuscité avec %d PV.", c.CurrentHP), 30*time.Millisecond)
	}
}
