package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from	*string
	help	*string
	//last	*string
	section	*string
	num		*int
)

func init()  {
	from = flag.String("from","","目录地址")
	help = flag.String("help","","帮助：键盘向上(⬆)上一行,键盘向上(⬇)下一行,键盘向左(⬅)上一章,键盘向右(➡)下一章,目前只支持(纵横中文网)")
	//last = flag.String("last","last","继续上次阅读")
	section = flag.String("section","","章节地址")
	num = flag.Int("num",0,"每行显示时间,自动下一行")
	flag.Parse()
	checkFrom()
}

func checkFrom()  {
	if *from == "" {
		fmt.Println("请设置 -from 参数")
		os.Exit(0)
	}
}

