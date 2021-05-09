package strings

import (
	"fmt"
	"testing"
)

// 在字符串中找到第一个只出现一次的字符
// 这里的求法有很多种，比如map、bitmap或者是位运算，前两种比较简单，下面是网上抄的位运算
func TestFindFirstOnceStr(t *testing.T) {
	s := "aaccbeeddd"
	// byte的范围0-128分开计算
	hn := uint64(0) // 64-128
	ln := uint64(0) // 0-63
	// 存放已经存在的
	ehn := uint64(0)
	eln := uint64(0)
	l := uint64(0)
	// 先循环找出所有只出现一次的字符
	for i := range s {
		if s[i] >= 64 {
			l = 1 << (s[i] - 64) // 没有严格处理128这个字符，128这里会报错，需要额外一个来储存(这里就不不实现了)，该解法只是提供思路
			if l&ehn == 0 {
				ehn |= l
				hn |= l
				continue
			}
			if hn&l != 0 {
				hn ^= l
			}
			continue
		}
		l = 1 << s[i]
		if l&eln == 0 {
			eln |= l
			ln |= l
			continue
		}
		if ln&l != 0 {
			ln ^= l
		}
	}
	ans := byte(' ')
	// 找到最新出现一次的字符
	for i := range s {
		if s[i] >= 64 {
			l = 1 << (s[i] - 64)
			if hn&l != 0 {
				ans = s[i]
				break
			}
			continue
		}
		l = 1 << s[i]
		if ln&l != 0 {
			ans = s[i]
			break
		}
	}

	// test1
	// test2
	// test3
	// test4
	// test/rebase merge
	fmt.Println(string(ans))
	// I want to test rebase
	// main merge
}
