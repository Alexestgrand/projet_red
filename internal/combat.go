package internal

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/color"
)

var goblinArt = `
       ,      ,
      /(.-""-.)\
      |\  \/  /|
      | \ / =/ |
      \(.\_v_/.)/
       \   |   /
        \  |  /
         )     (
        /       \
       (         )
        \       /
         \__ __/
          // \\
         ||   ||
`

var raykazaArt = `
                 /^\/^\
               _|__|  O|
      \/     /~     \_/ \
       \____|__________/  \
              \_______      \
                      \     \                 \
                       |     |                  \
                      /      /                    \
                     /     /                       \
                   /      /                         \ \
                  /     /                            \  \
                /     /             _----_            \   \
               /     /           _-~      ~-_         |   |
              (      (        _-~    _--_    ~-_     _/   |
               \      ~-____-~    _-~    ~-_    ~-_-~    /
                 ~-_           _-~          ~-_       _-~
                    ~--______-~                ~-___-~
`

//////////////////////////////////////////////////////
// 📦 Helpers for Bubble Tea menus
//////////////////////////////////////////////////////

type menuItem struct{ title string }

func (i menuItem) Title() string       { return i.title }
func (i menuItem) Description() string { return "" }
func (i menuItem) FilterValue() string { return i.title }

type menuModel struct {
	list   list.Model
	choice string
	quit   bool
}

func (m menuModel) Init() tea.Cmd { return nil }

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if selected, ok := m.list.SelectedItem().(menuItem); ok {
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

func (m menuModel) View() string {
	return m.list.View()
}

//////////////////////////////////////////////////////
// 🎮 Character Turn
//////////////////////////////////////////////////////

func (c *Character) characterTurn(monster *Monster) {
	TypeWriter("\n🌟 --- Tour du joueur --- 🌟", 20*time.Millisecond)

	// Build menu
	items := []list.Item{
		menuItem{"Attaquer"},
		menuItem{"Inventaire"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 20, 6)
	l.Title = color.Bold.Sprintf("👉 Choisis une action")

	m := menuModel{list: l}
	p := tea.NewProgram(m)
	finalModel, _ := p.Run()
	m = finalModel.(menuModel)

	switch m.choice {
	case "Attaquer":
		c.chooseAttack(monster)
	case "Inventaire":
		c.UseItem(monster)
	default:
		TypeWriter("⚠️ Choix invalide, votre tour est perdu !", 20*time.Millisecond)
	}
}

//////////////////////////////////////////////////////
// 🔥 Choose Attack (skills menu)
//////////////////////////////////////////////////////

func (c *Character) chooseAttack(monster *Monster) {
	TypeWriter(fmt.Sprintf("\n💥 Mana : %d/%d", c.Mana, c.MaxMana), 10*time.Millisecond)

	// Build skill menu
	items := []list.Item{}
	for _, spell := range c.Skills {
		items = append(items, menuItem{spell})
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 8)
	l.Title = color.Bold.Sprintf("📖 Attaques & Sorts")

	m := menuModel{list: l}
	p := tea.NewProgram(m)
	finalModel, _ := p.Run()
	m = finalModel.(menuModel)

	if m.choice != "" {
		c.CastSpell(m.choice, monster)
	} else {
		TypeWriter("⚠️ Aucun sort sélectionné.", 15*time.Millisecond)
	}
}

//////////////////////////////////////////////////////
// ⚔️ Cast Spell
//////////////////////////////////////////////////////

func (c *Character) CastSpell(spell string, target *Monster) {
	switch spell {
	case "Coup de poing":
		baseDamage := 8
		damage := baseDamage + int(math.Pow(float64(c.AtkPoints), 1.2))
		target.CurrentHP -= damage
		if target.CurrentHP < 0 {
			target.CurrentHP = 0
		}
		c.Mana += 2
		if c.Mana > c.MaxMana {
			c.Mana = c.MaxMana
		}
		TypeWriter(fmt.Sprintf("👊 %s inflige %d dégâts à %s (+2 mana)", c.Name, damage, target.Name), 15*time.Millisecond)

	case "Boule de feu":
		if c.Mana < 10 {
			TypeWriter("⚠️ Pas assez de mana !", 15*time.Millisecond)
			return
		}
		damage := 18 + int(math.Pow(float64(c.AtkPoints), 1.2))
		c.Mana -= 10
		target.CurrentHP -= damage
		TypeWriter(fmt.Sprintf("🔥 %s lance Boule de feu et inflige %d dégâts à %s !", c.Name, damage, target.Name), 15*time.Millisecond)

	case "Soin léger":
		if c.Mana < 5 {
			TypeWriter("⚠️ Pas assez de mana !", 15*time.Millisecond)
			return
		}
		c.Mana -= 5
		heal := 15
		c.CurrentHP += heal
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		TypeWriter(fmt.Sprintf("✨ %s se soigne de %d PV ! (%d/%d PV)", c.Name, heal, c.CurrentHP, c.MaxHP), 15*time.Millisecond)

	default:
		TypeWriter("⚠️ Sort inconnu.", 15*time.Millisecond)
	}

	if target.CurrentHP < 0 {
		target.CurrentHP = 0
	}
	TypeWriter(fmt.Sprintf("🛡️ %s → PV : %d/%d", target.Name, target.CurrentHP, target.MaxHP), 10*time.Millisecond)
}

//////////////////////////////////////////////////////
// 🎮 Training Fight
//////////////////////////////////////////////////////

// Dispatch monster attack depending on monster type
func monsterAttack(c *Character, m *Monster, tour int) {
	switch m.Name {
	case "Gobelin d’entraînement":
		goblinPattern(c, m, tour)
	case "Raykaza le Dragon Ancien":
		raykazaPattern(c, m, tour)
	default:
		// Basic attack if no special pattern
		damage := m.AtkPoints
		c.CurrentHP -= damage
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}
		TypeWriter(fmt.Sprintf("👹 %s attaque %s et inflige %d dégâts !", m.Name, c.Name, damage), 15*time.Millisecond)
	}
}

func trainingFight(c *Character) {
	// Choose monster
	items := []list.Item{
		menuItem{"Gobelin d’entraînement"},
		menuItem{"Raykaza le Dragon Ancien"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 30, 8)
	l.Title = color.Bold.Sprintf("⚔️ Choisis ton adversaire")

	m := menuModel{list: l}
	p := tea.NewProgram(m)
	finalModel, _ := p.Run()
	m = finalModel.(menuModel)

	if m.choice == "" {
		TypeWriter("⚠️ Aucun adversaire choisi.", 15*time.Millisecond)
		return
	}

	var monster Monster
	var monsterArt string
	var expReward int
	switch m.choice {
	case "Gobelin d’entraînement":
		monster = InitGoblin()
		monsterArt = goblinArt
		expReward = 30
	case "Raykaza le Dragon Ancien":
		monster = InitRaykaza()
		monsterArt = raykazaArt
		expReward = 100
	}

	TypeWriter(fmt.Sprintf("\n⚔️ Début du combat contre %s !", monster.Name), 20*time.Millisecond)
	fmt.Println(monsterArt)

	tour := 1
	playerFirst := c.Initiative >= monster.Initiative

	// Combat loop
	for c.CurrentHP > 0 && monster.CurrentHP > 0 {
		color.Bold.Printf("\n===== TOUR %d =====\n", tour)

		if playerFirst {
			c.characterTurn(&monster)
			if monster.CurrentHP <= 0 {
				break
			}
			monsterAttack(c, &monster, tour)
		} else {
			monsterAttack(c, &monster, tour)
			if c.CurrentHP <= 0 {
				isDead(c)
				return
			}
			c.characterTurn(&monster)
		}
		tour++
	}

	// End
	if c.CurrentHP > 0 {
		TypeWriter(fmt.Sprintf("✅ %s a vaincu %s !", c.Name, monster.Name), 20*time.Millisecond)
		c.GainExp(expReward)
		c.Pokedollar += monster.Reward
		color.Green.Printf("💰 Vous gagnez %d pokedollars ! (Total: %d)\n", monster.Reward, c.Pokedollar)
	} else {
		TypeWriter("💀 Vous avez perdu le combat...", 20*time.Millisecond)
	}
}
