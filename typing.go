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
		// 単語を表示
		fmt.Println(v)
		// 開始時間をセット
		stm := time.Now()
		// 入力欄を表示
		fmt.Print("input: ")
		// 入力内容を取得して比較
		var in string
		fmt.Scanln(&in)
		// かかった時間を取得
		etm := time.Since(stm)
		// 入力結果を表示
		if v == in {
			corcnt++
			fmt.Println("○", etm)
		} else {
			miscnt++
			fmt.Println("×", etm)
		}
	}

	// 結果の計算
	sumcnt = corcnt + miscnt
	corpct = int(float64(corcnt)/float64(sumcnt)*100)
	mispct = int(float64(miscnt)/float64(sumcnt)*100)
	
	// 結果の出力
	fmt.Printf("[正誤] 正解: %d, 誤り: %d\n", corcnt, miscnt)
	fmt.Printf("[正誤率] 正解: %d％, 誤り: %d％\n", corpct, mispct)
}
