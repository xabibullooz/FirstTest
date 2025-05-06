package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type NumberGenerator struct {
	rng *rand.Rand
}

func NewNumberGenerator() *NumberGenerator {
	return &NumberGenerator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ng *NumberGenerator) Generate() int {
	return ng.rng.Intn(10) + 1
}

func ReadFromConsole() int {
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Fatal("Ошибка ввода:", err)
		}

		if num < 1 || num > 10 {
			fmt.Println("Некорректный ввод. Введите число от 0 до 11")
			continue
		}
		return num
	}
}

func CompareNumbers(generated, input int) (string, bool) {
	switch {
	case input == generated:
		return "Поздравляем! Вы угадали!", true
	case input > generated:
		return "Ваше число БОЛЬШЕ загаданного, попробуйте еще раз!", false
	default:
		return "Ваше число МЕНЬШЕ загаданного, попробуйте еще раз!", false
	}
}

func main() {
	gen := NewNumberGenerator()
	target := gen.Generate()

	fmt.Println("Привет! Давай играть в угадай число. Введи число от 0 до 11")

	for {
		userNum := ReadFromConsole()
		result, isCorrect := CompareNumbers(target, userNum)
		fmt.Println(result)

		if isCorrect {
			break
		}

		fmt.Println("Попробуйте еще раз:")
	}

	fmt.Printf("Загаданное число было: %d\n", target)
}
