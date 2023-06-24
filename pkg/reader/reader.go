package reader

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Получение строк
func StrReader(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.Trim(data, "\n")
	data = strings.Replace(data, "\n", "", -1)
	return data, nil
}

// Получение чисел
func IntReader(r io.Reader) (int, error) {
	data, err := StrReader(r)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func IntArrReader(r io.Reader) ([]int, error) {
	var (
		ret = []int{}
	)
	data, err := StrReader(r)
	if err != nil {
		return nil, err
	}
	sList := strings.Split(data, " ")
	for _, v := range sList {
		if v == "" {
			continue
		}
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, num)

	}
	return ret, nil
}
