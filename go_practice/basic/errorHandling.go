package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./nothing.txt")
	if err != nil {
		log.Fatalln("ERROR!!")
	}
	defer file.Close()
	data := make([]byte, 100)
	// 上のerrとは別物。複数変数を渡したときにどちらかが一意になっていてば、 := が使える
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("ERROR!!")
	}
	fmt.Println(count, string(data))

	// ここでは := を使えないってこと 一つ上のerrをオーバーライドしている
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("ERROR!!")
	}
}