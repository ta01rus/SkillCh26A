package readers

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
