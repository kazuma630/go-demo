package main

import (
	"fmt"
	"math"
	"go-demo/pokemon"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=1; i<=10; i++ {
		z -= (z*z - x) / (2*z)

		// fmt.Sprintf 書式指定子に沿ってフォーマットした文字列を返す。%d 基数10。基数とは位が上がる毎に何倍になるかを表し、今回の基数10はすなわち10進数ということ。
		countNumber := fmt.Sprintf("%d番目", i)
		fmt.Println(countNumber, z)
		if math.Sqrt(x) == z {
			fmt.Println("近似値は下記になります。")
			break // 近似した場合は反復処理を抜ける。continueは、処理のスキップ（ループ自体は抜けない）
		}
	}
	return z
}

func main() {
  fmt.Println(pokemon.KantouBigThreePokemon())
	fmt.Println(pokemon.HouenBigThreePokemon())
	fmt.Println(pokemon.SecretBigThreePokemon())
	fmt.Println(Sqrt(2))
}
