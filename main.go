package main

import (
	"fmt"
	"os"
	"pinyin"
)

func main() {

	chineseWords := ""
	for i, v := range os.Args {
		if i != 0 {
			chineseWords += v
		}
		if i > 0 && i <= (len(os.Args)-1) {
			chineseWords += " "
		}
	}
	if len(chineseWords) == 0 {
		chineseWords = "hi,我是_中国人"
		fmt.Println("没有参数，使用例子`hi,我是_中国人`")
	}

	str, err := pinyin.New(chineseWords).Split("").Mode(pinyin.ToneUseNumeric).Convert()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(str) //输出 hi,wo3shi4_zhong1guo2ren2
	}

}
