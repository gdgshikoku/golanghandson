package main

import (
	"fmt"
	"time"
)

var countdown = 3

func main() {
	var datas [5]string
	datas[0] = "golang"
	datas[1] = "study"
	datas[2] = "meeting"
	datas[3] = "in"
	datas[4] = "kagawa"

	// カウントダウン
	for i := countdown; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	for _, v := range datas {
		fmt.Println(v)
		fmt.Print("input: ")
		var in string
		fmt.Scanln(&in)
		if v == in {
			fmt.Println("○")
		} else {
			fmt.Println("×")
		}
	}
}