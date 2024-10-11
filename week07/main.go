package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())

	var army string = "우리는 !가와 !민에 충성을 다하는 대한민! 육군이다."
	armyFixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.Println(armyFixed)

	fmt.Print("이름 입력 : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}
