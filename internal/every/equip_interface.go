package internal

import "fmt"

// Interface commune

type Equipment interface {
	GetName() string
	GetBonusHP() int
}

// ===== Casque =====
type Headset struct {
	Name    string
	BonusHP int
}

func (h Headset) GetName() string { return h.Name }
func (h Headset) GetBonusHP() int { return h.BonusHP }

// ===== Torse =====
type Torso struct {
	Name    string
	BonusHP int
}

func (t Torso) GetName() string { return t.Name }
func (t Torso) GetBonusHP() int { return t.BonusHP }

// ===== Bottes =====
type Foot struct {
	Name    string
	BonusHP int
}

func (f Foot) GetName() string { return f.Name }
func (f Foot) GetBonusHP() int { return f.BonusHP }

func (c *Character) updateMaxHP() {
	// PV de base selon la classe
	baseHP := 0
	switch c.Class {
	case "Humain":
		baseHP = 100
	case "Elfe":
		baseHP = 80
	case "Nain":
		baseHP = 120
	}

	// Ajout des bonus des équipements
	if c.Head != nil {
		baseHP += c.Head.BonusHP
	}
	if c.Body != nil {
		baseHP += c.Body.BonusHP
	}
	if c.Feet != nil {
		baseHP += c.Feet.BonusHP
	}

	// Mise à jour
	oldMax := c.MaxHP
	c.MaxHP = baseHP

	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
		fmt.Println("⚠️ Vos PV ont été réduits car vous avez retiré un équipement.")
	}

	// Feedback si les PV max ont augmenté
	if c.MaxHP > oldMax {
		fmt.Printf("💪 Vos PV maximum augmentent à %d !\n", c.MaxHP)
	}
}

func (c *Character) EquipementInterface() {
	choix := ""
	fmt.Println("\n=== Gestion de l'équipement ===")
	fmt.Println("Que voulez-vous équiper ?")
	fmt.Println("Chapeau / Tunique / Bottes")
	fmt.Print("Choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case "Chapeau":
		if c.Inventory["Chapeau de l’aventurier"] > 0 {
			// Retourne l’ancien casque dans l’inventaire
			if c.Head != nil {
				c.addInventory(c.Head.GetName())
			}
			// Conversion string -> Headset struct
			c.Head = CreateEquipment("Chapeau de l’aventurier").(*Headset)
			// Retrait de l’inventaire
			c.removeInventory("Chapeau de l’aventurier", 1)
			// Mise à jour PV max
			c.updateMaxHP()
			fmt.Println("✅ Vous avez équipé le Chapeau de l’aventurier 🎩")
		} else {
			fmt.Println("⚠️ Vous n’avez pas de Chapeau de l’aventurier.")
		}

	case "Tunique":
		if c.Inventory["Tunique de l’aventurier"] > 0 {
			if c.Body != nil {
				c.addInventory(c.Body.GetName())
			}
			c.Body = CreateEquipment("Tunique de l’aventurier").(*Torso)
			c.removeInventory("Tunique de l’aventurier", 1)
			c.updateMaxHP()
			fmt.Println("✅ Vous avez équipé la Tunique de l’aventurier 👕")
		} else {
			fmt.Println("⚠️ Vous n’avez pas de Tunique de l’aventurier.")
		}

	case "Bottes":
		if c.Inventory["Bottes de l’aventurier"] > 0 {
			if c.Feet != nil {
				c.addInventory(c.Feet.GetName())
			}
			c.Feet = CreateEquipment("Bottes de l’aventurier").(*Foot)
			c.removeInventory("Bottes de l’aventurier", 1)
			c.updateMaxHP()
			fmt.Println("✅ Vous avez équipé les Bottes de l’aventurier 👢")
		} else {
			fmt.Println("⚠️ Vous n’avez pas de Bottes de l’aventurier.")
		}

	default:
		fmt.Println("⚠️ Choix invalide !")
	}
}

func CreateEquipment(name string) interface{} {
	switch name {
	case "Chapeau de l’aventurier":
		return &Headset{Name: name, BonusHP: 10}
	case "Tunique de l’aventurier":
		return &Torso{Name: name, BonusHP: 25}
	case "Bottes de l’aventurier":
		return &Foot{Name: name, BonusHP: 15}
	default:
		return nil
	}
}
