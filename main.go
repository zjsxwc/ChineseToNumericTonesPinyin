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

	fmt.Println("把汉字转换为数字标记声调的拼音，比如 ā 为 a1、á 为 a2、ǎ 为 a3、à 为 a4，只能处理非多音字，多音字需要手动修改。")
	fmt.Println("比如 `hi,我是 中国人` 的结果是 `hi,wo3shi4 zhong1guo2ren2`。 ")

	if len(chineseWords) == 0 {
		fmt.Println("没有参数，请输入，例子 `./ChineseToNumericPinyin hi,我是 中国人`")
	}

	str, err := pinyin.New(chineseWords).Split("").Mode(pinyin.ToneUseNumeric).Convert()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print("`"+ chineseWords +"` 的转换结果是： ")
		fmt.Println(str)
	}

}
