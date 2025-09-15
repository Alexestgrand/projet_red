package internal

import (
	"fmt"
	"unicode"
)

type Character struct {
	Name         string
	Class        string
	Level        int
	MaxHP        int
	CurrentHP    int
	Inventory    map[string]int
	Skills       []string
	Pokedollar   int
	MaxInventory int
	UpgradeUsed  int

	// Slots d'équipement
	Head *Headset
	Body *Torso
	Feet *Foot
}

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int) Character {
	return Character{
		Name:         name,
		Class:        class,
		Level:        level,
		MaxHP:        maxHP,
		CurrentHP:    currentHP,
		Inventory:    inventory,
		Skills:       []string{"Coup de poing"},
		Pokedollar:   100,
		MaxInventory: 10,
		UpgradeUsed:  3,
	}
}

func CharacterCreation() Character {
	name := ""
	class := ""
	maxHP := 100
	level := 1
	for {
		fmt.Println("Quel nom souhaitez-vous avoir ?")
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
			fmt.Println("⚠️ Hep hep hep ! Que des lettres stp.")
		}
	}
	for {
		fmt.Println("Quelle classe souhaitez-vous avoir parmi : Humain, Elfe, Nain ?")
		fmt.Scanln(&class)

		switch class {
		case "Humain":
			maxHP = 100
			break
		case "Elfe":
			maxHP = 80
			break
		case "Nain":
			maxHP = 120
			break
		default:
			fmt.Println("⚠️ Classe invalide. Choisissez entre Humain, Elfe ou Nain.")
			continue
		}
		break
	}

	currentHP := maxHP / 2
	inventory := map[string]int{}

	character := InitCharacter(name, class, level, maxHP, currentHP, inventory)

	fmt.Println("\nFélicitations ! 🎉 Votre personnage a bien été créé :")
	character.DisplayInfo()

	return character
}

func (c Character) DisplayInfo() {
	fmt.Printf("=== Informations du personnage ===\n")
	fmt.Printf("Nom       : %s\n", c.Name)
	fmt.Printf("Classe    : %s\n", c.Class)
	fmt.Printf("Niveau    : %d\n", c.Level)
	fmt.Printf("PV        : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Inventaire: %v\n", c.Inventory)
	fmt.Printf("Compétences: ")
	if len(c.Skills) == 0 {
		fmt.Println("Aucune compétence apprise.")
	} else {
		for _, skill := range c.Skills {
			fmt.Printf("%s ", skill)
		}
		fmt.Println()
	}
	fmt.Println("=================================")
}

func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		fmt.Println("Le joueur est mort. 💀")
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("Le joueur est ressuscité avec %d points de vie.\n", c.CurrentHP)
	}
}
