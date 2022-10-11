package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GenerateData 生成第i门课的数据
func GenerateData(lessonId int) Lesson {

	rand.Seed(time.Now().Unix()) //设置随机种子

	//创建课程
	lesson := Lesson{
		Id:       lessonId,
		Students: nil,
	}

	students := make([]*Student, StudentNum)
	for i := 0; i < StudentNum; i++ {
		students[i] = &Student{
			Id:     i,
			Weight: 0,                      //权重初始化为0
			IsRun:  make([]bool, ClassNum), //20次记录，默认为false即不逃课
		}
	}

	lesson.Students = students //初始化学生列表

	alwaysRun := rand.Intn(3) + 5 //选出5-8个不同的人，随机人数5~8
	m := make(map[int]struct{})   //记录总是逃课的人的id的set集合
	for i := 0; i < alwaysRun; i++ {
		idx := rand.Intn(90) //0~89 学生id，对应也是课程中的
		if _, ok := m[idx]; ok {
			i--
		} else {
			m[idx] = struct{}{} //加入集合
		}
	}

	//对于这几个人随机让他们逃课
	for idx, _ := range m {
		student := lesson.Students[idx] //取出对应的学生
		for i := 0; i < ClassNum; i++ {
			f := rand.Float64() //返回0~1的随机数
			if f < 0.8 {
				student.IsRun[i] = true //80%概率逃课
			}
		}
	}
	//每一次都需要找一些人0~3让他们逃课
	for i := 0; i < ClassNum; i++ {
		count := rand.Intn(4) //生成随机人数0~3
		for i := 0; i < count; i++ {
			idx := rand.Intn(90)                 //随机选出下标
			lesson.Students[idx].IsRun[i] = true //设置为true表示逃课
		}
	}

	return lesson
}

//WriteToFile 将lesson数据写入到文件中
func WriteToFile(lesson *Lesson) error {
	//根据课程id获取文件路径
	fpath := fmt.Sprintf("lesson-%d.txt", lesson.Id)
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm) //以创建写入的方式打开文件
	if err != nil {
		return err
	}
	//将所有学生写入至文件中
	for i := range lesson.Students {
		student := lesson.Students[i] //获取student，以json格式进行写入
		jsonData, err := json.Marshal(student)
		if err != nil {
			return err
		}
		_, err = file.Write(jsonData) //写入文件
		_, _ = file.WriteString("\n") //写入换行符
		if err != nil {
			return err
		}
	}
	return file.Sync() //缓冲区数据刷入硬盘
}

//ReadFile 从文件中读取lesson并返回对应数据
func ReadFile(lessonId int) (*Lesson, error) {
	students := make([]*Student, 0)
	lesson := &Lesson{Id: lessonId}

	fpath := fmt.Sprintf("lesson-%d.txt", lesson.Id)
	f, err := os.Open(fpath) //打开文件
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Bytes() //读取一行数据
		student := &Student{}
		err := json.Unmarshal(data, student) //解析student json数据创建student
		if err != nil {
			return nil, err
		}
		students = append(students, student) //添加到lesson中去
	}
	lesson.Students = students
	return lesson, nil
}

// GenerateAllLessonFile  生成所有数据并且以文件形式进行保存
func GenerateAllLessonFile() {
	for i := 0; i < LessonNum; i++ {
		data := GenerateData(i + 1)
		WriteToFile(&data)
	}
}

// ReadAllLessonFile  读取五门课程的所有数据
func ReadAllLessonFile() ([]*Lesson, error) {
	lessons := make([]*Lesson, 0)
	for i := 0; i < LessonNum; i++ {
		lesson, err := ReadFile(i + 1)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
