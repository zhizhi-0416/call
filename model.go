package main

//定义需要使用到的结构体

type Manager struct {
	AllCount int //点名的所有次数
	HitCount int //点名的命中次数
}

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
