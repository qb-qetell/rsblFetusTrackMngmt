package rsblFetusTrackMngmt

import (
	"fmt"
	"github.com/qeetell/songTrack"
	"encoding/json"
	"testing"
)

func f1 (clap <-chan []string, flap chan<- []string) {
	_x1100 := <- clap
	fmt.Println ("f1:", _x1100)
	_x1200 := make (map[string]string)
	_x1200 ["startupSccsssStatusId"] = "1"
	_x1200 ["startupSccsssStatusDscrpt"] = "Hi"
	_x1300, _ := json.Marshal (_x1200)
	flap <- []string {"", "", string (_x1300)}

	for i := 1; i <= 200000; i ++ {
		flap <- []string {_x1100 [2], "b", "hi"}
		flap <- []string {_x1100 [2], "c", "hi"}
		flap <- []string {_x1100 [2], " ", "hi"}
	}
}
func f2 (clap <-chan []string, flap chan<- []string) {
	_x1100 := <- clap
	fmt.Println ("f2:", _x1100)
	_x1200 := make (map[string]string)
	_x1200 ["startupSccsssStatusId"] = "1"
	_x1200 ["startupSccsssStatusDscrpt"] = "Hi"
	_x1300, _ := json.Marshal (_x1200)
	flap <- []string {"", "", string (_x1300)}

	for {
		_x1400 := <- clap
		fmt.Println ("f2:", _x1400)
	}
}
func f3 (clap <-chan []string, flap chan<- []string) {
	_x1100 := <- clap
	fmt.Println ("f3:", _x1100)
	_x1200 := make (map[string]string)
	_x1200 ["startupSccsssStatusId"] = "1"
	_x1200 ["startupSccsssStatusDscrpt"] = "Hi"
	_x1300, _ := json.Marshal (_x1200)
	flap <- []string {"", "", string (_x1300)}

	for {
		_x1400 := <- clap
		fmt.Println ("f3:", _x1400)
	}
}

func TestPremier (t *testing.T) {
	tracks := [][3]interface {} {
		[3]interface {} {"a", songTrack.Track_Create (f1), "a"},
		[3]interface {} {"b", songTrack.Track_Create (f2), "b"},
		[3]interface {} {"c", songTrack.Track_Create (f3), "c"},
	}
	_x5100, _x6100 := ManageFetusTrack (tracks, nil)
	if _x5100 == true {
		return
	}
	ManageFetusTrack (nil, _x6100)
}
