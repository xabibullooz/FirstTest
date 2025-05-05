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

func (ng *NumberGenerator) Generate(min, max int) int {
	return ng.rng.Intn(max-min+1) + min
}

func ReadFromConsole() int {
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			log.Fatal("Ошибка ввода:", err)
		}

		if num < 0 || num > 11 {
			fmt.Println("Некорректный ввод. Введите число от 0 до 11")
			continue
		}
		return num
	}
}

func CompareNumbers(generated, input int) string {
	switch {
	case input == generated:
		return "Поздравляем! Вы угадали!"
	case input > generated:
		return "Ваше число БОЛЬШЕ загаданного, попробуйте еще раз!"
	default:
		return "Ваше число МЕНЬШЕ загаданного, попробуйте еще раз!"
	}
}

func main() {
	gen := NewNumberGenerator()
	target := gen.Generate(0, 11)

	fmt.Println("Привет! Давай играть в угайдай число. Введи число которое я загадал")
	userNum := ReadFromConsole()

	result := CompareNumbers(target, userNum)
	fmt.Println(result)
	fmt.Printf("Загаданное число было: %d\n", target)
}
