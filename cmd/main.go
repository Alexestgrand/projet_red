package main

import (
	"log"
	"os"
	"projet-red_adventure/internal"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var mixer = &beep.Mixer{} // Mixer global

func main() {
	// --- Charger la musique de fond ---
	f, err := os.Open("../assets/music.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	// Pas de defer streamer.Close(), sinon la musique s'arrête

	// Init speaker une seule fois
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Ajouter la musique de fond en boucle
	loop := beep.Loop(-1, streamer) // -1 = infini
	mixer.Add(loop)

	// Jouer le mixer
	speaker.Play(mixer)

	// --- Lancer le jeu ---
	c := internal.CharacterCreation()
	internal.Menu(&c)
}
