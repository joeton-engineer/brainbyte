// Incorporating random 10-second pauses during learning activities may
// accelerate neural repetitions, akin to deep sleep processes, potentially
// enhancing learning efficiency.  Every 2 minutes having range of 2 to 3 mins.

package pause

import (
	"fmt"
	"math/rand"
	"time"
)

func StartPause() {
	fmt.Println("TESTING....")
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the duration for the learning session
	learningSessionDuration := 10 * time.Minute

	// Define the intervals for incorporating pauses
	pauseInterval := 2 * time.Minute
	pauseRange := time.Duration(rand.Intn(61)+120) * time.Second // Random range between 2 to 3 minutes

	// Initialize the timer for the learning session
	learningSessionTimer := time.NewTimer(learningSessionDuration)

	fmt.Println("Learning session started...")

	// Loop for the duration of the learning session
	for {
		select {
		case <-learningSessionTimer.C:
			// Learning session completed
			fmt.Println("Learning session completed.")
			return
		default:
			// Conduct learning activities
			fmt.Println("Conducting learning activities...")

			// Check if it's time for a pause
			if time.Since(learningSessionTimer.Reset(pauseInterval)) >= pauseRange {
				// Pause for random duration
				pauseDuration := time.Duration(rand.Intn(11)) * time.Second // Random 10-second pause
				fmt.Printf("Taking a %s pause...\n", pauseDuration)
				time.Sleep(pauseDuration)
			}
		}
	}
}
