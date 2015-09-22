package fetcher

import (
	"scraper/scheduler"
	"testing"
	"time"
)

func TestSchedulableRSS(t *testing.T) {
	t.Skip("skipping RSS test for now")
	s := scheduler.MakeScheduler(5, 3)
	s.Start()

	rss := CreateSchedulableRSS(&WSJRSS{}, 0)
	s.AddSchedulable(rss)
	time.Sleep(time.Duration(6) * time.Second)
	s.Stop()
}
