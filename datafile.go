package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func Getfloat(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName) /// открывает файл с указаным именем kills.txt в этом случае
	if err != nil {                /// выводит ошибку, если она возникла на этапе открытия
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file) /// создается новое значение scanner для переменной file
	for scanner.Scan() {              /// читает файл
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) /// выводит строки содержащиеся в файле
		if err != nil {
			return numbers, err
		}
		i++
		/// цикл завершается по окончанию файла
	}
	err = file.Close() /// закрывает файл, освобождая ресурсы
	if err != nil {    /// выводит ошибку, которая могла возникнуть при закрытии файла
		return numbers, err
	}
	if scanner.Err() != nil { /// выводит ошибку, которая могла возникнуть в процессе сканирования файла
		return numbers, scanner.Err()
	}
	return numbers, nil
}
