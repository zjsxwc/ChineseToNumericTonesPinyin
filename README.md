
## 把汉字转换为数字标记声调的拼音

#### 编译

```bash
rm ChineseToNumericPinyin
make
```

#### 用法
./ChineseToNumericPinyin 飞龙_在天
把汉字转换为数字标记声调的拼音，比如 ā 为 a1、á 为 a2、ǎ 为 a3、à 为 a4，只能处理非多音字，多音字需要手动修改。
比如 `hi,我是 中国人` 的结果是 `hi,wo3shi4 zhong1guo2ren2`。 
`飞龙_在天 ` 的转换结果是： fei1long2_zai4tian1 