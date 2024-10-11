package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err != in.ReadString('\n')
	if err != nil {1
		log.Fatal(err)
	} 
	score = strings.TrimSpace(score) //줄바꿈,띄어쓰기, 탭 등 제거
	realScore, _ := strconv.ParseInt(score,16,32) //실수형 64비트 타입으로 형변환
	if realScore >= 90 {
		grade :="A"
	} else {
		grade :="BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다.\n", realScore,grade)
}
