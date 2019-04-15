package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitter sync.WaitGroup

	waitter.Add(len(feeds))

	for _, feed := range feeds {
		match, exist := matchers[feed.Type]
		if !exist {
			match = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(feed, matcher, searchTerm, results)
			waitter.Done()
		}(match, feed)

		go func() {
			waitter.Wait()
			close(results)
		}()
	}
	Display(results)
}

// Register 注册器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
