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

	var corcnt, miscnt, sumcnt int
	var corpct, mispct int

	for _, v := range datas {
		fmt.Println(v)
		fmt.Print("input: ")
		var in string
		fmt.Scanln(&in)
		if v == in {
			corcnt++
			fmt.Println("○")
		} else {
			miscnt++
			fmt.Println("×")
		}
	}

	sumcnt = corcnt + miscnt
	corpct = int(float64(corcnt)/float64(sumcnt)*100)
	mispct = int(float64(miscnt)/float64(sumcnt)*100)
	
	// 結果の出力
	fmt.Printf("[正誤] 正解: %d, 誤り: %d\n", corcnt, miscnt)
	fmt.Printf("[正誤率] 正解: %d％, 誤り: %d％\n", corpct, mispct)
}
