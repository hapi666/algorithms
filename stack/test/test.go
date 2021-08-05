package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	str := now.Format("today is Monday, On January 2，现在为您报时：当前时间是15时4分5秒")
	fmt.Println(str)
}
