//пакет пердназначет для чтения данных из файлов
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// getFloats читает значение float64 из каждой строки файла
func getFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64         //объявляем возвращаемый массив
	file, err := os.Open(fileName) //открываем файл с переданным именем
	if err != nil {
		return numbers, err
	}
	i := 0 //переменная для индексирования массива
	sca := bufio.NewScanner(file)
	for sca.Scan() {
		numbers[i], err = strconv.ParseFloat(sca.Text(), 64) //строка прочитанная из файла преобразуется в float64
		if err != nil {
			return numbers, err
		}
		i++ // переход к следующему индексу массива

		err = file.Close()
		if err != nil {
			return numbers, err
		}
		if sca.Err() != nil {
			return numbers, sca.Err()
		}

	}
	return numbers, nil
}
