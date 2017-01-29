package main

// 依存パッケージのインポート
import (
	"fmt"
)

// Person構造体の定義
type Person struct {
	FirstName string // メンバー変数の宣言。型が後に来る
	LastName  string
}

// Person構造体のオブジェクトを作る関数の定義。
// 戻り値の型は後ろに書く。
func NewPerson(firstName string, lastName string) *Person {
	// Person{…}は新しい構造体を作り、&はそのアドレスを取る。
	p := &Person{
		FirstName: firstName,
		LastName:  lastName,
	}
	return p
}

// Person構造体のHelloメソッドの定義。
// レシーバ、関数シグニチャ、戻り値の型という順番で記述する。
func (p *Person) Hello() {
	fmt.Printf("Hello, I'm %s %s.\n", p.FirstName, p.LastName)
}

// エントリーポイントであるmain関数の定義。
func main() {
	m := NewPerson("Michael", "Jackson")
	m.Hello()
}
