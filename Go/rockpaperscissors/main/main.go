package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"example.com/rockpaperscissors"
)

//メイン処理
func main() {

	msg := rockpaperscissors.GetMessage(0)
	//じゃんけんぽん。。（1：グー、2：チョキ、3；パー）
	fmt.Printf(msg)
	isNum := true
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		//入力チェック
		chk := rockpaperscissors.IsInput(text, isNum)
		if chk {
			if isNum {
				rand.Seed(time.Now().Unix())
				opponent := rand.Intn(2) + 1
				myself, err := strconv.Atoi(text)
				if err == nil {
					result := rockpaperscissors.Game(myself, opponent)
					msg = rockpaperscissors.GetMessage(result)
					fmt.Println("あなた：" + rockpaperscissors.GetName(myself))
					fmt.Println("あいて：" + rockpaperscissors.GetName(opponent))
					fmt.Printf(msg)
					if result < 3 {
						isNum = false
					}
				}
			} else {
				if text == "Y" {
					isNum = true
					msg = rockpaperscissors.GetMessage(0)
					fmt.Printf(msg)
				} else {
					break
				}
			}
		} else {
			fmt.Printf(msg)
		}
	}
}
