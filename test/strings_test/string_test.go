package strings_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	//分割
	list := strings.Split("a,b,c", ",")
	fmt.Println(fmt.Sprintf("%#v", list))
	//合并
	str := strings.Join(list, "-")
	fmt.Println(str)
}

func TestContains(t *testing.T) {
	fmt.Println(strings.Contains("abcdefgef", "ef"))
	fmt.Println(strings.Count("abcdefgef", "ef"))
	fmt.Println(strings.Index("ratretei", "re"))
}

func TestPrefix(t *testing.T) {
	fmt.Println(strings.HasPrefix("abc-xyz", "abc"))
	fmt.Println(strings.HasSuffix("abc-xyz", "xyz"))
	fmt.Println(strings.HasSuffix("abc-xyz", "xyzg"))
}

func TestTrim(t *testing.T) {
	fmt.Println(strings.Trim("1 abc 2 ", "1 "))
	fmt.Println(strings.TrimSpace(" abc "))
	fmt.Println(strings.TrimPrefix("abcefgab", "ab"))
	fmt.Println(strings.TrimLeft("abcefgab", "ab"))
	fmt.Println(strings.TrimSuffix("abcefgab", "ab"))
	fmt.Println(strings.TrimRight("abcefgab", "ab"))
}

func TestReplace(t *testing.T) {
	fmt.Println(strings.Replace("tomo", "o", "i", 1))
	fmt.Println(strings.ReplaceAll("tomo", "o", "i"))
}

func TestUp(t *testing.T) {
	fmt.Println(strings.ToLower("ABC"))
	fmt.Println(strings.ToUpper("xyz"))
	fmt.Println(strings.ToLower("ABCtt"))
	fmt.Println(strings.ToUpper("xyzGG"))
}
