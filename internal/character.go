package internal

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Inventory map[string]int //Tableau en 2 dimensions ""Telephone" 1ere dimensions 2 c'est la deuxième dim quantité"
}

func InitCharacter(name, class string, level, maxHP, currentHP int, inventory map[string]int) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
	}
}

func (c Character) DisplayInfo() {
	fmt.Printf("=== Informations du personnage ===\n")
	fmt.Printf("Nom       : %s\n", c.Name)
	fmt.Printf("Classe    : %s\n", c.Class)
	fmt.Printf("Niveau    : %d\n", c.Level)
	fmt.Printf("PV        : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Inventaire: %v\n", c.Inventory)
	fmt.Println("=================================")
}

func (c Character) AccessInventory() {
	fmt.Println("=== Inventaire ===")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	counter := 1
	for item := range c.Inventory {
		fmt.Printf("%d. %s\n", counter, item)
		counter = counter + 1
	}

	fmt.Println("=================")
}

func (c *Character) TakePot() {
	found := false
	for item := range c.Inventory {
		if item == "Potion" {
			c.Inventory[item]--
			if c.Inventory[item] == 0 {
				delete(c.Inventory, item)
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Aucune potion disponible !")
		return
	}

	c.CurrentHP = c.CurrentHP + 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	fmt.Printf("Potion utilisée ! PV actuels :%d/%d\n", c.CurrentHP, c.MaxHP)
}
