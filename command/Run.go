package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)

	var x int64
	cells := make([]int64, 3000)
	curr := int64(00)
	for {
		fmt.Print("\n>> ")

		input, _ := reader.ReadString('\n')

		for _, c := range input {
			switch c {
			case 's':
				cells[curr] = cells[curr] * cells[curr]
			case 'i':
				cells[curr]++
			case 'd':
				cells[curr]--
			case 'o':
				fmt.Print(strconv.FormatInt(cells[curr], 10) + " ")
			case 'c':
				fmt.Print(string(rune(cells[curr])))
			case 'h':
				curr++
			case 'r':
				curr--
			case 'm':
				fmt.Print("\nMemory:")
				for index, cell := range cells {
					if cell != 0 {
						fmt.Print("\n", index, " ", cell, " ", string(rune(cell)))
					}
				}
			}
			if x < 0 || x == 256 {
				x = 0
			}
		}
	}
}
