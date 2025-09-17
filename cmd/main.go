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

func main() {
	// Ouvre le fichier audio
	f, err := os.Open("../internal/music.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Décode le fichier
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialise le haut-parleur
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Lance la musique en parallèle
	go func() {
		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done
	}()

	// 🚀 Le jeu démarre directement
	c := internal.CharacterCreation()
	internal.Menu(&c)
}
