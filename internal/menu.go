package internal
 
import "fmt"
 
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
                c1.TakePot()
            } else if invChoice == 2 {
 
            } else {
                fmt.Println("Choix invalide.")
            }
 
        case 3:
            c1.MarchantInterface()
 
        case 4:
            fmt.Println("Au revoir !")
            return // on quitte la boucle et donc le programme
 
        default:
            fmt.Println("Choix invalide !")
        }
    }
}
 
// Ajoute un item à l'inventaire
func (c *Character) addInventory(item string) {
    c.Inventory[item] = 1
}
 
// Retire un item de l'inventaire (par index)
func (c *Character) removeInventory(index int) {
    if index < 0 || index >= len(c.Inventory) {
        return
    }
    for key := range c.Inventory {
        if index == 0 {
            delete(c.Inventory, key)
            break
        }
        index--
    }
}
 
func (c *Character) MarchantInterface() {
    for {
        fmt.Println("\n=== Marchand ===")
        fmt.Println("1. Potion de vie (gratuite)")
        fmt.Println("2. Retour")
        fmt.Print("Choix : ")
 
        var choix int
        fmt.Scanln(&choix)
 
        switch choix {
        case 1:
            c.addInventory("Potion de vie")
            fmt.Println("Vous avez obtenu : Potion de vie")
        case 2:
            return
        default:
            fmt.Println("Choix invalide !")
        }
    }
}
 