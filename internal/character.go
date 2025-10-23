package internal

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/color"
)

// //////////////////////////////////////////////////////
// 🧙 CHARACTER STRUCTURE
// //////////////////////////////////////////////////////
type Character struct {
	Name         string
	Class        string
	Level        int
	MaxHP        int
	CurrentHP    int
	Inventory    map[string]int
	Skills       []string
	Pokedollar   int
	Initiative   int
	MaxInventory int
	UpgradeUsed  int
	AtkPoints    int

	// Experience
	Exp      int
	ExpToLvl int

	// Mana
	Mana    int
	MaxMana int

	// Equipment
	Head *Headset
	Body *Torso
	Feet *Foot
}

// //////////////////////////////////////////////////////
// ✍️ TYPEWRITER EFFECT
// //////////////////////////////////////////////////////
func TypeWriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

// //////////////////////////////////////////////////////
// ⚙️ CHARACTER INITIALIZATION
// //////////////////////////////////////////////////////
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
		Mana:         30,
		MaxMana:      30,
		Exp:          0,
		ExpToLvl:     50,
		AtkPoints:    5,
	}
}

// //////////////////////////////////////////////////////
// 📜 CLASS SELECTION (Bubble Tea)
// //////////////////////////////////////////////////////

// Class item for Bubble Tea
type classItem struct{ title string }

func (i classItem) Title() string       { return i.title }
func (i classItem) Description() string { return "" }
func (i classItem) FilterValue() string { return i.title }

// List of available classes
var classChoices = []classItem{
	{title: "Humain (100 PV, Atk 5)"},
	{title: "Elfe (80 PV, Atk 7)"},
	{title: "Nain (120 PV, Atk 4)"},
}

// Bubble Tea model for class selection
type model struct {
	list   list.Model
	choice string
	quit   bool
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if selected, ok := m.list.SelectedItem().(classItem); ok {
				m.choice = selected.title
				m.quit = true
				return m, tea.Quit
			}
		case "q", "esc":
			m.quit = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quit && m.choice != "" {
		return fmt.Sprintf("\n✅ Classe choisie : %s\n", m.choice)
	}
	return "\nChoisis ta classe :\n" + m.list.View()
}

// //////////////////////////////////////////////////////
// ✨ CHARACTER CREATION FLOW
// //////////////////////////////////////////////////////
func CharacterCreation() Character {
	// ASCII banner
	banner := color.Style{color.FgGreen, color.OpBold}.Sprintf(`
██████╗  ██████╗ ██╗  ██╗███████╗██╗    ██╗ ██████╗ ██████╗ ██╗     ██████╗ 
██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝██║    ██║██╔═══██╗██╔══██╗██║     ██╔══██╗
██████╔╝██║   ██║█████╔╝ █████╗  ██║ █╗ ██║██║   ██║██████╔╝██║     ██║  ██║
██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║███╗██║██║   ██║██╔══██╗██║     ██║  ██║
██║     ╚██████╔╝██║  ██╗███████╗╚███╔███╔╝╚██████╔╝██║  ██║███████╗ █████╔╝
╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝ ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝ ╚════╝ 
`)
	fmt.Println(banner)

	// Ask for name
	var name string
	for {
		TypeWriter(color.Style{color.FgGreen, color.OpBold}.Sprintf("Quel est le nom du jeune Aventurier ?"), 20*time.Millisecond)
		fmt.Scanln(&name)

		// Vérification : uniquement des lettres
		valid := true
		for _, letter := range name {
			if !unicode.IsLetter(letter) {
				valid = false
				break
			}
		}

		if !valid || len(name) == 0 {
			TypeWriter("⚠️ Que des lettres uniquement !", 15*time.Millisecond)
			continue
		}

		// Normalisation → 1ère majuscule + reste en minuscule
		name = strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])

		break
	}

	// Build Bubble Tea list of classes
	items := make([]list.Item, len(classChoices))
	for i, c := range classChoices {
		items[i] = c
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 10)
	l.Title = color.Bold.Sprintf("Choisis ta classe")

	// Run Bubble Tea program
	m := model{list: l}
	p := tea.NewProgram(m)
	finalModel, _ := p.Run()
	m = finalModel.(model)

	class := m.choice
	maxHP, atk := 100, 5
	switch class {
	case "Humain (100 PV, Atk 5)":
		maxHP, atk = 100, 5
	case "Elfe (80 PV, Atk 7)":
		maxHP, atk = 80, 7
	case "Nain (120 PV, Atk 4)":
		maxHP, atk = 120, 4
	}

	// Create character
	currentHP := maxHP / 2
	inventory := map[string]int{}
	character := InitCharacter(name, class, 1, maxHP, currentHP, inventory)
	character.AtkPoints = atk

	TypeWriter("\n🎉 Félicitations ! Ton personnage a bien été créé :", 15*time.Millisecond)
	character.DisplayInfo()
	return character
}

// //////////////////////////////////////////////////////
// ⭐ EXPERIENCE & LEVEL-UP
// //////////////////////////////////////////////////////
func (c *Character) GainExp(amount int) {
	TypeWriter(fmt.Sprintf("⭐ %s gagne %d XP!", c.Name, amount), 15*time.Millisecond)
	c.Exp += amount

	for c.Exp >= c.ExpToLvl {
		c.Level++
		c.Exp -= c.ExpToLvl
		c.ExpToLvl += 25
		c.MaxHP += 10
		c.MaxMana += 5
		c.AtkPoints += 2
		c.CurrentHP = c.MaxHP
		c.Mana = c.MaxMana

		TypeWriter(fmt.Sprintf("🎉 %s passe au niveau %d ! ❤️ PV = %d | 💧 Mana = %d | 🗡️ Atk = %d",
			c.Name, c.Level, c.MaxHP, c.MaxMana, c.AtkPoints), 20*time.Millisecond)
		PlaySFX("assets/levelup.mp3")
	}
	c.displayExpBar()
}

// //////////////////////////////////////////////////////
// 📊 XP BAR
// //////////////////////////////////////////////////////
func (c *Character) displayExpBar() {
	totalBars := 20
	filledBars := (c.Exp * totalBars) / c.ExpToLvl
	bar := ""
	for i := 0; i < totalBars; i++ {
		if i < filledBars {
			bar += color.Green.Sprintf("█")
		} else {
			bar += color.Gray.Sprintf("░")
		}
	}
	fmt.Printf("🔸 XP [%s] %d/%d\n", bar, c.Exp, c.ExpToLvl)
}

// //////////////////////////////////////////////////////
// 📋 CHARACTER INFO
// //////////////////////////////////////////////////////
func (c Character) DisplayInfo() {
	fmt.Println("\n========== CHARACTER INFO ==========")
	fmt.Printf("🧑 Nom    : %s\n", color.Cyan.Sprintf("%s", c.Name))
	fmt.Printf("🏹 Classe : %s\n", color.Yellow.Sprintf("%s", c.Class))
	fmt.Printf("⭐ Niveau : %d\n", c.Level)
	fmt.Printf("❤️ PV     : %s/%s\n", color.Red.Sprintf("%d", c.CurrentHP), color.Red.Sprintf("%d", c.MaxHP))
	fmt.Printf("💧 Mana   : %s/%s\n", color.Blue.Sprintf("%d", c.Mana), color.Blue.Sprintf("%d", c.MaxMana))
	fmt.Printf("🗡️ Atk    : %d\n", c.AtkPoints)
	fmt.Printf("💰 Argent : %d pokedollars\n", c.Pokedollar)

	// Inventory
	fmt.Printf("🎒 Inventaire: %v\n", c.Inventory)

	// Skills
	fmt.Print("🔥 Compétences : ")
	if len(c.Skills) == 0 {
		fmt.Println("Aucune")
	} else {
		for _, skill := range c.Skills {
			fmt.Printf("%s ", color.Magenta.Sprintf("%s", skill))
		}
		fmt.Println()
	}
	c.displayExpBar()
	fmt.Println("====================================")
}

// //////////////////////////////////////////////////////
// 💀 DEATH & RESURRECTION
// //////////////////////////////////////////////////////
func isDead(c *Character) {
	if c.CurrentHP <= 0 {
		TypeWriter("\n💀 Le joueur est mort...", 30*time.Millisecond)
		c.CurrentHP = c.MaxHP / 2
		c.Mana = c.MaxMana / 2
		TypeWriter(fmt.Sprintf("✨ Ressuscité avec %d PV et %d mana.", c.CurrentHP, c.Mana), 30*time.Millisecond)
		PlaySFX("assets/death.mp3")
	}
}
