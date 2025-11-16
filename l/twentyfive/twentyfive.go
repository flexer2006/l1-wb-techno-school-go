package twentyfive

import "time"

func Sleep(d time.Duration) {
	if d <= 0 {
		return
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	<-timer.C
}
