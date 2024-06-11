package dye

import "fmt"

func Red(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[31m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Rred(data ...interface{}) (string){
	var all string 
	for _, d := range data {
		all = fmt.Sprint("\x1b[31m", d, "\x1b[0m")
	}
	return all
}

func Green(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[32m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Cyan(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[36m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Blue(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[34m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Bblue(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[94m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Magenta(data ...interface{}) {
	for _, d := range data {
		fmt.Print("\x1b[35m", d, "\x1b[0m")
	}
	fmt.Println()
}

func Bmagenta(data ...interface{}){
	for _, d := range data{
		fmt.Print("\x1b[95m", d, "\x1b[0m")
	}
	fmt.Println()
}

func White(data ...interface{}) {
	for _, d := range data {
		fmt.Print(d)
	}
	fmt.Println()
}