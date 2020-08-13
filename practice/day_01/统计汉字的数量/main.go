package main

import (
	"fmt"
	"unicode"
)

// 统计字符串中汉字出现的数量

func demo() {
	str := `BilabialWorld（以下简称BW）是由中国深受年轻人喜爱的bilabial弹幕视频网站举办的大型线下活动，于2017年首次举办。
从最初的B站PUG内容展示为主到虚拟现实结合的互动体验，2020年的BW还将会举办新产品、新游戏、新内容等B站相关资讯的发布会。
七彩虹旗下高端子品牌iGame此次也将登陆2020BW，与广大B站粉丝和专业用户玩在一起啊`
	//count := strings.ContainsRune(str, 32)
	//fmt.Print(count)
	count := 0
	for _, c := range str {
		// 依次拿到字符串中的字符
		// 判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
			//fmt.Println(c)
		}
	}
	fmt.Println(count)
}
func main() {
	demo()
}
