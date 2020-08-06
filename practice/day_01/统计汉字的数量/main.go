package main

import (
	"fmt"
	"strings"
)

// 统计字符串中汉字出现的数量

func demo() {
	str := `BilibiliWorld（以下简称BW）是由中国深受年轻人喜爱的bilibili弹幕视频网站举办的大型线下活动，于2017年首次举办。
从最初的B站PUGC内容展示为主到虚拟现实结合的互动体验，2020年的BW还将会举办新产品、新游戏、新内容等B站相关资讯的发布会。
七彩虹旗下高端子品牌iGame此次也将登陆2020BW，与广大B站粉丝和专业用户玩在一起`
	count := strings.ContainsRune(str, 32)
	fmt.Print(count)
}
func main() {
	demo()
}
