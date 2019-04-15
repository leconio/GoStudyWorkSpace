package search

import "log"

// Result 匹配结果
type Result struct {
	Field   string
	Content string
}

// Matcher 匹配器接口
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 匹配器
func Match(feed *Feed, matcher Matcher, searchTerm string, results chan<- *Result) {
	result, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range result {
		results <- content
	}
}

func Display(results chan *Result) {

	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
