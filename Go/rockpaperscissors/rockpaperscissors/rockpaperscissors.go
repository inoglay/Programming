package rockpaperscissors

import (
	"strconv"
)

//IsInput(inputText string, isNum bool)
func IsInput(inputText string, isNum bool) bool {

	// 数値チェックの場合
	if isNum {
		getint, err := strconv.Atoi(inputText)
		if err != nil {
			//数の変換でエラーが発生。つまり、数字ではありません。
			return false
		} else {
			if 1 <= getint && getint <= 3 {
				return true
			}
			return false
		}
	} else {
		if inputText == "Y" || inputText == "N" {
			return true
		}
		return false
	}
}

//GetMessage(messageNo int)
func GetMessage(messageNo int) string {
	switch messageNo {
	case 0:
		return "じゃんけんぽん。。（1：グー、2：チョキ、3；パー）"
	case 1:
		return "あなたの勝ち！！！つづけますか？（Y:はい、N：いいえ）"
	case 2:
		return "あなたの負け・・・つづけますか？（Y:はい、N：いいえ）"
	case 3:
		return "あいこでしょ。。。（1：グー、2：チョキ、3；パー）"
	default:
		return ""
	}
}

//GetName(nameNo int)
func GetName(nameNo int) string {
	switch nameNo {
	case 1:
		return "グー"
	case 2:
		return "チョキ"
	case 3:
		return "パー"
	default:
		return ""
	}
}

//Game(myself int,opponent int)
func Game(myself int, opponent int) int {

	//あいこ
	if myself == opponent {
		return 3
	}

	switch myself {
	//グー
	case 1:
		//チョキ
		if opponent == 2 {
			return 1
		}
		//パー
		return 2
	//チョキ
	case 2:
		//パー
		if opponent == 3 {
			return 1
		}
		//グー
		return 2
	//パー
	default:
		//グー
		if opponent == 1 {
			return 1
		}
		//チョキ
		return 2
	}
}
