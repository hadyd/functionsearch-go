package main

import "fmt"

func main() {

	kota := []string{"Jakarta", "Bogor", "Depok", "Tangerang", "Bekasi"}

	var searchInput string = "Bogor"

	key, found := search(kota, searchInput)
	if found {
		fmt.Println("Data di Temukan, Ada di Index Ke:", key)
	} else {
		fmt.Println("Data Tidak di Temukan")
	}

}

func search(kota []string, searchInput string) (int, bool) {
	for i, item := range kota {
		if item == searchInput {
			return i, true
		}
	}
	return 0, false
}
