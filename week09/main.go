package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1~6
	fmt.Println(answer)

	var win bool
	win = true
	for guessess := 0; guessess < 3; guessess++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guessess)
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input) //줄바꿈,띄어쓰기, 탭 등 제거
		// guess, err := strconv.ParseInt(input,10,32) //정수형 32비트 타입으로 형변환, 16진수 입력
		guess, err := strconv.Atoi(input) //정수형 32비트 타입으로 형변환, 16진수 입력
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작은 수 입니다.")
		} else { //answer < guess
			fmt.Println("입력하신 값은 정답보다 큰 수 입니다.")
		}
	}

	if win {
		fmt.Println("당신이 이겼습니다.")
	} else {
		fmt.Printf("당신이 졌습니다. 정답은 %d입니다.\n", answer)
	}
}
