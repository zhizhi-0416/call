package main

import (
	"fmt"
	"time"
)

func main() {

	GenerateAllLessonFile()
	for i := 1; i <= LessonNum; i++ {
		lesson, err := ReadFile(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		Call(lesson)
		everyResult := GetResult()
		fmt.Printf("第 %d 类课获取的 E 为 %f \n", i, everyResult)
	}
	result := GetResult()
	fmt.Printf("总计 点名 %d 次，命中 %d次 最终 E 为 %f \n", AllCount, HitCount, result)

	time.Sleep(time.Minute)
}
