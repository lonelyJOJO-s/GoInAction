package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults {
		results <- result // 此步有疑问
	}
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}


