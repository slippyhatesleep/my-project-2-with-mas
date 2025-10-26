package main

import (
	"fmt"
	"os"
)

var sum int

func summa(mas []int) int {
	sum = 0
	for _, i := range mas {
		sum += i
	}
	return sum
}
func average(mas []int) float32 {
	if len(mas) == 0 {
		return 0
	}
	return float32(summa(mas)) / float32(len(mas))
}
func filetxt(mas []int) error {
	file, err := os.Create("input.txt")
	defer file.Close()
	_, err = file.WriteString("Массив: " + fmt.Sprint(mas) + "\n")
	_, err = file.WriteString("Среднее значение: " + fmt.Sprint(average(mas)) + "\n")
	_, err = file.WriteString("Сумма массива: " + fmt.Sprint(summa(mas)) + "\n")
	if err != nil {
		fmt.Println("Error input.txt")
	}
	return nil
}
func main() {
	var mas []int
	var input string
	var inputmas int
	for {
		fmt.Println("Впиши цифру от 1 до 4\n" +
			"1 - добавить массив\n" +
			"2 - сумма чисел в массиве\n" +
			"3 - средний показатель\n" +
			"4 - выход\n" +
			"5 - сохранить массив в txt файл\n")
		fmt.Scan(&input)
		if input == "4" {
			fmt.Println("zakrit")
			break
		}
		if input == "1" {
			fmt.Println("Введите число для массива")
			fmt.Scan(&inputmas)
			mas = append(mas, inputmas)
			fmt.Println(mas)
			continue
		}
		if input == "2" {
			sum = summa(mas)
			fmt.Printf("Сумма элементов массива %v\n", sum)
			continue
		}
		if input == "3" {
			fmt.Printf("Среднее значение элементов  %.2f\n", average(mas))
			continue
		}
		if input == "5" {
			file := filetxt(mas)
			sum = summa(mas)
			average(mas)
			if file != nil {
				fmt.Errorf("Ошибка %v", file)
			} else {
				fmt.Println("Ваш файл сохранен в input.txt")
				fmt.Println("Содержимое файла: сам массив: ", mas, "\n", " Среднее значение: ", average(mas), "\n", " Сумма массива: ", summa(mas))
			}
			continue
		}
	}
}
