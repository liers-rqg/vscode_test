package demo6

import (
	"bufio"
	"fmt"
	"os"
)

type CountFromFile struct {
	Chinese int
	English int
	Number  int
	Space   int
	Other   int
}

// os.open中的路径是从项目路径开始找的(src/demo6/test.txt)
func CountChar(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("open error: %v !\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	cff := new(CountFromFile)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil {
			fmt.Println("read error !")
			break
		}
		for _, c := range line {
			switch {
			case c >= 'a' && c <= 'z':
				fallthrough
			case c >= 'A' && c <= 'Z':
				cff.English++
			case c == ' ' || c == '\t':
				cff.Space++
			case c >= '0' && c <= '9':
				cff.Number++
			default:
				cff.Other++
			}
		}
	}
	fmt.Println(cff)

}
