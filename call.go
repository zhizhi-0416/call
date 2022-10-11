package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

//提供主要的点名逻辑

var (
	AllCount   = 0    //点名的所有次数
	HitCount   = 0    //点名的命中次数
	FirstCount = 45   //第一次点的人数
	MinCount   = 5    //至少点到人数
	Threshold  = 30.0 //阈值，超过这个阈值就一定点名
)

const (
	Rate        float64 = 1.7 //每次下降的比率
	PlusCount   float64 = 15  //每次提高权重
	ReduceCount float64 = -6  //每次减少的权重
)

// Call 一次课程点名的逻辑
// 初始时所有的学生的权重都为0，都可能被点名
// 为了快速筛选，第一次点45名学生，之后以一定比率下降，下降至每次八人
// 每一次被点到并且逃课了权重值+16，否则权重值-6，每次点名优先点权重最高的学生，需要进行排序
// 由于80%概率逃课，对于一个80%逃课的学生来说 连续到三次的概率是非常低的，
// 即三次点错抵消增加的权重值，可以快速筛选出来经常逃课的学生
// 暂时将下降比率设置为1.7，在下降到最低人次8人时能大致覆盖所有的学生
func Call(lesson *Lesson) {
	count := FirstCount //初始化点到人数，可以进行参数调整

	//第i表示第i轮点名,先打乱数组，再进行排序
	//排序后点到前count个，如果大于阈值的人数大于count，就更新count
	for i := 0; i < ClassNum; i++ {
		aboveThreshold := 0
		randShuffle(lesson.Students) //打乱数组，随机选择
		students := lesson.Students
		sort.Sort(Students(students)) //进行排序操作，按照权重值从大到小排
		for j := 0; j < len(students); j++ {
			if students[j].Weight > Threshold {
				aboveThreshold++
			}
			if aboveThreshold > count { //如果超过阈值的人数大于当前挑选人数那么就进行改变
				count = aboveThreshold
			}
		}
		for j := 0; j < count; j++ {
			student := students[j]
			student.Weight += hit(i, student) //增加或者减小权重
		}
		count = int(math.Floor(float64(count) / Rate)) //计算下一次轮次人数
		if count < MinCount {
			count = MinCount //保证不低于最低点到人数
		}

		if aboveThreshold > count { //如果超过阈值的人数大于当前挑选人数那么就进行改变
			count = aboveThreshold
		}
	}

}

// hit 表示点名是否命中
func hit(round int, student *Student) float64 {
	AllCount++ //增加点到次数
	if student.IsRun[round] == true {
		HitCount++ //增加命中次数
		return PlusCount
	}
	return ReduceCount
}

// GetResult 点五次名字，拿到所有数据
func GetResult() float64 {
	return float64(HitCount) / float64(AllCount)
}

// 随机打乱数组
func randShuffle(slice []*Student) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
