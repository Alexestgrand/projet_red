package internal

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 🎭 INTERFACE COMMUNE DES ÉQUIPEMENTS
////////////////////////////////////////////////////////////////////////////////

// Tous les équipements doivent pouvoir fournir leur nom et leur bonus de PV
type Equipment interface {
	GetName() string
	GetBonusHP() int
}

////////////////////////////////////////////////////////////////////////////////
// 🪖 STRUCTURES DES ÉQUIPEMENTS
////////////////////////////////////////////////////////////////////////////////

// ===== Casque (slot : Head) =====
type Headset struct {
	Name    string
	BonusHP int
}

func (h Headset) GetName() string { return h.Name }
func (h Headset) GetBonusHP() int { return h.BonusHP }

// ===== Torse (slot : Body) =====
type Torso struct {
	Name    string
	BonusHP int
}

func (t Torso) GetName() string { return t.Name }
func (t Torso) GetBonusHP() int { return t.BonusHP }

// ===== Bottes (slot : Feet) =====
type Foot struct {
	Name    string
	BonusHP int
}

func (f Foot) GetName() string { return f.Name }
func (f Foot) GetBonusHP() int { return f.BonusHP }

////////////////////////////////////////////////////////////////////////////////
// ❤️ MISE À JOUR DES POINTS DE VIE EN FONCTION DES ÉQUIPEMENTS
////////////////////////////////////////////////////////////////////////////////

func (c *Character) updateMaxHP() {
	// --- Définition des PV de base selon la classe du personnage ---
	baseHP := 0
	switch c.Class {
	case "Humain":
		baseHP = 100
	case "Elfe":
		baseHP = 80
	case "Nain":
		baseHP = 120
	}

	// --- Ajout des bonus des équipements portés ---
	if c.Head != nil {
		baseHP += c.Head.GetBonusHP()
	}
	if c.Body != nil {
		baseHP += c.Body.GetBonusHP()
	}
	if c.Feet != nil {
		baseHP += c.Feet.GetBonusHP()
	}

	// --- Mise à jour des valeurs ---
	oldMax := c.MaxHP
	c.MaxHP = baseHP

	// Ajustement si les PV actuels dépassent les PV max
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
		TypeWriter("⚠️ Vos PV ont été réduits car vous avez retiré un équipement.", 20*time.Millisecond)
	}

	// Feedback si les PV max augmentent
	if c.MaxHP > oldMax {
		TypeWriter(fmt.Sprintf("💪 Vos PV maximum augmentent à %d !", c.MaxHP), 20*time.Millisecond)
	}
}

////////////////////////////////////////////////////////////////////////////////
// 🛡️ INTERFACE UTILISATEUR : MENU D'ÉQUIPEMENT
////////////////////////////////////////////////////////////////////////////////

func (c *Character) EquipementInterface() {
	choix := ""
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	TypeWriter("⚔️  GESTION DE L'ÉQUIPEMENT", 15*time.Millisecond)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	TypeWriter("Que voulez-vous équiper ?", 15*time.Millisecond)
	TypeWriter("➡️  Chapeau / Tunique / Bottes", 15*time.Millisecond)
	fmt.Print("👉 Votre choix : ")
	fmt.Scanln(&choix)

	switch choix {
	// --- Casque ---
	case "Chapeau":
		if c.Inventory["Chapeau de l’aventurier"] > 0 {
			if c.Head != nil {
				c.addInventory(c.Head.GetName())
			}
			c.Head = CreateEquipment("Chapeau de l’aventurier").(*Headset)
			c.removeInventory("Chapeau de l’aventurier", 1)
			c.updateMaxHP()
			TypeWriter("✅ Vous avez équipé le Chapeau de l’aventurier 🎩", 15*time.Millisecond)
		} else {
			TypeWriter("⚠️ Vous n’avez pas de Chapeau de l’aventurier.", 15*time.Millisecond)
		}

	// --- Tunique ---
	case "Tunique":
		if c.Inventory["Tunique de l’aventurier"] > 0 {
			if c.Body != nil {
				c.addInventory(c.Body.GetName())
			}
			c.Body = CreateEquipment("Tunique de l’aventurier").(*Torso)
			c.removeInventory("Tunique de l’aventurier", 1)
			c.updateMaxHP()
			TypeWriter("✅ Vous avez équipé la Tunique de l’aventurier 👕", 15*time.Millisecond)
		} else {
			TypeWriter("⚠️ Vous n’avez pas de Tunique de l’aventurier.", 15*time.Millisecond)
		}

	// --- Bottes ---
	case "Bottes":
		if c.Inventory["Bottes de l’aventurier"] > 0 {
			if c.Feet != nil {
				c.addInventory(c.Feet.GetName())
			}
			c.Feet = CreateEquipment("Bottes de l’aventurier").(*Foot)
			c.removeInventory("Bottes de l’aventurier", 1)
			c.updateMaxHP()
			TypeWriter("✅ Vous avez équipé les Bottes de l’aventurier 👢", 15*time.Millisecond)
		} else {
			TypeWriter("⚠️ Vous n’avez pas de Bottes de l’aventurier.", 15*time.Millisecond)
		}

	default:
		TypeWriter("❌ Choix invalide, réessayez.", 15*time.Millisecond)
	}
}

////////////////////////////////////////////////////////////////////////////////
// 🏭 FACTORY (Création d’équipements)
////////////////////////////////////////////////////////////////////////////////

// Retourne un objet Equipment correspondant au nom donné
func CreateEquipment(name string) Equipment {
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
