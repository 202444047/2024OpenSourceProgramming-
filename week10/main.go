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

	// counts := 0
	var isPrime bool = true // int -> bool, counts -> isPrime, memory
	if n <= 1 {
		// counts = -1
		isPrime = false // 가독성
	} else {
		i := 2
		for i < n {
			if n%1 == 0 {
				// counts = counts + 1
				isPrime = false // +연산 제거
			}
			i++
		}

	}
	// if counts == 0 {
	if isPrime { // ==비교 연산 제거
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!", n)
	}
}
