# Red Adventure 🎮

Un jeu d'aventure RPG en ligne de commande écrit en Go.

## Description

Red Adventure est un jeu de rôle textuel où vous créez votre personnage et partez à l'aventure. Affrontez des monstres, gérez votre inventaire, achetez des équipements chez le marchand, et améliorez vos armes chez le forgeron.

## Fonctionnalités

- 🧙 Création de personnage
- ⚔️ Système de combat
- 🎒 Gestion d'inventaire
- 🏪 Marchand pour acheter des objets
- 🔨 Forgeron pour améliorer vos équipements
- 🎵 Musique de fond et effets sonores
- ✨ Système de sorts

## Prérequis

- Go 1.25.1 ou supérieur

## Installation

```bash
# Cloner le dépôt
git clone <url-du-repo>
cd projet-red_adventure

# Installer les dépendances
go mod download
```

## Lancement du jeu

```bash
# Depuis la racine du projet
cd cmd
go run main.go
```

## Technologies utilisées

- [Go](https://golang.org/) - Langage de programmation
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework TUI
- [Beep](https://github.com/faiface/beep) - Bibliothèque audio
- [Gookit Color](https://github.com/gookit/color) - Couleurs dans le terminal

## Structure du projet

```
.
├── cmd/              # Point d'entrée de l'application
├── internal/         # Logique du jeu
├── assets/           # Fichiers audio
└── utils/            # Utilitaires
```

## Licence

Tous droits réservés.

