package main

import "fmt"

func main() {
	week := 5

	switch week {
	case 1:
		fmt.Println("月曜日")
	case 2:
		fmt.Println("火曜日")
	case 3:
		fmt.Println("水曜日")
	case 4:
		fmt.Println("木曜日")
	case 5:
		fmt.Println("金曜日")
	case 6:
		fmt.Println("土曜日")
	case 7:
		fmt.Println("日曜日")
	default:
		fmt.Println("範囲対象外です")

	}

}

