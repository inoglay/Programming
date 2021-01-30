// rockpaperscissors_test.go
package rockpaperscissors

import (
	"testing"
)

//数値フラグTrue時チェック
func TestRockpaperscissors(t *testing.T) {

	//入力チェック(数値)
	got := IsInput("1", true)
	if got != true {
		t.Errorf(`IsInput(1,true) is %t`, got)
	}

	got = IsInput("2", true)
	if got != true {
		t.Errorf(`IsInput(1,true) is %t`, got)
	}

	got = IsInput("3", true)
	if got != true {
		t.Errorf(`IsInput(1,true) is %t`, got)
	}

	got = IsInput("0", true)
	if got != false {
		t.Errorf(`IsInput(1,true) is %t`, got)
	}

	got = IsInput("4", true)
	if got != false {
		t.Errorf(`IsInput(1,true) is %t`, got)
	}

	got = IsInput("A", true)
	if got != false {
		t.Errorf(`IsInput("A", true) is %t`, got)
	}

	got = IsInput("", true)
	if got != false {
		t.Errorf(`IsInput("", true) is %t`, got)
	}

	//入力チェック(文字)
	got = IsInput("Y", false)
	if got != true {
		t.Errorf(`IsInput("Y", false) is %t`, got)
	}

	got = IsInput("N", false)
	if got != true {
		t.Errorf(`IsInput("N", false) is %t`, got)
	}

	got = IsInput("A", false)
	if got != false {
		t.Errorf(`IsInput("A", false) is %t`, got)
	}

	got = IsInput("1", false)
	if got != false {
		t.Errorf(`IsInput("1", false) is %t`, got)
	}

	got = IsInput("", false)
	if got != false {
		t.Errorf(`IsInput("", false) is %t`, got)
	}

	//メッセージ取得
	message := GetMessage(0)
	if message != "じゃんけんぽん。。（１：グー、２：チョキ、３；パー）" {
		t.Errorf(`GetMessage(0) is %q`, message)
	}

	message = GetMessage(1)
	if message != "あなたの勝ち！！！つづけますか？（Y:はい、N：いいえ）" {
		t.Errorf(`GetMessage(1) is %q`, message)
	}

	message = GetMessage(2)
	if message != "あなたの負け・・・つづけますか？（Y:はい、N：いいえ）" {
		t.Errorf(`GetMessage(2) is %q`, message)
	}

	message = GetMessage(3)
	if message != "あいこでしょ。。。（１：グー、２：チョキ、３；パー）" {
		t.Errorf(`GetMessage(3) is %q`, message)
	}

	message = GetMessage(4)
	if message != "" {
		t.Errorf(`GetMessage(4) is %q`, message)
	}

	//名称取得
	name := GetName(0)
	if name != "" {
		t.Errorf(`GetName(0) is %q`, name)
	}

	name = GetName(1)
	if name != "グー" {
		t.Errorf(`GetName(1) is %q`, name)
	}

	name = GetName(2)
	if name != "チョキ" {
		t.Errorf(`GetName(2) is %q`, name)
	}

	name = GetName(3)
	if name != "パー" {
		t.Errorf(`GetName(3) is %q`, name)
	}

	name = GetName(4)
	if name != "" {
		t.Errorf(`GetName(3) is %q`, name)
	}

	//じゃんけんゲーム
	result := Game(1, 1)
	if result != 3 {
		t.Errorf(`Game(1, 1) is %v`, result)
	}

	result = Game(1, 2)
	if result != 1 {
		t.Errorf(`Game(1, 2) is %v`, result)
	}

	result = Game(1, 3)
	if result != 2 {
		t.Errorf(`Game(1, 3) is %v`, result)
	}

	result = Game(2, 2)
	if result != 3 {
		t.Errorf(`Game(2, 2) is %v`, result)
	}

	result = Game(2, 3)
	if result != 1 {
		t.Errorf(`Game(2, 3) is %v`, result)
	}

	result = Game(2, 1)
	if result != 2 {
		t.Errorf(`Game(2, 1) is %v`, result)
	}

	result = Game(3, 3)
	if result != 3 {
		t.Errorf(`Game(3, 3) is %v`, result)
	}

	result = Game(3, 1)
	if result != 1 {
		t.Errorf(`Game(3, 1) is %v`, result)
	}

	result = Game(3, 2)
	if result != 2 {
		t.Errorf(` Game(3, 2) is %v`, result)
	}

}
