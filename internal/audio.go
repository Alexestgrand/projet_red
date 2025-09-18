package internal

import (
	"fmt"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// internal/sound.go
func PlaySFX(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("⚠️ Impossible de charger le son:", path, "-", err)
		return
	}
	defer f.Close()

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("⚠️ Erreur décodage audio:", err)
		return
	}

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		streamer.Close()
	})))
}
