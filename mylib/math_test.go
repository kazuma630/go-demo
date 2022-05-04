package mylib
import (
	"testing" // Goの標準テスト用パッケージ
	// "fmt"
)

// var Debug bool = true

// テスト用の関数の定義.
// テスト用の関数名は TestXxx という形式にする.
func TestAverage(t *testing.T) {
	// if Debug {
	// 	t.Skip("Skip Reason")
	// }
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		// fmt.Sprintln("OK")
		t.Errorf("Expected 3, got %d", v)
	}
}