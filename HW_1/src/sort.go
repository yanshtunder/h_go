package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"os"
	"sort"
	"strings"
)

func main() {

	/**
	 * -f = string file for sorting
	 * -k = int    compare by column number
	 * -r = bool   reverse the result of comparisons
	 * -u = bool   not output repeated rows
	 */
	f, k, r, u := initFlags()

	line, err := readLines(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("--- BEFORE SORTING ---")
	for _, v := range line {
		fmt.Println(v)
	}

	maxLenRow := getMaxLenRow(&line)
	if k >= maxLenRow {
		fmt.Println("out of bounds")
		os.Exit(1)
	}

	/**
	 * Preparation of a two-dimensional slice.
	 * If the length of some line is less than maxLenRow, it is necessary to fill the remaining space with "".
	 */
	arr := Make2D[string](len(line), maxLenRow)
	for i, item := range line {
		arr[i] = strings.Split(item, " ")
		for len(arr[i]) < maxLenRow {
			arr[i] = append(arr[i], "")
		}
	}

	/**
	 * Sorting and print slice
	 */
	var row int
	arr, row = sortSlice(&arr, line, maxLenRow, k, r, u)
	fmt.Printf("\n--- AFTER SORTING ---\n")
	for i := 0; i < row; i++ {
		for j := 0; j < maxLenRow; j++ {
			fmt.Printf("%s ", arr[i][j])
		}
		fmt.Println()
	}

}

func initFlags() (string, int, bool, bool) {
	f := flag.String("f", "../input.txt", "file for sorting")
	k := flag.Int("k", 0, "compare by column number")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "not output repeated rows")
	flag.Parse()
	return *f, *k, *r, *u
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getMaxLenRow(row *[]string) int {
	maxLen := 0
	for _, item := range *row {
		arr := strings.Split(item, " ")
		if maxLen <= len(arr) {
			maxLen = len(arr)
		}
	}
	return maxLen
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func sortSlice(arr *[][]string, line []string, maxLenRow int, k int, r bool, u bool) ([][]string, int) {
	cl := collate.New(language.Russian)
	sort.Slice(*arr, func(i, j int) bool {
		if r {
			return cl.CompareString((*arr)[i][k], (*arr)[j][k]) == 1
		} else {
			return cl.CompareString((*arr)[i][k], (*arr)[j][k]) == -1
		}
	})

	if u {
		prev := 1
		for curr := 1; curr < len(line); curr++ {
			for j := 0; j < maxLenRow; j++ {
				if (*arr)[curr-1][j] != (*arr)[curr][j] {
					(*arr)[prev] = (*arr)[curr]
					prev++

					break
				}
			}
		}
		*arr = (*arr)[:prev]
	}
	return *arr, len(*arr)
}
