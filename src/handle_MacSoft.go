package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
|      |      |      |
| ---- | ---- | ---- |
|      |      |      |
|      |      |      |
|      |      |      |
*/
func main() {

	fmt.Println(`|  序号   |  软件 | 中文名  |  说明   |
| ---- | ---- | ---- | ---- |`)

	f, _ := os.Open("soft.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Fields(line)
		if len(items) < 9 {
			continue
		}
		fn := strings.Join(items[8:], " ")
		if fn == "soft.txt" {
			continue
		}
		i++
		fmt.Printf("| %d ", i)
		//r := strings.NewReader(line)
		//str_scner := bufio.NewScanner(r)
		fmt.Printf("| %s ", fn)
		fmt.Printf("| %s | xxx |\n", fn)
	}
}
