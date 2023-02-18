package search

type defaultMatcher struct {
}

// 注册默认匹配器
func init() {
	var matcher defaultMatcher
	Register("default", matcher)

}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error){
	return nil, nil
}

// 注意， 不管传入的接受者为指针还是引用，均可正常调用其实现的函数
// func (m defaultMatcher) Search(feed *Feed, searchTerm string)

// dm := new(defaultMatcher)

// dm.search(feed, "test")

// func (m *defaultMatcher) Search(feed *Feed, searchTerm string)

// var dm defaultMatcher

// dm.search(feed, "test")

