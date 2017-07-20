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
	//local_bin from ls -l
	f, _ := os.Open("local_bin.txt")
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
		i++
		fmt.Printf("| %d ", i)
		terms := strings.Split(fn, "->")
		fmt.Printf("| %s ", fn)
		fmt.Printf("| %s | xxx |\n", terms[0])
	}
}
