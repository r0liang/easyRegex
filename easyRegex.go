package easyRegex

import (
	"io"
	"regexp"
	"sync"
)

var (
	regexMute = sync.RWMutex{}
	regexMap  = make(map[string]*regexp.Regexp)
)

func getRegexpObject(s string) (rex *regexp.Regexp, err error) {
	// 核心功能
	// 1.自动将正则对象缓存在内存中的map里
	// 2.确保并发安全
	// 3.未来版本中会实现类似于redis的过期、hash等高级功能
	// Core functions
	// 1. Automatically cache regular objects in a map in memory
	// 2. Ensure concurrency security
	// 3. In future versions, advanced functions similar to redis expiration and hashing will be implemented
	regexMute.RLock()
	rex = regexMap[s]
	regexMute.RUnlock()
	if rex == nil {
		return addRegexObject(s)
	} else {
		return
	}
}

func addRegexObject(s string) (rex *regexp.Regexp, err error) {
	// 将正则对象编译后缓存在内存中
	// The regular object is compiled and cached in memory
	rex, err = regexp.Compile(s)
	regexMute.Lock()
	regexMap[s] = rex
	regexMute.Unlock()
	return
}

func Match(pattern string, b []byte) (bool, error) {
	// Match报告字节切片b是否包含正则表达式模式的任何匹配。
	// 更复杂的查询需要使用Compile和完整的Regexp接口。
	// Match reports whether the byte slice b contains any match of the regular expression pattern.
	// More complicated queries need to use Compile and the full Regexp interface.
	obj, err := getRegexpObject(pattern)
	if err != nil {
		return false, err
	}
	return obj.Match(b), err
}

func MatchReader(pattern string, r io.RuneReader) (bool, error) {
	// MatchReader报告RuneReader返回的文本是否包含正则表达式模式的任何匹配项
	// 更复杂的查询需要使用Compile和完整的Regexp接口。
	// MatchReader reports whether the text returned by the RuneReader contains any match of the regular expression pattern.
	// More complicated queries need to use Compile and the full Regexp interface.
	obj, err := getRegexpObject(pattern)
	if err != nil {
		return false, err
	}
	return obj.MatchReader(r), err
}

func MatchString(pattern string, s string) (bool, error) {
	// MatchString报告字符串s是否包含任何正则表达式模式的匹配项。
	// 更复杂的查询需要使用Compile和完整的Regexp接口。
	// MatchString reports whether the string s contains any match of the regular expression pattern.
	// More complicated queries need to use Compile and the full Regexp interface.
	obj, err := getRegexpObject(pattern)
	if err != nil {
		return false, err
	}
	return obj.MatchString(s), err
}
