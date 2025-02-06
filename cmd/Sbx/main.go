package main

import "fmt"

func main() {
	fmt.Println("Hallo Welt")

	texts := []string{"Daniel", "Stonebuilt", "Hosana", "Natieli", "Das ist", "ein", "test"}

	fmt.Printf("Size: %d\n", len(texts))
	fmt.Printf("Capacity: %d\n", cap(texts))
	fmt.Println(texts)
	//fmt.Println(texts[cap(texts)-1:])
	r1 := pop(texts)
	fmt.Println(r1)
	r2 := pop(r1)
	fmt.Println(r2)
	//fmt.Println(texts[cap(texts)-1:])
	//fmt.Println(deleteElement(texts, 1))
	//fmt.Println(texts[0:2])
	//fmt.Println(texts[2:])
	fmt.Printf("Size R1: %d\n", len(r1))
	fmt.Printf("Capacity R1: %d\n", cap(r1))

	fmt.Printf("Size R2: %d\n", len(r2))
	fmt.Printf("Capacity R2: %d\n", cap(r2))
}

func deleteElement(slice []string, index int) []string {

	println("pos init: ", slice[:index])
	println("pos end: ", slice[index+1:])
	return append(slice[:index], slice[index+1:]...)
}

func pop(slice []string) []string {

	if len(slice) == 0 {
		return nil
	}
	return append([]string{}, slice[:cap(slice)-1]...)
}
