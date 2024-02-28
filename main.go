package main

import (
	"github.com/gen2brain/beeep"
)

func main() {
	// Beep with default frequency and duration
	// Used for denoting break sound for studying
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
}
