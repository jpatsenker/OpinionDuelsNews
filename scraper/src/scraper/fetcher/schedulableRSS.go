package fetcher

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"scraper/scheduler"
	"time"
)

// make an RSS schedulable
type SchedulableRSS struct {
	RSSFeed RSS
	delay   int
	start   time.Time
}

func (rss *SchedulableRSS) DoWork(scheduler *scheduler.Scheduler) {
	rss.start = time.Now() // reset the timer at the top of the loop
	resp, err := http.Get(rss.RSSFeed.GetLink())
	fmt.Println("link is:", rss.RSSFeed.GetLink())
	if err != nil {
		// TODO: error checking here
		fmt.Println("error getting RSS:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: error handling
		return
	}

	err = GetStories(rss.RSSFeed, body)
	if err != nil {
		return
	}

	fmt.Println("OK")
	// TODO: use config file to control the timing here
	toSchedule := CreateSchedulableArticle(rss.RSSFeed.GetChannel().GetArticle(0), 1)
	go scheduler.AddSchedulable(toSchedule)
	toSchedule = CreateSchedulableArticle(rss.RSSFeed.GetChannel().GetArticle(1), 2)
	go scheduler.AddSchedulable(toSchedule)
}

func (rss *SchedulableRSS) GetTimeRemaining() int {
	remainingTime := float64(rss.delay) - time.Since(rss.start).Seconds()
	if remainingTime <= 0 {
		return 0
	}
	return int(math.Ceil(remainingTime))
}

func (rss *SchedulableRSS) IsLoopable() bool {
	// TODO: make this true once out of testing
	return false
}

func (rss *SchedulableRSS) SetTimeRemaining(remaining int) {
	rss.delay = remaining
}

// factory to make schedulable rss
func CreateSchedulableRSS(rss RSS, delay int) *SchedulableRSS {
	return &SchedulableRSS{rss, delay, time.Now()}
}

// check that we implemented this properly
var _ scheduler.Schedulable = (*SchedulableRSS)(nil)
