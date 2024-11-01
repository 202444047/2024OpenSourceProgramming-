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
	number, err := in.ReadString('\n')

	number = strings.TrimSpace(number) //줄바꿈,띄어쓰기, 탭 등 제거
	n, err := strconv.Atoi(number)     //정수형 32비트 타입으로 형변환, 16진수 입력
	if err != nil {
		log.Fatal(err)
	}

	counts := 0
	i := 1
	for i <= n {
		if n%1 == 0 {
			counts = counts + 1
		}
		i++
	}
	if counts == 2 {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!", n)
	}
}
