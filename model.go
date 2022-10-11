package main

//定义需要用到的常量
const (
	StudentNum = 90 //一节课90个学生
	ClassNum   = 20 //一共上20节课
	LessonNum  = 5  //一共无门课
)

//定义需要使用到的结构体

// Lesson 课程结构体,五门课程，每个课程一次90人，20次
type Lesson struct {
	Id       int        //课程id
	Students []*Student //学生集合
}

// Student 学生结构体
type Student struct {
	Id     int     //学生id
	Weight float64 //权重，权重越高被点到的概率越大
	IsRun  []bool  //记录每次是否逃课，长度为20，true表示逃了，false表示没逃
}

// Students 实现sort接口，支持排序，按照权重值从大到小排
type Students []*Student

func (s Students) Len() int {
	return len(s)
}

// Less 按照权重值比较，从大到小
func (s Students) Less(i, j int) bool {
	return s[i].Weight >= s[j].Weight
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
