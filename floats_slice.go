//пакет пердназначет для чтения данных из файлов
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// getFloats читает значение float64 из каждой строки файла
func getFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0 //переменная для индексирования массива
	sca := bufio.NewScanner(file)
	for sca.Scan() {
		number, err = strconv.ParseFloat(sca.Text(), 64) //строка прочитанная из файла преобразуется в float64
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)

		if sca.Err() != nil {
			return numbers, sca.Err()
		}

	}
	return numbers, nil
}
