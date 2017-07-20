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
	d := make(map[string]int)
	//brew.txt from history|grep 'brew install'  &  history|grep 'brew cask'
	//f, _ := os.Open("brew.txt")
	//pip.txt from history|grep 'pip install'
	f, _ := os.Open("pip.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		items := strings.Fields(line)
		fn := items[len(items)-1]
		_, ok := d[fn]
		if ok {
			continue
		}
		d[fn] = 1
		//fn := strings.Join(items[], " ")
		i++
		fmt.Printf("| %d ", i)
		//r := strings.NewReader(line)
		//str_scner := bufio.NewScanner(r)
		fmt.Printf("| %s ", fn)
		fmt.Printf("| %s | xxx |\n", fn)
	}
}
