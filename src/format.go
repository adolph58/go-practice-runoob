package main

import "fmt"

func main() {
	var stackcode = 123
	var enddate = "2025-10-11"
	var url = "Code=%d&endData=%s"
	var target_url = fmt.Sprintf(url, stackcode, enddate)
	fmt.Println(target_url)
}
