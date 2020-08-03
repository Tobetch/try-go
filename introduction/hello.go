// 標準ライブラリ一覧
// math：算術用。Sin, Cos, Tan, Sqrtなど
// compress：圧縮用
// crypto：圧縮用
// image：画像処理用
package main

// フォーマット用関数 "fmt" をインポート
import (
	"fmt"
)

// main関数の宣言
// Goプログラムはmainパッケージのmain関数から実行される
func main() {
	variousLanguages()
}

func variousLanguages() {
	fmt.Println("Hello, world") //Javaにおける"System.out.println"メソッド
	fmt.Println("你好，世界")
	fmt.Println("Здравствуй, мир")
	fmt.Println("Hola, mundo")
}
