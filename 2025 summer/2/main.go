package main

import (
	"bufio"
	// "fmt"
	"io"
	"os"
	"strings"

	// "regexp"
	"strconv"
	// "strings"
	// "math"
)


func main(){
	in, out := getBuffers()
	defer out.Flush()

	str, _ := in.ReadString('\n')
	parts := strings.Fields(str)
	m, _ := strconv.Atoi(parts[0]) // ширина
	n, _ := strconv.Atoi(parts[1]) // высота
	m += 1
	n += 1
	width, _ := strconv.Atoi(parts[2])
	height, _ := strconv.Atoi(parts[3])
	quantity, _ := strconv.Atoi(parts[4])
	sl := make([][]string, n+1)
	for j := range sl {
		sl[j] = make([]string, m+1)
	}

	sl_fig := make([][]string, height*2 + 1)
	for j := range sl_fig {
		sl_fig[j] = make([]string, height * 2 + width)
	}

	sl_fig = build(height, width, sl_fig)
	sl = build_schema(m, n, sl)
	// for _, item := range sl{
	// 	res := strings.Join(item,"") + "\n"
	// 	out.WriteString(res)
	// }
	// Читаю sl_fig, вставляю в sl прибавляя + 1
	width_fig := height * 2 + width
	height_fig := 2*height + 1
	start_next_fig := height + width // начиная от начала фигуры в первой строке

	quantity_x_1 := (m-1)/(width_fig + width)
	quantity_x_2 := (m-1-start_next_fig)/(width_fig + width)
	koef_x := 1
	koef_y := 1
	cur_quantity := 0
	line_quantity := 0
	quantity_figs := 0
	for {
		if line_quantity%2==0 {
			koef_x = 1
			quantity_figs = quantity_x_1
		} else {
			koef_x = 1 + height + width
			quantity_figs = quantity_x_2
		}
		for range(quantity_figs){
			for i := range(height_fig){
				for j := range(width_fig){
					x := j + koef_x
					y := i + koef_y
					if sl[y][x] == " "{
						sl[y][x] = sl_fig[i][j]
					}
				}
			}
			cur_quantity += 1
			if cur_quantity == quantity {
				break
			}

			koef_x += width_fig + width
		}
			koef_y += height
			line_quantity += 1
			if cur_quantity == quantity {
				break
			}
	}
	for _, item := range sl{
		res := strings.Join(item,"") + "\n"
		out.WriteString(res)
	}
}

func build_schema(m int, n int, sl [][]string)([][]string){

	for i := range(n+1){
		for j := range(m+1){
			if i==0 && j==0{
				sl[i][j] = "+"
			} else if i==n && j==m{
				sl[i][j] = "+"
			} else if i==n && j==0{
				sl[i][j] = "+"
			} else if i==0 && j==m{
				sl[i][j] = "+"
			}  else if i==0 && j!=0{
				sl[i][j] = "-"
			} else if i!=0 && j==0{
				sl[i][j] = "|"
			} else if i==n && j!=0{
				sl[i][j] = "-"
			} else if i!=0 && j==m{
				sl[i][j] = "|"
			} else {
				sl[i][j] = " "
			} 
		}
	}
	return sl
}

func build(height int, width int, sl [][]string) ([][]string){
	for j := range height*2 + 1 { // Верх
		if j == 0 {
			for i := range height * 2 + width {
				if i < height {
					sl[j][i] = "~"
				} else if i < height + width {
					sl[j][i] = "_"
				} else if i < height * 2 + width {
					sl[j][i] = "~"
				}
			}
		} else if j < height + 1 { // Середина
			for i := range height * 2 + width {
				if i < height - j {
					sl[j][i] = "~"
				} else if i == height - j {
					sl[j][i] = "/"
				} else if i == height + width + j - 1 {
					sl[j][i] = "\\"
				} else if i < height * 2 + width {
					sl[j][i] = "~"
				}
			}
		} else if j >= height + 1 { // Низ
			for i := range height * 2 + width {
				if height + i + 1 == j {
					sl[j][i] = "\\"
				} else if (j == 2 * height) && (i >= height) && (i < height + width) {
					sl[j][i] = "_"
				} else if i == 2*height + width + height - j {
					sl[j][i] = "/"
				} else if i < height * 2 + width {
					sl[j][i] = "~"
				}
			}
		}
	}
	// for _, item := range sl {
	// 	fmt.Println(item)
	// }
	return sl
}


func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var input io.Reader = os.Stdin

	for i, arg := range os.Args {
		if arg == "-f" && i+1 < len(os.Args) {
			f, err := os.Open(os.Args[i+1])
			if err == nil {
				input = f
			}
		}
	}
	return bufio.NewReader(input), bufio.NewWriter(os.Stdout)
}
