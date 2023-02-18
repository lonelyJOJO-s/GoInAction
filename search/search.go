package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher) // matchers 作为search包的包级变量， 但注意大写开头的变量维公开标识符，小写的为不公开

func Run(searchTerm string) {

	feeds, err := RetrieveFeeds() // 多个网站下去找，feeds包含多个网站信息

	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result) // 一个通道类型

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {

		matcher, exists := matchers[feed.Type] // 一开始便会注册到matchers

		if !exists {
			matcher = matchers["default"] // 得到使用的matcher对象实例
		}

		// 启动一个goruntine来执行搜索结果
		go func(matcher Matcher, feed *Feed) { // 此处为一个匿名函数
			Match(matcher, feed, searchTerm, results) // 闭包原理，并非有副本，而是访问外层函数作用域
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		// 等待所有工作的完成
		waitGroup.Wait()

		close(results) //关闭通道

	}()

	Display(results) // 展示通道内容

}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Macther already registered")
	}

	log.Println("register", feedType, "matcher")
	matchers[feedType] = matcher

}
