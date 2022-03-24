//-- p --
package rsblFetusTrackMngmt

//-- r --
import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/qeetell/songTrack"
	"strings"
	"sync"
)

func frwrdMssg (clap <-chan []string, flap chan<- []string, tracks [][3]interface {}) {
	accsss := []*sync.Mutex {}
	for _, _ = range tracks {
		access := &sync.Mutex {}
		accsss = append (accsss, access)
	}
	
	for {
		mssg := <- clap
		for id, track := range tracks {
			if strings.Index (mssg [1], track [2].(string)) == 0 {
				go func (i int, t [3]interface {}) {
					accsss [i].Lock ()
					t [1].(*songTrack.Track).CLAP_Fill (mssg, 0)
					accsss [i].Unlock ()
				} (id, track)
				break
			}
			if (id + 1) == len (tracks) {
				fmt.Println ("dropped")
			}
		}
	}
}

//-- i --
func ManageFetusTrack (tracks, exctnFile [][3]interface {}) (errorEncntrStatus bool,
	presentExctnFile [][3]interface {}) {
	
	if exctnFile == nil {
		for _, track5011 := range tracks {
			go track5011 [1].(*songTrack.Track).Run ()
			track5011 [1].(*songTrack.Track).CLAP_Fill ([]string {"", "",
				track5011 [2].(string)}, 0)
		
			_, _y5111 :=  track5011 [1].(*songTrack.Track).FLAP_Read (0)
			if gjson.Get (_y5111 [2], "startupSccsssStatusId").String () == "0" {
				_y5211 := fmt.Sprintf ("---- E: Task could not be started. " +
					"[Sub-task '%s' could not be started. [%s]]",
					track5011 [0].(string),
					gjson.Get (_y5111 [2],
						"startupSccsssStatusDscrpt").String (),
				)
				fmt.Println (_y5211)
				errorEncntrStatus = true
				presentExctnFile = nil
				return
			}
	
			_y5311 := fmt.Sprintf ("---- S: Sub-task '%s' started up successfully.",
				track5011 [0].(string))
			fmt.Println (_y5311)
		}

		errorEncntrStatus = false
		presentExctnFile = tracks
		return
	} else {
		clap := make (chan []string)
		flap := make (chan []string)
		go frwrdMssg (clap, flap, exctnFile)
		
		chnl := make (chan []string)
		chnlAccess := &sync.Mutex {}
		for _, track5411 := range exctnFile {
			go func (t [3]interface {}) {
				for {
					_, _x5511 := t [1].(*songTrack.Track).FLAP_Read (0)
					chnlAccess.Lock ()
					chnl <- _x5511
					chnlAccess.Unlock ()
				}
			} (track5411)
		}
		
		for {
			_x7100 := <- chnl
			clap <- _x7100
		}
		
		errorEncntrStatus = false
		presentExctnFile = tracks
		return
	}
}
