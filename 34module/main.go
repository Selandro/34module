package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Открытие входного файла
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer inputFile.Close()

	// Создание файла для выводов результатов,
	//os.Create создает новый файл или при открытии очищает
	//существуюший
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	// Регулярное выражение для поиска математических выражений
	re := regexp.MustCompile(`(\d+)([+\-*/])(\d+)=\?`)

	// Сканер для построчного чтения входного файла
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		// Поиск соответствия регулярному выражению
		find := re.FindStringSubmatch(line)
		if find != nil {
			num1, _ := strconv.Atoi(find[1])
			operator := find[2]
			num2, _ := strconv.Atoi(find[3])

			var result int
			// Выполнение математической операции
			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 != 0 {
					result = num1 / num2
				} else {
					// Пропускаем деление на ноль
					continue
				}
			}

			// Запись результата в выходной файл
			outputLine := fmt.Sprintf("%s%d\n", line[:len(line)-1], result)
			writer.WriteString(outputLine)
		}
	}

	// Проверка на наличие ошибок при сканировании файла
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}

	// Записываем все строки в файл
	writer.Flush()
}
