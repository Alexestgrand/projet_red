package internal
 
import (
    "fmt"
    "time"
)
 
// ===== MENU PRINCIPAL =====
func Menu(c1 *Character) {
    for {
        var choix int
        fmt.Println("\n=== MENU ===")
        fmt.Println("1. Infos du personnage")
        fmt.Println("2. Inventaire")
        fmt.Println("3. Marchand")
        fmt.Println("4. Quitter")
        fmt.Print("Choix : ")
        fmt.Scanln(&choix)
 
        switch choix {
        case 1:
            c1.DisplayInfo()
 
        case 2:
            fmt.Println("\n--- Inventaire ---")
            c1.AccessInventory()
            fmt.Println("1. Utiliser une potion")
            fmt.Println("2. Retour")
            fmt.Print("Choix : ")
            var invChoice int
            fmt.Scanln(&invChoice)
 
            if invChoice == 1 {
                if len(c1.Inventory) == 0 {
                    fmt.Println("⚠️ Votre inventaire est vide.")
                    continue
                }
 
                // Choix de l’objet à utiliser
                fmt.Println("\nQuel objet voulez-vous utiliser ?")
                keys := make([]string, 0, len(c1.Inventory))
                for item := range c1.Inventory {
                    keys = append(keys, item)
                }
                for i, item := range keys {
                    fmt.Printf("%d. %s\n", i+1, item)
                }
                var itemChoice int
                fmt.Print("Choix : ")
                fmt.Scanln(&itemChoice)
 
                if itemChoice > 0 && itemChoice <= len(keys) {
                    selected := keys[itemChoice-1]
                    switch selected {
                    case "Potion de vie":
                        c1.TakePot()
                    case "Potion de poison":
                        c1.PoisonPot()
                    default:
                        fmt.Println("⚠️ Cet objet ne peut pas être utilisé.")
                    }
                    // Supprime l’objet utilisé
                    c1.removeInventory(selected)
                } else {
                    fmt.Println("⚠️ Choix invalide.")
                }
            }
 
        case 3:
            c1.MarchantInterface()
 
        case 4:
            fmt.Println("Au revoir !")
            return
 
        default:
            fmt.Println("⚠️ Choix invalide !")
        }
    }
}
 
// ===== INVENTAIRE =====
func (c *Character) addInventory(item string) {
    c.Inventory[item]++
}
 
func (c *Character) removeInventory(item string) {
    delete(c.Inventory, item)
}
 
// ===== MARCHAND =====
func (c *Character) MarchantInterface() {
    var choix int
    fmt.Println("\n=== Marchand ===")
    fmt.Println("1. Potion de vie (gratuite)")
    fmt.Println("2. Potion de poison (gratuite)")
    fmt.Println("3. Retour")
    fmt.Print("Choix : ")
    fmt.Scanln(&choix)
 
    switch choix {
    case 1:
        c.addInventory("Potion de vie")
        fmt.Println("Vous avez obtenu : Potion de vie")
    case 2:
        c.addInventory("Potion de poison")
        fmt.Println("Vous avez obtenu : Potion de poison")
    case 3:
        return
    default:
        fmt.Println("⚠️ Choix invalide !")
    }
}
 
// ===== POTIONS =====
func (c *Character) PoisonPot() {
    damage := 10
    duration := 3
 
    fmt.Println("⚠️ Vous utilisez une Potion de poison !")
 
    for i := 1; i <= duration; i++ {
        c.CurrentHP -= damage
        if c.CurrentHP < 0 {
            c.CurrentHP = 0
        }
        fmt.Printf("⏳ Tour %d : -%d HP → %d/%d PV restants\n",
            i, damage, c.CurrentHP, c.MaxHP)
 
        time.Sleep(1 * time.Second)
 
        if c.CurrentHP == 0 {
            fmt.Println("💀 Vous êtes mort à cause du poison !")
            break
        }
    }
}